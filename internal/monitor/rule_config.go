package monitor

type RuleConfig struct {
	Name          string
	QueryType     QueryType
	QueryArgs     []any
	ConditionType ConditionType
	ConditionArgs float64
}
