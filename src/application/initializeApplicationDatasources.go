package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/wires"
)

func (t *App) initializeApplicationDatasources() {

	datasource, err := wires.InitializeHybridApplicationDatasource(hybridClientConfig)

	if err != nil {
		panic("Error: initializeApplicationContext dependency injection failed")
	}

	t.applicationDatasources = append(t.applicationDatasources, datasource)

	datasource, err = wires.InitializeCloudhubApplicationDatasource(cloudhubClientConfig)

	if err != nil {
		panic("Error: initializeApplicationContext dependency injection failed")
	}

	t.applicationDatasources = append(t.applicationDatasources, datasource)

	datasource, err = wires.InitializeFabricApplicationDatasource(fabricClientConfig, runtimeFabricManagementClientConfig)

	if err != nil {
		panic("Error: initializeApplicationContext dependency injection failed")
	}

	t.applicationDatasources = append(t.applicationDatasources, datasource)
}
