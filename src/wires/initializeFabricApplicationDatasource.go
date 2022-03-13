//go:build wireinject
// +build wireinject

package wires

import (
	fabricHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/fabric"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications"
	fabricService "github.com/aljrubior/amc-ui-rest-facade/services/fabric"
	"github.com/google/wire"
)

func InitializeFabricApplicationDatasource(configClient config.FabricConfigClient) (datasources.ApplicationDatasource, error) {

	wire.Build(
		fabricHttpClient.NewDefaultHttpClient,
		fabricService.NewDefaultService,
		applications.NewFabricApplicationDatasource,
		wire.Bind(new(fabricHttpClient.HttpClient), new(fabricHttpClient.DefaultHttpClient)),
		wire.Bind(new(fabricService.Service), new(fabricService.DefaultService)),
		wire.Bind(new(datasources.ApplicationDatasource), new(applications.FabricApplicationDatasource)),
	)

	return applications.FabricApplicationDatasource{}, nil
}
