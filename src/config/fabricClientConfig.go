package config

func NewFabricConfigClient() FabricConfigClient {
	return FabricConfigClient{
		Protocol:        "http",
		Host:            "application-manager-service.arm.svc",
		Port:            "8080",
		DeploymentsPath: "/api/v2/organizations/%s/environments/%s/deployments",
	}
}

type FabricConfigClient struct {
	Protocol,
	Host,
	Port,
	DeploymentsPath string
}
