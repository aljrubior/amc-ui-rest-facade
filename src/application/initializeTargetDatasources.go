package application

import "github.com/aljrubior/amc-ui-rest-facade/wires"

func (t *App) initializeTargetDatasources() {

	datasource, err := wires.InitializeCloudhubTargetDatasource(cloudhubClientConfig)

	if err != nil {
		panic("Error: initializeApplicationContext dependency injection failed")
	}

	t.targetDatasources = append(t.targetDatasources, datasource)

}
