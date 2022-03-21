//go:build wireinject
// +build wireinject

package wires

import (
	coreServiceClient "github.com/aljrubior/go-facade/clients/accessManagement"
	"github.com/aljrubior/go-facade/config"
	"github.com/aljrubior/go-facade/controllers/user"
	"github.com/aljrubior/go-facade/services/accessManagement"
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
