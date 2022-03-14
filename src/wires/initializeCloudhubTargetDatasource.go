//go:build wireinject
// +build wireinject

package wires

import (
	cloudhubHttpClient "github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets/cloudhub"
	cloudhubService "github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
	"github.com/google/wire"
)

func InitializeCloudhubTargetDatasource(configClient config.CloudhubConfigClient) (targets.TargetDatasource, error) {

	wire.Build(
		cloudhubHttpClient.NewDefaultHttpClient,
		cloudhubService.NewDefaultService,
		cloudhub.NewDatasource,
		wire.Bind(new(cloudhubHttpClient.HttpClient), new(cloudhubHttpClient.DefaultHttpClient)),
		wire.Bind(new(cloudhubService.Service), new(cloudhubService.DefaultService)),
		wire.Bind(new(targets.TargetDatasource), new(cloudhub.Datasource)),
	)

	return cloudhub.Datasource{}, nil
}
