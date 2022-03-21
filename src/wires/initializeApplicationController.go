//go:build wireinject
// +build wireinject

package wires

import (
	applicationController "github.com/aljrubior/go-facade/controllers/application"
	"github.com/aljrubior/go-facade/datasources/applications"
	applicationService "github.com/aljrubior/go-facade/services/application"
	"github.com/google/wire"
)

func InitializeApplicationController(datasources map[string]applications.Datasource) (applicationController.Controller, error) {

	wire.Build(
		applicationService.NewDefaultService,
		applicationController.NewDefaultController,
		wire.Bind(new(applicationService.Service), new(applicationService.DefaultService)),
		wire.Bind(new(applicationController.Controller), new(applicationController.DefaultController)),
	)

	return applicationController.DefaultController{}, nil
}
