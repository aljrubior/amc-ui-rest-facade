//go:build wireinject
// +build wireinject

package wires

import (
	alertController "github.com/aljrubior/go-facade/controllers/alert"
	"github.com/aljrubior/go-facade/datasources/alerts"
	"github.com/aljrubior/go-facade/services/alert"
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
