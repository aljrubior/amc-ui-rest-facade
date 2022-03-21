//go:build wireinject
// +build wireinject

package wires

import (
	cloudhubHttpClient "github.com/aljrubior/go-facade/clients/cloudhub"
	"github.com/aljrubior/go-facade/config"
	"github.com/aljrubior/go-facade/datasources/targets"
	"github.com/aljrubior/go-facade/datasources/targets/cloudhub"
	cloudhubService "github.com/aljrubior/go-facade/services/cloudhub"
	"github.com/google/wire"
)

func InitializeCloudhubTargetDatasource(configClient config.CloudhubConfigClient) (targets.Datasource, error) {

	wire.Build(
		cloudhubHttpClient.NewDefaultHttpClient,
		cloudhubService.NewDefaultService,
		cloudhub.NewDatasource,
		wire.Bind(new(cloudhubHttpClient.HttpClient), new(cloudhubHttpClient.DefaultHttpClient)),
		wire.Bind(new(cloudhubService.Service), new(cloudhubService.DefaultService)),
		wire.Bind(new(targets.Datasource), new(cloudhub.Datasource)),
	)

	return cloudhub.Datasource{}, nil
}
