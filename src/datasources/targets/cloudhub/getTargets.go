package cloudhub

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets/cloudhub/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/target"
)

func (t Datasource) GetTargets(token, orgId, envId string) (*[]target.Response, error) {

	apps, err := t.cloudhubService.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return transformers.NewDefaultTransformer(apps).Transform()
}
