package config

func NewCloudhubConfigClient() CloudhubConfigClient {
	return CloudhubConfigClient{
		Protocol:   "http",
		Host:       "console.qax.msap.io",
		Port:       "8080",
		AlertsPath: "/api/v2/alerts",
		AlertPath:  "/api/v2/alerts/%s",
	}
}

type CloudhubConfigClient struct {
	Protocol,
	Host,
	Port,
	AlertsPath,
	AlertPath string
}
