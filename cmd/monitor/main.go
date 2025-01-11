package main

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/internal/example"
	"github.com/holin20/catcatcat/internal/monitor"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	scope := ezgo.NewScopeWithDefaultLogger("Monitor")
	defer scope.Close()

	ctx := context.Background()

	var ruleConfigs = map[string]*monitor.RuleConfig{
		"0": {
			RuleId:        "0",
			Name:          example.CATS[0].Name,
			QueryType:     monitor.ZapTail,
			QueryArgs:     []any{"logs/cdp_0.txt", "ts", "price"},
			ConditionType: monitor.LessCondition,
			ConditionArgs: 1100.0,

			AlertCriteria:       "price < $1100.0",
			QueryResultTemplate: "current price is $%f",
		},
		"1": {
			RuleId:        "1",
			Name:          example.CATS[1].Name,
			QueryType:     monitor.ZapTail,
			QueryArgs:     []any{"logs/cdp_1.txt", "ts", "inStock"},
			ConditionType: monitor.EqualCondition,
			ConditionArgs: 1,

			AlertCriteria:       "in stock",
			QueryResultTemplate: "in-stock status: %f",
		},
	}

	monitor := monitor.NewMonitor(scope).
		WithRuleConfigs(ruleConfigs).
		WithEvalInterval(time.Minute)

	monitor.Kickoff(ctx)

	time.Sleep(24 * time.Hour)

	monitor.Terminate()
}
