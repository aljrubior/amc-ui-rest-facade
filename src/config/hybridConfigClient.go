package config

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/utils/getenvs"
)

func NewHybridConfigClient() HybridConfigClient {

	basePath := getenvs.GetString("AMC_REST_PATH", "/api/v1")

	return HybridConfigClient{
		Protocol:                 getenvs.GetString("AMC_REST_PROTOCOL", "http"),
		Host:                     getenvs.GetString("AMC_REST_HOST", ""),
		Port:                     getenvs.GetString("AMC_REST_PORT", "8080"),
		AlertsPath:               fmt.Sprintf("%s%s", basePath, "/alerts"),
		AlertPath:                fmt.Sprintf("%s%s", basePath, "/alerts/%s"),
		AlertHistoryPath:         fmt.Sprintf("%s%s", basePath, "/alerts/%s/history"),
		ResourceAlertHistoryPath: fmt.Sprintf("%s%s", basePath, "/alerts/resource/%s/history"),
		ApplicationsPath:         fmt.Sprintf("%s%s", basePath, "/applications"),
		ServerGroupsPath:         fmt.Sprintf("%s%s", basePath, "/serverGroups"),
		ClustersPath:             fmt.Sprintf("%s%s", basePath, "/clusters"),
		ServersPath:              fmt.Sprintf("%s%s", basePath, "/servers"),
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
	ServerGroupsPath,
	ClustersPath,
	ServersPath string
}
