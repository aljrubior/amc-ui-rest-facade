//go:build wireinject
// +build wireinject

package wires

import (
	targetController "github.com/aljrubior/go-facade/controllers/target"
	"github.com/aljrubior/go-facade/datasources/targets"
	targetService "github.com/aljrubior/go-facade/services/target"
	"github.com/google/wire"
)

func InitializeTargetController(datasources map[string]targets.Datasource) (targetController.Controller, error) {

	wire.Build(
		targetService.NewDefaultService,
		targetController.NewDefaultController,
		wire.Bind(new(targetService.Service), new(targetService.DefaultService)),
		wire.Bind(new(targetController.Controller), new(targetController.DefaultController)),
	)

	return targetController.DefaultController{}, nil
}
