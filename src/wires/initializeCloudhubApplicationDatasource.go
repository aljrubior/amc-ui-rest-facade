//go:build wireinject
// +build wireinject

package wires

import (
	cloudhubHttpClient "github.com/aljrubior/go-facade/clients/cloudhub"
	"github.com/aljrubior/go-facade/config"
	"github.com/aljrubior/go-facade/datasources/applications"
	"github.com/aljrubior/go-facade/datasources/applications/cloudhub"
	cloudhubService "github.com/aljrubior/go-facade/services/cloudhub"
	"github.com/google/wire"
)

func InitializeCloudhubApplicationDatasource(configClient config.CloudhubConfigClient) (applications.Datasource, error) {

	wire.Build(
		cloudhubHttpClient.NewDefaultHttpClient,
		cloudhubService.NewDefaultService,
		cloudhub.NewDefaultDatasource,
		wire.Bind(new(cloudhubHttpClient.HttpClient), new(cloudhubHttpClient.DefaultHttpClient)),
		wire.Bind(new(cloudhubService.Service), new(cloudhubService.DefaultService)),
		wire.Bind(new(applications.Datasource), new(cloudhub.DefaultDatasource)),
	)

	return cloudhub.DefaultDatasource{}, nil
}
