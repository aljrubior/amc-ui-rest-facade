//go:build wireinject
// +build wireinject

package wires

import (
	alertController "github.com/aljrubior/amc-ui-rest-facade/controllers/alert"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/services/alert"
	"github.com/google/wire"
)

func InitializeAlertController(datasources map[string]alerts.Datasource) (alertController.Controller, error) {

	wire.Build(
		alert.NewDefaultService,
		alertController.NewDefaultController,
		wire.Bind(new(alert.Service), new(alert.DefaultService)),
		wire.Bind(new(alertController.Controller), new(alertController.DefaultController)),
	)

	return alertController.DefaultController{}, nil
}
