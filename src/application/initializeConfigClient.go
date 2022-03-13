package application

import "github.com/aljrubior/amc-ui-rest-facade/config"

var (
	hybridClientConfig   config.HybridConfigClient
	cloudhubClientConfig config.CloudhubConfigClient
	fabricClientConfig   config.FabricConfigClient
)

func (t *App) initializeConfigClient() {

	hybridClientConfig = config.NewHybridConfigClient()
	cloudhubClientConfig = config.NewCloudhubConfigClient()
	fabricClientConfig = config.NewFabricConfigClient()
}
