package application

import "github.com/aljrubior/amc-ui-rest-facade/wires"

func (t *App) initializePermissionDatasource() {

	datasource, err := wires.InitializePermissionDatasource(coreServiceConfigClient)

	if err != nil {
		panic("Error: InitializePermissionDatasource dependency injection failed: Cloudhub")
	}

	t.permissionDatasource = datasource
}
