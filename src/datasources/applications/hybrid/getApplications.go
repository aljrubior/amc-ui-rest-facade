package hybrid

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/hybrid/formatters"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/hybrid/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func (t Datasource) GetApplications(token, orgId, envId string) (*[]application.Response, error) {

	rawApps, err := t.hybridService.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	formattedApps := formatters.NewResponseFormatter(rawApps).Format()

	return transformers.NewDefaultTransformer(formattedApps).Transform(), nil
}
