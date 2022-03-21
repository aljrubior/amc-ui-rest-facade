package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets"
	"github.com/aljrubior/amc-ui-rest-facade/wires"
)

func (t *App) initializeTargetDatasources() {

	t.targetDatasources = make(map[string]targets.Datasource)

	datasource, err := wires.InitializeCloudhubTargetDatasource(cloudhubClientConfig)

	if err != nil {
		panic("Error: InitializeTargetDatasources dependency injection failed: Cloudhub")
	}

	t.targetDatasources[CLOUDHUB_PRODUCT_NAME] = datasource

	datasource, err = wires.InitializeHybridTargetDatasource(hybridClientConfig)

	if err != nil {
		panic("Error: InitializeTargetDatasources dependency injection failed: Hybrid")
	}

	t.targetDatasources[HYBRID_PRODUCT_NAME] = datasource
}
