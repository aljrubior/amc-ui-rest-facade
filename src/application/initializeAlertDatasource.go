package application

import (
	"github.com/aljrubior/go-facade/datasources/alerts"
	"github.com/aljrubior/go-facade/wires"
)

func (t *App) initializeAlertDatasources() {

	t.alertDatasources = make(map[string]alerts.Datasource)

	datasource, err := wires.InitializeHybridAlertDatasource(hybridClientConfig)

	if err != nil {
		panic("Error: InitializeHybridAlertDatasource dependency injection failed")
	}

	t.alertDatasources[HYBRID_PRODUCT_NAME] = datasource

	datasource, err = wires.InitializeCloudhubAlertDatasource(cloudhubClientConfig)

	if err != nil {
		panic("Error: InitializeCloudhubAlertDatasource dependency injection failed")
	}

	t.alertDatasources[CLOUDHUB_PRODUCT_NAME] = datasource
}
