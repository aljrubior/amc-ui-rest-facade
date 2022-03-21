package alerts

type AlertHistoryContext struct {
	Name     string `json:"name"`
	Resource string `json:"resource"`
	User     string `json:"user,omitempty"`
	App      string `json:"app,omitempty"`
}
