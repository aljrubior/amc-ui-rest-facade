package fabric

import (
	"github.com/aljrubior/go-facade/services/fabric"
	"github.com/aljrubior/go-facade/services/runtimeFabricManagement"
)

func NewDatasource(
	fabricService fabric.Service,
	runtimeFabricManagementService runtimeFabricManagement.Service) Datasource {
	return Datasource{
		fabricService:                  fabricService,
		runtimeFabricManagementService: runtimeFabricManagementService,
	}
}

type Datasource struct {
	fabricService                  fabric.Service
	runtimeFabricManagementService runtimeFabricManagement.Service
}
