//go:build wireinject
// +build wireinject

package wires

import (
	fabricHttpClient "github.com/aljrubior/go-facade/clients/fabric"
	"github.com/aljrubior/go-facade/clients/runtimeFabricManagement"
	"github.com/aljrubior/go-facade/config"
	"github.com/aljrubior/go-facade/datasources/applications"
	"github.com/aljrubior/go-facade/datasources/applications/fabric"
	fabricService "github.com/aljrubior/go-facade/services/fabric"
	runtimeFabricManagementService "github.com/aljrubior/go-facade/services/runtimeFabricManagement"
	"github.com/google/wire"
)

func InitializeFabricApplicationDatasource(
	fabricConfigClient config.FabricConfigClient,
	runtimeFabricManagementClientConfig config.RuntimeFabricManagementClientConfig) (applications.Datasource, error) {

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
		wire.Bind(new(applications.Datasource), new(fabric.Datasource)),
	)

	return fabric.Datasource{}, nil
}
