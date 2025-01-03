package monitor

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/pkg/ezgo"
)

type Monitor struct {
	scope     *ezgo.Scope
	scheduler *ezgo.Scheduler
	rules     []*Rule[float64]
}

var RULES_CONFIG = []struct {
	name          string
	queryType     QueryType
	conditionType ConditionType
	conditionArgs float64
}{
	{"Macbook", eEcho, eLess, 1050.0},
	{"Face", eEcho, eLess, 100000.0},
}

func NewMonitor(scope *ezgo.Scope) *Monitor {
	m := &Monitor{
		scope:     scope,
		scheduler: ezgo.NewScheduler(scope),
	}
	m.buildRules()
	return m
}

func (m *Monitor) Kickoff(ctx context.Context) {
	m.scheduler.Repeat(ctx, time.Minute, "Monitor", func() {
		if err := m.evalRules(ctx, time.Now()); err != nil {
			ezgo.LogCausesf(m.scope.GetLogger(), err, "evalRules")
		}
	})
}

func (m *Monitor) evalRules(ctx context.Context, now time.Time) error {
	for _, r := range m.rules {
		met, err := r.Eval(ctx, now)
		if err != nil {
			ezgo.LogCausesf(m.scope.GetLogger(), err, "Eval")
			continue
		}
		if met {
			m.notify(r, now)
		}
	}
	return nil
}

func (m *Monitor) notify(r *Rule[float64], now time.Time) {
}

func (m *Monitor) buildRules() error {
	var rules []*Rule[float64]
	for _, rc := range RULES_CONFIG {
		rules = append(rules, &Rule[float64]{
			query:     ezgo.Must(BuildQuery[float64](rc.queryType)),
			condition: ezgo.Must(BuildCondition[float64](rc.conditionType, rc.conditionArgs)),
		})
	}
	m.rules = rules
	return nil
}

func (m *Monitor) Terminate() {
	m.scheduler.Terminate()
}
