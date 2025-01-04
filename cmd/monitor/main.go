package main

import (
	"context"
	"encoding/csv"
	"strings"
	"time"

	"github.com/holin20/catcatcat/internal/monitor"
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func main() {
	scope := ezgo.NewScopeWithDefaultLogger()
	defer scope.Close()

	ctx := context.Background()

	dpCsv := ezgo.SliceApply(
		[]float64{1, 2, 3, 4, 5, 6, 7},
		ezgo.FloatToString,
	)
	csvReader := csv.NewReader(strings.NewReader(strings.Join(dpCsv, "\n")))

	var ruleConfigs = []*monitor.RuleConfig{
		{"Macbook", monitor.FloatCsvQuery, []any{csvReader}, monitor.LessCondition, 1050.0},
		{"Face", monitor.FloatCsvQuery, []any{csvReader}, monitor.LessCondition, 100000.0},
	}

	monitor := monitor.NewMonitor(scope).
		WithRuleConfigs(ruleConfigs).
		WithEvalInterval(time.Second)

	monitor.Kickoff(ctx)

	time.Sleep(10 * time.Second)

	monitor.Terminate()
}
