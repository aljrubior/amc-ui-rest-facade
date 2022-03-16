package requests

type AlertRequest struct {
	Name      string         `json:"name"`
	Severity  string         `json:"severity"`
	Condition AlertCondition `json:"condition"`
	Actions   []AlertAction  `json:"actions"`
}
