//go:build wireinject
// +build wireinject

package wires

import (
	fabricHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/fabric"
	"github.com/aljrubior/amc-ui-rest-facade/clients/runtimeFabricManagement"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/fabric"
	fabricService "github.com/aljrubior/amc-ui-rest-facade/services/fabric"
	runtimeFabricManagementService "github.com/aljrubior/amc-ui-rest-facade/services/runtimeFabricManagement"
	"github.com/google/wire"
)

func InitializeFabricApplicationDatasource(
	fabricConfigClient config.FabricConfigClient,
	runtimeFabricManagementClientConfig config.RuntimeFabricManagementClientConfig) (datasources.ApplicationDatasource, error) {

	wire.Build(
		fabricHttpClient.NewDefaultHttpClient,
		runtimeFabricManagement.NewDefaultHttpClient,
		fabricService.NewDefaultService,
		runtimeFabricManagementService.NewDefaultService,
		fabric.NewDatasource,
		wire.Bind(new(fabricHttpClient.HttpClient), new(fabricHttpClient.DefaultHttpClient)),
		wire.Bind(new(runtimeFabricManagement.HttpClient), new(runtimeFabricManagement.DefaultHttpClient)),
		wire.Bind(new(fabricService.Service), new(fabricService.DefaultService)),
		wire.Bind(new(runtimeFabricManagementService.Service), new(runtimeFabricManagementService.DefaultService)),
		wire.Bind(new(datasources.ApplicationDatasource), new(fabric.Datasource)),
	)

	return fabric.Datasource{}, nil
}
