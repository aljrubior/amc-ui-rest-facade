package applications

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
	"github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
)

func NewCloudhubApplicationDatasource(cloudhubService cloudhub.Service) CloudhubApplicationDatasource {
	return CloudhubApplicationDatasource{
		cloudhubService: cloudhubService,
	}
}

type CloudhubApplicationDatasource struct {
	cloudhubService cloudhub.Service
}

func (t CloudhubApplicationDatasource) GetApplications(token, orgId, envId string) (*application.DataResponse, error) {

	data := make([]application.Response, 0)

	resp, err := t.cloudhubService.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	data = append(data, transformers.NewCloudhubApplicationTransformer(resp).Transform()...)

	return application.NewDataResponse(&data), nil
}
