package responses

type AlertAction struct {
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Type    string   `json:"type"`
	Content string   `json:"content"`
}
