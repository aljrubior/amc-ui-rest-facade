//go:build wireinject
// +build wireinject

package wires

import (
	cloudhubHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications"
	cloudhubService "github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
	"github.com/google/wire"
)

func initializeCloudhubApplicationDatasource(configClient config.CloudhubConfigClient) (datasources.ApplicationDatasource, error) {

	wire.Build(
		cloudhubHttpClient.NewDefaultHttpClient,
		cloudhubService.NewDefaultService,
		applications.NewCloudhubApplicationDatasource,
		wire.Bind(new(cloudhubHttpClient.HttpClient), new(cloudhubHttpClient.DefaultHttpClient)),
		wire.Bind(new(cloudhubService.Service), new(cloudhubService.DefaultService)),
		wire.Bind(new(datasources.ApplicationDatasource), new(applications.CloudhubApplicationDatasource)),
	)

	return applications.CloudhubApplicationDatasource{}, nil
}
