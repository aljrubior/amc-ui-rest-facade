package config

import (
	"fmt"
	"github.com/aljrubior/go-facade/utils/getenvs"
)

func NewAccessManagementConfigClient() AccessManagementConfigClient {

	basePath := getenvs.GetString("CORESERVICES_REST_PATH", "/api")

	return AccessManagementConfigClient{
		Protocol:      getenvs.GetString("CORESERVICES_REST_PROTOCOL", "http"),
		Host:          getenvs.GetString("CORESERVICES_REST_HOST", ""),
		Port:          getenvs.GetString("CORESERVICES_REST_PORT", "3004"),
		MembersPath:   fmt.Sprintf("%s%s", basePath, "/organizations/%s/members"),
		AuthorizePath: fmt.Sprintf("%s%s", basePath, "/authorize"),
	}
}

type AccessManagementConfigClient struct {
	Protocol,
	Host,
	Port,
	MembersPath,
	AuthorizePath string
}
