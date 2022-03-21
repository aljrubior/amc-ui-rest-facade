//go:build wireinject
// +build wireinject

package wires

import (
	serverController "github.com/aljrubior/go-facade/controllers/server"
	"github.com/aljrubior/go-facade/datasources/targets"
	serverService "github.com/aljrubior/go-facade/services/server"
	"github.com/google/wire"
)

func InitializeServerController(datasources []targets.Datasource) (serverController.Controller, error) {

	wire.Build(
		serverService.NewDefaultService,
		serverController.NewDefaultController,
		wire.Bind(new(serverService.Service), new(serverService.DefaultService)),
		wire.Bind(new(serverController.Controller), new(serverController.DefaultController)),
	)

	return serverController.DefaultController{}, nil
}
