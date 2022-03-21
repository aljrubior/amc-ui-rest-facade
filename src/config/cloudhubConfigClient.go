package config

import (
	"fmt"
	"github.com/aljrubior/go-facade/utils/getenvs"
)

func NewCloudhubConfigClient() CloudhubConfigClient {

	basePath := getenvs.GetString("CLOUDHUB_REST_V2_PATH", "/api/v2")

	return CloudhubConfigClient{
		Protocol:         getenvs.GetString("CLOUDHUB_REST_PROTOCOL", "http"),
		Host:             getenvs.GetString("CLOUDHUB_REST_HOST", ""),
		Port:             getenvs.GetString("CLOUDHUB_REST_PORT", "8080"),
		AlertsPath:       fmt.Sprintf("%s%s", basePath, "/alerts"),
		AlertPath:        fmt.Sprintf("%s%s", basePath, "/alerts/%s"),
		AlertHistoryPath: fmt.Sprintf("%s%s", basePath, "/alerts/%s/history"),
		ApplicationsPath: fmt.Sprintf("%s%s", basePath, "/applications"),
	}
}

type CloudhubConfigClient struct {
	Protocol,
	Host,
	Port,
	AlertsPath,
	AlertPath,
	AlertHistoryPath,
	ApplicationsPath string
}
