//go:build wireinject
// +build wireinject

package wires

import (
	cloudhubHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub"
	fabricHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/fabric"
	hybridHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	applicationController "github.com/aljrubior/amc-ui-rest-facade/controllers/application"
	cloudhubService "github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
	fabricService "github.com/aljrubior/amc-ui-rest-facade/services/fabric"
	hybridService "github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
	"github.com/google/wire"
)

func InitializeApplicationController(
	hybridConfigClient config.HybridConfigClient,
	cloudhubConfigClient config.CloudhubConfigClient,
	fabricConfigClient config.FabricConfigClient) (applicationController.Controller, error) {

	wire.Build(
		hybridHttpClient.NewDefaultHttpClient,
		cloudhubHttpClient.NewDefaultHttpClient,
		fabricHttpClient.NewDefaultHttpClient,
		hybridService.NewDefaultService,
		cloudhubService.NewDefaultService,
		fabricService.NewDefaultService,
		applicationController.NewDefaultController,
		wire.Bind(new(hybridHttpClient.HttpClient), new(hybridHttpClient.DefaultHttpClient)),
		wire.Bind(new(cloudhubHttpClient.HttpClient), new(cloudhubHttpClient.DefaultHttpClient)),
		wire.Bind(new(fabricHttpClient.HttpClient), new(fabricHttpClient.DefaultHttpClient)),
		wire.Bind(new(hybridService.Service), new(hybridService.DefaultService)),
		wire.Bind(new(cloudhubService.Service), new(cloudhubService.DefaultService)),
		wire.Bind(new(fabricService.Service), new(fabricService.DefaultService)),
		wire.Bind(new(applicationController.Controller), new(applicationController.DefaultController)),
	)

	return applicationController.DefaultController{}, nil
}
