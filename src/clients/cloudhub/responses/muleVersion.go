package responses

type MuleVersion struct {
	Version          string `json:"version"`
	UpdateId         string `json:"updateId"`
	LatestUpdateId   string `json:"latestUpdateId"`
	EndOfSupportDate int64  `json:"endOfSupportDate"`
}
