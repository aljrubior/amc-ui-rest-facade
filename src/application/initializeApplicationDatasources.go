package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications"
	"github.com/aljrubior/amc-ui-rest-facade/wires"
)

func (t *App) initializeApplicationDatasources() {

	t.applicationDatasources = make(map[string]applications.Datasource)

	datasource, err := wires.InitializeHybridApplicationDatasource(hybridClientConfig)

	if err != nil {
		panic("Error: initializeApplicationContext dependency injection failed")
	}

	t.applicationDatasources[HYBRID_PRODUCT_NAME] = datasource

	datasource, err = wires.InitializeCloudhubApplicationDatasource(cloudhubClientConfig)

	if err != nil {
		panic("Error: initializeApplicationContext dependency injection failed")
	}

	t.applicationDatasources[CLOUDHUB_PRODUCT_NAME] = datasource

	datasource, err = wires.InitializeFabricApplicationDatasource(fabricClientConfig, runtimeFabricManagementClientConfig)

	if err != nil {
		panic("Error: initializeApplicationContext dependency injection failed")
	}

	t.applicationDatasources[RUNTIME_FABRIC_PRODUCT_NAME] = datasource
}
