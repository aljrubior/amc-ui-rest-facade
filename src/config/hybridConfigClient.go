package config

func NewHybridConfigClient() HybridConfigClient {
	return HybridConfigClient{
		Protocol:                 "http",
		Host:                     "application-manager-service.arm.svc",
		Port:                     "8080",
		AlertsPath:               "/api/v1/alerts",
		AlertPath:                "/api/v1/alerts/%s",
		AlertHistoryPath:         "/api/v1/alerts/%s/history",
		ResourceAlertHistoryPath: "/api/v1/alerts/resource/%s/history",
		ApplicationsPath:         "/api/v1/applications",
		ServerGroupsPath:         "/api/v1/serverGroups",
	}
}

type HybridConfigClient struct {
	Protocol,
	Host,
	Port,
	AlertsPath,
	AlertPath,
	AlertHistoryPath,
	ResourceAlertHistoryPath,
	ApplicationsPath,
	ServerGroupsPath string
}
