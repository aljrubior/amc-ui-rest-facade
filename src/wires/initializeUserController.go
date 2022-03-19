//go:build wireinject
// +build wireinject

package wires

import (
	coreServiceClient "github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/user"
	"github.com/aljrubior/amc-ui-rest-facade/services/accessManagement"
	"github.com/google/wire"
)

func InitializeUserController(coreServiceConfigClient config.AccessManagementConfigClient) (user.Controller, error) {

	wire.Build(
		coreServiceClient.NewDefaultHttpClient,
		accessManagement.NewDefaultService,
		user.NewDefaultController,
		wire.Bind(new(coreServiceClient.HttpClient), new(coreServiceClient.DefaultHttpClient)),
		wire.Bind(new(accessManagement.Service), new(accessManagement.DefaultService)),
		wire.Bind(new(user.Controller), new(user.DefaultController)),
	)

	return user.DefaultController{}, nil
}
