//go:build wireinject
// +build wireinject

package wires

import (
	hybridHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets/hybrid"
	hybridService "github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
	"github.com/google/wire"
)

func InitializeHybridTargetDatasource(hybridConfigClient config.HybridConfigClient) (targets.Datasource, error) {

	wire.Build(
		hybridHttpClient.NewDefaultHttpClient,
		hybridService.NewDefaultService,
		hybrid.NewDatasource,
		wire.Bind(new(hybridHttpClient.HttpClient), new(hybridHttpClient.DefaultHttpClient)),
		wire.Bind(new(hybridService.Service), new(hybridService.DefaultService)),
		wire.Bind(new(targets.Datasource), new(hybrid.Datasource)),
	)

	return hybrid.Datasource{}, nil
}
