package application

import "github.com/aljrubior/amc-ui-rest-facade/wires"

func (t *App) initializeTargetDatasources() {

	datasource, err := wires.InitializeCloudhubTargetDatasource(cloudhubClientConfig)

	if err != nil {
		panic("Error: InitializeTargetDatasources dependency injection failed: Cloudhub")
	}

	t.targetDatasources = append(t.targetDatasources, datasource)

	hybridDatasource, err = wires.InitializeHybridTargetDatasource(hybridClientConfig)

	if err != nil {
		panic("Error: InitializeTargetDatasources dependency injection failed: Hybrid")
	}

	t.targetDatasources = append(t.targetDatasources, hybridDatasource)
}
