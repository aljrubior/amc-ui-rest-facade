package application

import "github.com/aljrubior/amc-ui-rest-facade/config"

func (t *App) initializeConfigClient() {

	hybridClientConfig = config.NewHybridConfigClient()
	cloudhubClientConfig = config.NewCloudhubConfigClient()
	fabricClientConfig = config.NewFabricConfigClient()
	runtimeFabricManagementClientConfig = config.NewRuntimeFabricManagementClientConfig()
}
