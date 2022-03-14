//go:build wireinject
// +build wireinject

package wires

import (
	hybridHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/hybrid"
	hybridService "github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
	"github.com/google/wire"
)

func InitializeHybridApplicationDatasource(hybridConfigClient config.HybridConfigClient) (datasources.ApplicationDatasource, error) {

	wire.Build(
		hybridHttpClient.NewDefaultHttpClient,
		hybridService.NewDefaultService,
		hybrid.NewDatasource,
		wire.Bind(new(hybridHttpClient.HttpClient), new(hybridHttpClient.DefaultHttpClient)),
		wire.Bind(new(hybridService.Service), new(hybridService.DefaultService)),
		wire.Bind(new(datasources.ApplicationDatasource), new(hybrid.Datasource)),
	)

	return hybrid.Datasource{}, nil
}
