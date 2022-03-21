package alerts

type AlertHistoryResponse struct {
	TriggeredAt int64                 `json:"triggeredAt"`
	Context     AlertHistoryContext   `json:"context"`
	Severity    string                `json:"severity"`
	Condition   AlertHistoryCondition `json:"condition"`
	Action      []AlertHistoryAction  `json:"actions"`
}
