package application

import (
	"github.com/aljrubior/go-facade/datasources/targets"
	"github.com/aljrubior/go-facade/wires"
)

func (t *App) initializeControllers() {

	var err error

	t.applicationController, err = wires.InitializeApplicationController(t.applicationDatasources)

	if err != nil {
		panic("Error: Application Controller dependency injection failed")
	}

	t.targetController, err = wires.InitializeTargetController(t.targetDatasources)

	if err != nil {
		panic("Error: Target Controller dependency injection failed")
	}

	t.serverController, err = wires.InitializeServerController([]targets.Datasource{hybridDatasource})

	if err != nil {
		panic("Error: Server Controller dependency injection failed")
	}

	t.alertController, err = wires.InitializeAlertController(t.alertDatasources)

	if err != nil {
		panic("Error: Alert Controller dependency injection failed")
	}

	t.userController, err = wires.InitializeUserController(coreServiceConfigClient)

	if err != nil {
		panic("Error: User Controller dependency injection failed")
	}

	t.permissionController, err = wires.InitializePermissionController(t.permissionDatasource)

	if err != nil {
		panic("Error: Permission Controller dependency injection failed")
	}
}
