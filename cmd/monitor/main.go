package main

import (
	"context"
	"time"

	"github.com/holin20/catcatcat/internal/example"
	"github.com/holin20/catcatcat/internal/monitor"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	scope := ezgo.NewScopeWithDefaultLogger()
	defer scope.Close()

	ctx := context.Background()

	var ruleConfigs = []*monitor.RuleConfig{
		{
			Name:          example.CATS[0].Name,
			QueryType:     monitor.ZapTail,
			QueryArgs:     []any{"logs/cdp_0.txt", "ts", "price"},
			ConditionType: monitor.LessCondition,
			ConditionArgs: 1050.0,
		},
		{
			Name:          example.CATS[1].Name,
			QueryType:     monitor.ZapTail,
			QueryArgs:     []any{"logs/cdp_1.txt", "ts", "inStock"},
			ConditionType: monitor.EqualCondition,
			ConditionArgs: 1,
		},
	}

	monitor := monitor.NewMonitor(scope).
		WithRuleConfigs(ruleConfigs).
		WithEvalInterval(time.Minute)

	monitor.Kickoff(ctx)

	time.Sleep(10 * time.Second)

	monitor.Terminate()
}
