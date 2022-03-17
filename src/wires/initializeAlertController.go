//go:build wireinject
// +build wireinject

package wires

import (
	cloudhubHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub"
	hybridHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	alertController "github.com/aljrubior/amc-ui-rest-facade/controllers/alert"
	"github.com/aljrubior/amc-ui-rest-facade/services/alert"
	"github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
	"github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
	"github.com/google/wire"
)

func InitializeAlertController(hybridConfigClient config.HybridConfigClient, configClient config.CloudhubConfigClient) (alertController.Controller, error) {

	wire.Build(
		hybridHttpClient.NewDefaultHttpClient,
		cloudhubHttpClient.NewDefaultHttpClient,
		hybrid.NewDefaultService,
		cloudhub.NewDefaultService,
		alert.NewDefaultService,
		alertController.NewDefaultController,
		wire.Bind(new(hybridHttpClient.HttpClient), new(hybridHttpClient.DefaultHttpClient)),
		wire.Bind(new(cloudhubHttpClient.HttpClient), new(cloudhubHttpClient.DefaultHttpClient)),
		wire.Bind(new(hybrid.Service), new(hybrid.DefaultService)),
		wire.Bind(new(cloudhub.Service), new(cloudhub.DefaultService)),
		wire.Bind(new(alert.Service), new(alert.DefaultService)),
		wire.Bind(new(alertController.Controller), new(alertController.DefaultController)),
	)

	return alertController.DefaultController{}, nil
}
