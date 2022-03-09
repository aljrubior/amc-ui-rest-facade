package config

func NewApplicationManagerV1ConfigClient() ApplicationManagerV1ConfigClient {
	return ApplicationManagerV1ConfigClient{
		Protocol:   "http",
		Host:       "application-manager-service.arm.svc",
		Port:       "8080",
		AlertsPath: "/api/v1/alerts",
		AlertPath:  "/api/v1/alerts/%s",
	}
}

type ApplicationManagerV1ConfigClient struct {
	Protocol,
	Host,
	Port,
	AlertsPath string
	AlertPath string
}
