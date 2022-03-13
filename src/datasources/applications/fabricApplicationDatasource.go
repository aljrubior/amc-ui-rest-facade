package applications

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
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

func (t FabricApplicationDatasource) GetApplications(token, orgId, envId string) (*[]application.Response, error) {

	data := make([]application.Response, 0)

	resp, err := t.fabricService.GetDeployments(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	apps := transformers.NewFabricApplicationTransformer(resp).Transform()

	data = append(data, apps...)

	return &data, nil
}
