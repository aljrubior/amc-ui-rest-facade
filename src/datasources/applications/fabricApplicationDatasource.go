package applications

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
	"github.com/aljrubior/amc-ui-rest-facade/services/fabric"
)

func NewFabricApplicationDatasource(fabricService fabric.Service) FabricApplicationDatasource {
	return FabricApplicationDatasource{
		fabricService: fabricService,
	}
}

type FabricApplicationDatasource struct {
	fabricService fabric.Service
}

func (t FabricApplicationDatasource) GetApplications(token, orgId, envId string) (*application.DataResponse, error) {

	data := make([]application.Response, 0)

	resp, err := t.fabricService.GetDeployments(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	data = append(data, transformers.NewFabricApplicationTransformer(resp).Transform()...)

	return application.NewDataResponse(&data), nil
}
