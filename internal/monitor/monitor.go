package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
	"go.uber.org/zap"
)

type Monitor struct {
	scope     *ezgo.Scope
	scheduler *ezgo.Scheduler

	rules       []*Rule[float64]
	ruleConfigs []*RuleConfig

	evalInterval time.Duration
}

func NewMonitor(scope *ezgo.Scope) *Monitor {
	scope = scope.WithLogger(scope.GetLogger().Named("Monitor"))
	m := &Monitor{
		scope:     scope,
		scheduler: ezgo.NewScheduler(scope),
	}
	return m
}

func (m *Monitor) WithRuleConfigs(ruleConfigs []*RuleConfig) *Monitor {
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
	awaitbles := ezgo.SliceApplyAsync(m.rules, func(i int, r *Rule[float64]) bool {
		m.scope.GetLogger().Info("Evaluating rule", zap.String("name", r.GetName()), zap.Int("rule#", i))
		met, queryResult, queryResultTime, err := r.Eval(ctx, now)
		if err != nil {
			ezgo.LogCausesf(m.scope.GetLogger(), err, "Eval")
			return false
		}
		if met {
			m.notify(r, queryResult, queryResultTime, now)
		}
		return met
	})

	metStatusSlice := ezgo.Await(awaitbles...)
	metStatusSlice = ezgo.SliceFilter(metStatusSlice, func(b bool) bool { return b })
	m.scope.GetLogger().Info(fmt.Sprintf("%d rules are met", len(metStatusSlice)))
	return nil
}

func (m *Monitor) notify(
	r *Rule[float64],
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
}

func (m *Monitor) buildRules() {
	var rules []*Rule[float64]
	for _, rc := range m.ruleConfigs {
		rules = append(rules, &Rule[float64]{
			name: rc.Name,
			query: ezgo.Must(BuildQuery[float64](
				rc.QueryType,
				rc.QueryArgs...,
			)),
			condition: ezgo.Must(BuildCondition[float64](
				rc.ConditionType,
				rc.ConditionArgs,
			)),
		})
	}
	m.rules = rules

	m.scope.GetLogger().Info("Rules built", zap.Int("rule count", len(m.rules)))
}

func (m *Monitor) Terminate() {
	m.scheduler.Terminate()
}
