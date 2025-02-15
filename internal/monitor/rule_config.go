package monitor

type CdpRuleConfig struct {
	RuleId        string
	Name          string
	CatId         string
	MonitorField  CdpField
	ConditionType ConditionType
	ConditionArg  float64

	AlertCriteria       string
	QueryResultTemplate string
}
