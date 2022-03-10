package alerts

type ResourceAlertHistory struct {
	AlertId     string                `json:"alertId"`
	AlertName   string                `json:"alertName"`
	TriggeredAt int64                 `json:"triggeredAt"`
	Severity    string                `json:"severity"`
	Condition   AlertHistoryCondition `json:"condition"`
}