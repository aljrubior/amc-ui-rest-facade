package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/wires"
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
}