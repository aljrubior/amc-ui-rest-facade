package cloudhub

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/cloudhub/formatters"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/cloudhub/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func (t DefaultDatasource) GetApplications(token, orgId, envId string) (*[]application.Response, error) {

	resp, err := t.cloudhubService.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	formatterResult := formatters.NewResponseFormatter(resp).Format()

	return transformers.NewDefaultTransformer(formatterResult).Transform(), nil
}
