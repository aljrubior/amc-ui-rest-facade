//go:build wireinject
// +build wireinject

package wires

import (
	hybridHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/alerts/hybrid"
	hybridService "github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
	"github.com/google/wire"
)

func InitializeHybridAlertDatasource(hybridConfigClient config.HybridConfigClient) (alerts.Datasource, error) {

	wire.Build(
		hybridHttpClient.NewDefaultHttpClient,
		hybridService.NewDefaultService,
		hybrid.NewDefaultDatasource,
		wire.Bind(new(hybridHttpClient.HttpClient), new(hybridHttpClient.DefaultHttpClient)),
		wire.Bind(new(hybridService.Service), new(hybridService.DefaultService)),
		wire.Bind(new(alerts.Datasource), new(hybrid.DefaultDatasource)),
	)

	return hybrid.DefaultDatasource{}, nil
}
