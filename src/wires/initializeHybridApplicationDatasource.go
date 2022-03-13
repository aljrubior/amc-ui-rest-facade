//go:build wireinject
// +build wireinject

package wires

import (
	hybridHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications"
	hybridService "github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
	"github.com/google/wire"
)

func initializeHybridApplicationDatasource(hybridConfigClient config.HybridConfigClient) (datasources.ApplicationDatasource, error) {

	wire.Build(
		hybridHttpClient.NewDefaultHttpClient,
		hybridService.NewDefaultService,
		applications.NewHybridApplicationDatasource,
		wire.Bind(new(hybridHttpClient.HttpClient), new(hybridHttpClient.DefaultHttpClient)),
		wire.Bind(new(hybridService.Service), new(hybridService.DefaultService)),
		wire.Bind(new(datasources.ApplicationDatasource), new(applications.HybridApplicationDatasource)),
	)

	return applications.HybridApplicationDatasource{}, nil
}
