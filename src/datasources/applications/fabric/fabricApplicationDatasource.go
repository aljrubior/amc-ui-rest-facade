package fabric

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/fabric"
	"github.com/aljrubior/amc-ui-rest-facade/services/runtimeFabricManagement"
)

func NewFabricApplicationDatasource(
	fabricService fabric.Service,
	runtimeFabricManagementService runtimeFabricManagement.Service) FabricApplicationDatasource {
	return FabricApplicationDatasource{
		fabricService:                  fabricService,
		runtimeFabricManagementService: runtimeFabricManagementService,
	}
}

type FabricApplicationDatasource struct {
	fabricService                  fabric.Service
	runtimeFabricManagementService runtimeFabricManagement.Service
}
