package alerts

type Response struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Severity       string         `json:"severity"`
	Enabled        bool           `json:"enabled"`
	OrganizationId string         `json:"organizationId"`
	EnvironmentId  string         `json:"environmentId"`
	LastModified   int64          `json:"lastModified"`
	CreatedAt      int64          `json:"createdAt"`
	IsSystem       bool           `json:"isSystem"`
	ProductName    string         `json:"productName"`
	ConditionType  string         `json:"conditionType"`
	Condition      AlertCondition `json:"condition"`
	Actions        []AlertAction  `json:"actions"`
}
