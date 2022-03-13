package config

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/utils/getenvs"
)

func NewFabricConfigClient() FabricConfigClient {

	basePath := getenvs.GetString("AM_REST_PATH", "/api/v2")

	return FabricConfigClient{
		Protocol:        getenvs.GetString("AM_REST_PROTOCOL", "http"),
		Host:            getenvs.GetString("AM_REST_HOST", ""),
		Port:            getenvs.GetString("AM_REST_PORT", "8080"),
		DeploymentsPath: fmt.Sprintf("%s%s", basePath, "/api/v2/organizations/%s/environments/%s/deployments"),
	}
}

type FabricConfigClient struct {
	Protocol,
	Host,
	Port,
	DeploymentsPath string
}
