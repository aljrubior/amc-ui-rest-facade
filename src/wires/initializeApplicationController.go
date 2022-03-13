//go:build wireinject
// +build wireinject

package wires

import (
	applicationController "github.com/aljrubior/amc-ui-rest-facade/controllers/application"
	"github.com/aljrubior/amc-ui-rest-facade/datasources"
	applicationService "github.com/aljrubior/amc-ui-rest-facade/services/application"
	"github.com/google/wire"
)

func InitializeApplicationController(datasources []datasources.ApplicationDatasource) (applicationController.Controller, error) {

	wire.Build(
		applicationService.NewDefaultService,
		applicationController.NewDefaultController,
		wire.Bind(new(applicationService.Service), new(applicationService.DefaultService)),
		wire.Bind(new(applicationController.Controller), new(applicationController.DefaultController)),
	)

	return applicationController.DefaultController{}, nil
}
