package monitor

type RuleConfig struct {
	RuleId        string
	Name          string
	QueryType     QueryType
	QueryArgs     []any
	ConditionType ConditionType
	ConditionArgs float64

	// for notification content
	AlertCriteria       string
	QueryResultTemplate string // result will be provided as %f
}
