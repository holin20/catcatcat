package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
	"go.uber.org/zap"
)

type alertState struct {
	alerting bool

	notifyHistory []*ezgo.Tuple2_[time.Time, float64]
}

type Monitor struct {
	scope     *ezgo.Scope
	scheduler *ezgo.Scheduler
	notifier  *Notifier

	ruleConfigs map[string]*RuleConfig

	// materialized rules and alert states
	rules       map[string]*Rule[float64]
	alertStates map[string]*alertState

	evalInterval   time.Duration
	notifyInterval time.Duration
}

func NewMonitor(scope *ezgo.Scope) *Monitor {
	scope = scope.WithLogger(scope.GetLogger().Named("Monitor"))
	m := &Monitor{
		scope:          scope,
		scheduler:      ezgo.NewScheduler(scope),
		notifier:       NewNotifier(),
		notifyInterval: 24 * time.Hour,
	}
	return m
}

func (m *Monitor) WithRuleConfigs(ruleConfigs map[string]*RuleConfig) *Monitor {
	m.ruleConfigs = ruleConfigs
	m.buildRules()
	return m
}

func (m *Monitor) WithEvalInterval(evalInterval time.Duration) *Monitor {
	m.evalInterval = evalInterval
	return m
}

func (m *Monitor) Kickoff(ctx context.Context) {
	evalInterval := m.evalInterval
	if evalInterval == 0 {
		evalInterval = time.Minute
	}
	m.scheduler.Repeat(ctx, evalInterval, "Monitor", func() {
		if err := m.evalRules(ctx, time.Now()); err != nil {
			ezgo.LogCausesf(m.scope.GetLogger(), err, "evalRules")
		}
	})
}

func (m *Monitor) evalRules(ctx context.Context, now time.Time) error {
	awaitbles := ezgo.MapApplyAsync(m.rules, func(ruleId string, r *Rule[float64]) bool {
		m.scope.GetLogger().Info("Evaluating rule", zap.String("name", r.GetName()), zap.String("rule#", ruleId))
		met, queryResult, queryResultTime, err := r.Eval(ctx, now)
		if err != nil {
			ezgo.LogCausesf(m.scope.GetLogger(), err, "Eval")
			return false
		}
		if met {
			m.alert(r, m.ruleConfigs[ruleId], queryResult, queryResultTime, now)
		} else if rs := m.alertStates[ruleId]; rs != nil && rs.alerting {
			rs.alerting = false // reset the alerting state.
			// TODO - notify "alert dismissed"
		}
		return met
	})

	metStatusSlice := ezgo.AwaitMapAll(awaitbles)
	metStatusSlice = ezgo.MapTrueFilter(metStatusSlice)
	m.scope.GetLogger().Info(fmt.Sprintf("%d rules are met", len(metStatusSlice)))
	return nil
}

func (m *Monitor) alert(
	r *Rule[float64],
	ruleConfig *RuleConfig,
	queryResult float64,
	queryResultTime time.Time,
	queryTime time.Time,
) {
	m.scope.GetLogger().Info(
		"Notify rule is met",
		zap.String("rule", r.GetName()),
		zap.Float64("query_result", queryResult),
		zap.Time("query_result_time", queryResultTime),
		zap.Duration("result_delay", queryTime.Sub(queryResultTime)),
		zap.Time("query_time", queryTime),
	)

	now := time.Now()
	alertState := ezgo.EnsureMapValue(m.alertStates, r.id)
	if alertState.alerting {
		// Currently in alerting state. Check if we need to notify again. Reset alerting flag if needed.
		if len(alertState.notifyHistory) == 0 {
			ezgo.LogCauses(m.scope.GetLogger(), fmt.Errorf("empty notif history"), "alertHistory")
		} else {
			lastNotifyTime, _ := ezgo.Last(alertState.notifyHistory).Unpack()
			if now.Sub(lastNotifyTime) > m.notifyInterval {
				alertState.alerting = false
			}
		}
	}

	if !alertState.alerting {
		alertState.alerting = true
		alertState.notifyHistory = append(alertState.notifyHistory, ezgo.Tuple2(now, queryResult))
	}

	if err := m.notifier.NotifyEmail(
		r.GetName(),
		ruleConfig,
		queryTime,
		queryResult,
		queryResultTime,
	); ezgo.IsErr(err) {
		ezgo.LogCauses(m.scope.GetLogger(), err, "NotifyEmail")
	}
}

func (m *Monitor) buildRules() {
	rules := make(map[string]*Rule[float64])
	for i, rc := range m.ruleConfigs {
		rules[i] = &Rule[float64]{
			id:   rc.RuleId,
			name: rc.Name,
			query: ezgo.Must(BuildQuery[float64](
				rc.QueryType,
				rc.QueryArgs...,
			)),
			condition: ezgo.Must(BuildCondition[float64](
				rc.ConditionType,
				rc.ConditionArgs,
			)),
		}
	}
	m.rules = rules

	m.scope.GetLogger().Info("Rules built", zap.Int("rule count", len(m.rules)))
}

func (m *Monitor) Terminate() {
	m.scheduler.Terminate()
}
