package config

func NewCloudhubConfigClient() CloudhubConfigClient {
	return CloudhubConfigClient{
		Protocol:   "http",
		Host:       "console.qax.msap.io",
		Port:       "8080",
		Path:       "/api/v2",
		AlertsPath: "/alerts",
	}
}

type CloudhubConfigClient struct {
	Protocol,
	Host,
	Port,
	Path,
	AlertsPath string
}
