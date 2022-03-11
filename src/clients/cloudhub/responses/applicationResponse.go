package responses

type ApplicationResponse struct {
	VersionId                         string            `json:"versionId"`
	Domain                            string            `json:"domain"`
	FullDomain                        string            `json:"fullDomain"`
	Properties                        map[string]string `json:"properties"`
	PropertiesOptions                 map[string]string `json:"propertiesOptions"`
	Status                            string            `json:"status"`
	Worker                            Worker            `json:"worker"`
	LastUpdateTime                    int64             `json:"lastUpdateTime"`
	FileName                          string            `json:"fileName"`
	MuleVersion                       MuleVersion       `json:"muleVersion"`
	PersistentQueues                  bool              `json:"persistentQueues"`
	PersistentQueuesEncryptionEnabled bool              `json:"persistentQueuesEncryptionEnabled"`
	PersistentQueuesEncrypted         bool              `json:"persistentQueuesEncrypted"`
	MonitoringEnabled                 bool              `json:"monitoringEnabled"`
	MonitoringAutoRestart             bool              `json:"monitoringAutoRestart"`
	StaticIPsEnabled                  bool              `json:"staticIPsEnabled"`
	HasFile                           bool              `json:"hasFile"`
	SecureDataGatewayEnabled          bool              `json:"secureDataGatewayEnabled"`
	LoggingNgEnabled                  bool              `json:"loggingNgEnabled"`
	LoggingCustomLog4JEnabled         bool              `json:"loggingCustomLog4JEnabled"`
	CloudObjectStoreRegion            string            `json:"cloudObjectStoreRegion"`
	InsightsReplayDataRegion          string            `json:"insightsReplayDataRegion"`
	IsDeploymentWaiting               bool              `json:"isDeploymentWaiting"`
	DeploymentGroup                   DeploymentGroup   `json:"deploymentGroup"`
	UpdateRuntimeConfig               bool              `json:"updateRuntimeConfig"`
	LogLevels                         []LogLevel        `json:"logLevels"`
	IPAddresses                       []IPAddress       `json:"ipAddresses"`
}
