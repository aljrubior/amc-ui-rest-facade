package applications

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
	"github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
)

func NewHybridApplicationDatasource(hybridService hybrid.Service) HybridApplicationDatasource {
	return HybridApplicationDatasource{
		hybridService: hybridService,
	}
}

type HybridApplicationDatasource struct {
	hybridService hybrid.Service
}

func (t HybridApplicationDatasource) GetApplications(token, orgId, envId string) (*application.DataResponse, error) {

	data := make([]application.Response, 0)

	hybridApps, err := t.hybridService.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	data = append(data, transformers.NewHybridApplicationTransformer(hybridApps).Transform()...)

	return application.NewDataResponse(&data), nil
}
