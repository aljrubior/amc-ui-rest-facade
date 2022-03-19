//go:build wireinject
// +build wireinject

package wires

import (
	accessManagementClient "github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/permissions"
	accessManagementDatasource "github.com/aljrubior/amc-ui-rest-facade/datasources/permissions/accessManagement"
	"github.com/aljrubior/amc-ui-rest-facade/services/accessManagement"
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
