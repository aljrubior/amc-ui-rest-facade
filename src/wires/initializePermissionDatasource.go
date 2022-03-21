//go:build wireinject
// +build wireinject

package wires

import (
	accessManagementClient "github.com/aljrubior/go-facade/clients/accessManagement"
	"github.com/aljrubior/go-facade/config"
	"github.com/aljrubior/go-facade/datasources/permissions"
	accessManagementDatasource "github.com/aljrubior/go-facade/datasources/permissions/accessManagement"
	"github.com/aljrubior/go-facade/services/accessManagement"
	"github.com/google/wire"
)

func InitializePermissionDatasource(accessManagementConfigClient config.AccessManagementConfigClient) (permissions.Datasource, error) {

	wire.Build(
		accessManagementClient.NewDefaultHttpClient,
		accessManagement.NewDefaultService,
		accessManagementDatasource.NewDefaultDatasource,
		wire.Bind(new(accessManagementClient.HttpClient), new(accessManagementClient.DefaultHttpClient)),
		wire.Bind(new(accessManagement.Service), new(accessManagement.DefaultService)),
		wire.Bind(new(permissions.Datasource), new(accessManagementDatasource.DefaultDatasource)),
	)

	return accessManagementDatasource.DefaultDatasource{}, nil
}
