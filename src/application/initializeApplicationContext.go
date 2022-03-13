package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/wires"
)

var (
	hybridClientConfig   config.HybridConfigClient
	cloudhubClientConfig config.CloudhubConfigClient
	fabricClientConfig   config.FabricConfigClient
)

func (t *App) initializeApplicationContext() {

	hybridClientConfig = config.NewHybridConfigClient()
	cloudhubClientConfig = config.NewCloudhubConfigClient()
	fabricClientConfig = config.NewFabricConfigClient()

	var err error

	t.applicationController, err = wires.InitializeApplicationController(hybridClientConfig, cloudhubClientConfig, fabricClientConfig)

	if err != nil {
		panic("Error: ApplicationController dependency injection failed")
	}
}
