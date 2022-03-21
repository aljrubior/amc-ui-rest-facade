package config

import (
	"fmt"
	"github.com/aljrubior/go-facade/utils/getenvs"
)

func NewAtlasConfigClient() AtlasConfigClient {

	basePath := getenvs.GetString("ATLAS_REST_PATH", "/api/v1")

	return AtlasConfigClient{
		Protocol:    getenvs.GetString("ATLAS_REST_PROTOCOL", "http"),
		Host:        getenvs.GetString("ATLAS_REST_HOST", ""),
		Port:        getenvs.GetString("ATLAS_REST_PORT", "8080"),
		TargetsPath: fmt.Sprintf("%s%s", basePath, "/organizations/%s/providers/%s/targets"),
	}
}

type AtlasConfigClient struct {
	Protocol,
	Host,
	Port,
	TargetsPath string
}
