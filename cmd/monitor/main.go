package main

import (
	"context"
	"time"

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

	var cdpRuleConfig = []*monitor.CdpRuleConfig{
		{
			Name:                "Macbook air 13 price\"",
			CatId:               "0",
			MonitorField:        monitor.CdpPrice,
			ConditionType:       monitor.LessCondition,
			ConditionArg:        1100.0,
			AlertCriteria:       "price < $1100",
			QueryResultTemplate: "current price is $%f",
		},
		{
			Name:                "Macbook air 13 availability\"",
			CatId:               "0",
			MonitorField:        monitor.CdpInStock,
			ConditionType:       monitor.EqualCondition,
			ConditionArg:        1.0,
			AlertCriteria:       "in stock",
			QueryResultTemplate: "current in-stock $%f",
		},
		{
			Name:                "Pretty face availability",
			CatId:               "1",
			MonitorField:        monitor.CdpInStock,
			ConditionType:       monitor.EqualCondition,
			ConditionArg:        1.0,
			AlertCriteria:       "in stock",
			QueryResultTemplate: "in-stock status: %f",
		},
		{
			Name:                "Macbook pro 14 price\"",
			CatId:               "2",
			MonitorField:        monitor.CdpPrice,
			ConditionType:       monitor.LessCondition,
			ConditionArg:        1000.0,
			AlertCriteria:       "price < $1000",
			QueryResultTemplate: "current price is $%f",
		},
		{
			Name:                "Macbook pro 14 availability\"",
			CatId:               "2",
			MonitorField:        monitor.CdpInStock,
			ConditionType:       monitor.EqualCondition,
			ConditionArg:        1.0,
			AlertCriteria:       "in stock",
			QueryResultTemplate: "in-stock $%f",
		},
	}

	monitor := monitor.NewMonitor(scope, db).
		WithCdpRuleConfigs(cdpRuleConfig).
		WithEvalInterval(5 * time.Minute)

	monitor.Kickoff(ctx)

	time.Sleep(24 * time.Hour * 7)

	monitor.Terminate()
}
