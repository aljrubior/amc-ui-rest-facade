package application

import (
	"github.com/aljrubior/go-facade/config"
	"github.com/aljrubior/go-facade/datasources/targets"
)

var (
	hybridClientConfig                  config.HybridConfigClient
	cloudhubClientConfig                config.CloudhubConfigClient
	fabricClientConfig                  config.FabricConfigClient
	runtimeFabricManagementClientConfig config.RuntimeFabricManagementClientConfig
	coreServiceConfigClient             config.AccessManagementConfigClient

	cloudhubDatasource targets.Datasource
	hybridDatasource   targets.Datasource
)
