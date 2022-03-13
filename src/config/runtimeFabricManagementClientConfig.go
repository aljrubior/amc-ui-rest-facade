package config

import (
	"github.com/aljrubior/amc-ui-rest-facade/utils/getenvs"
)

func NewRuntimeFabricManagementClientConfig() RuntimeFabricManagementClientConfig {

	return RuntimeFabricManagementClientConfig{
		Protocol:    getenvs.GetString("RF_REST_PROTOCOL", "http"),
		Host:        getenvs.GetString("RF_REST_HOST", ""),
		Port:        getenvs.GetString("RF_REST_PORT", "80"),
		TargetsPath: "/organizations/%s/targets",
	}
}

type RuntimeFabricManagementClientConfig struct {
	Protocol,
	Host,
	Port,
	TargetsPath string
}
