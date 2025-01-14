package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
	"go.uber.org/zap"
)

const (
	defaultEvalInterval   = time.Minute
	defaultNotifyInterval = 24 * time.Hour
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
		notifyInterval: defaultNotifyInterval,
		evalInterval:   defaultEvalInterval,
		alertStates:    make(map[string]*alertState),
	}
	return m
}

func (m *Monitor) WithRuleConfigs(ruleConfigs map[string]*RuleConfig) *Monitor {
	m.ruleConfigs = ruleConfigs
	m.buildRules()
	return m
}

func (m *Monitor) WithEvalInterval(evalInterval time.Duration) *Monitor {
	if evalInterval > 0 {
		m.evalInterval = evalInterval
	}
	return m
}

func (m *Monitor) WithNotifyInterval(notifyInterval time.Duration) *Monitor {
	if m.notifyInterval > 0 {
		m.notifyInterval = notifyInterval
	}
	return m
}

func (m *Monitor) Kickoff(ctx context.Context) {
	m.scope.GetLogger().Info(
		"Kicking off monitor!",
		zap.Duration("eval_interval", m.evalInterval),
		zap.Duration("notify_interval", m.notifyInterval),
	)

	m.scheduler.Repeat(ctx, m.evalInterval, "Monitor", func() {
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
			m.scope.GetLogger().Info("Rule no longer meets the alerting criteria", zap.String("ruleId", ruleId))
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
				m.scope.GetLogger().Info(
					"Alert notification interval passed. Alert again",
					zap.String("rule_id", r.id),
					zap.Time("last_notify_time", lastNotifyTime),
				)
				alertState.alerting = false
			} else {
				m.scope.GetLogger().Info(
					"Inside alert notification interval. Do not notify",
					zap.String("rule_id", r.id),
					zap.Time("last_notify_time", lastNotifyTime),
					zap.Time("next_notify_time", lastNotifyTime.Add(m.notifyInterval)),
					zap.Duration("notify_interval", m.notifyInterval),
				)
			}
		}
	}

	if !alertState.alerting {
		alertState.alerting = true
		alertState.notifyHistory = append(alertState.notifyHistory, ezgo.Tuple2(now, queryResult))
		m.scope.GetLogger().Info("Notify this alert!", zap.String("ruleId", r.id))

		// Notify!
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
