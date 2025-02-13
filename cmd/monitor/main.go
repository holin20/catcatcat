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

	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "NewLocalPostgresDB")
	defer db.Close()

	var ruleConfigs = map[string]*monitor.RuleConfig{
		"0": {
			RuleId:        "0",
			Name:          example.CATS[0].Name,
			QueryType:     monitor.EntCdp,
			QueryArgs:     []any{db, "0", monitor.CdpPrice},
			ConditionType: monitor.LessCondition,
			ConditionArgs: 1100.0,

			AlertCriteria:       "price < $1100.0",
			QueryResultTemplate: "current price is $%f",
		},
		"1": {
			RuleId:        "1",
			Name:          example.CATS[1].Name,
			QueryType:     monitor.EntCdp,
			QueryArgs:     []any{db, "1", monitor.CdpInStock},
			ConditionType: monitor.EqualCondition,
			ConditionArgs: 1,

			AlertCriteria:       "in stock",
			QueryResultTemplate: "in-stock status: %f",
		},
		"2": {
			RuleId:        "2",
			Name:          example.CATS[2].Name,
			QueryType:     monitor.EntCdp,
			QueryArgs:     []any{db, "2", monitor.CdpPrice},
			ConditionType: monitor.LessCondition,
			ConditionArgs: 1000,

			AlertCriteria:       "price < $1000",
			QueryResultTemplate: "current price is $%f",
		},
	}

	monitor := monitor.NewMonitor(scope, db).
		WithRuleConfigs(ruleConfigs).
		WithEvalInterval(5 * time.Minute)

	monitor.Kickoff(ctx)

	time.Sleep(24 * time.Hour * 7)

	monitor.Terminate()
}
