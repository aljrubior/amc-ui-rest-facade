package config

func NewApplicationManagerV1ConfigClient() ApplicationManagerV1ConfigClient {
	return ApplicationManagerV1ConfigClient{
		Protocol:   "http",
		Host:       "application-manager-service.arm.svc",
		Port:       "8080",
		Path:       "/api/v1",
		AlertsPath: "/alerts",
		AlertPath:  "/alerts/%s",
	}
}

type ApplicationManagerV1ConfigClient struct {
	Protocol,
	Host,
	Port,
	Path,
	AlertsPath string
	AlertPath string
}
