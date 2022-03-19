package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets"
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
