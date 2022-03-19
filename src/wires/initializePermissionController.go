//go:build wireinject
// +build wireinject

package wires

import (
	permissionController "github.com/aljrubior/amc-ui-rest-facade/controllers/permission"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/permissions"
	"github.com/aljrubior/amc-ui-rest-facade/services/permission"
	"github.com/google/wire"
)

func InitializePermissionController(datasource permissions.Datasource) (permissionController.Controller, error) {

	wire.Build(
		permission.NewDefaultService,
		permissionController.NewDefaultController,
		wire.Bind(new(permission.Service), new(permission.DefaultService)),
		wire.Bind(new(permissionController.Controller), new(permissionController.DefaultController)),
	)

	return permissionController.DefaultController{}, nil
}
