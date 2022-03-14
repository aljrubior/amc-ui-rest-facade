package fabric

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/fabric/formatters"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/fabric/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func (t Datasource) GetApplications(token, orgId, envId string) (*[]application.Response, error) {

	deployments, err := t.fabricService.GetDeployments(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	targets, err := t.runtimeFabricManagementService.GetTargets(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	formatterResult := formatters.NewResponseFormatter(deployments, targets).Format()

	return transformers.NewDefaultTransformer(formatterResult).Transform(), nil

}
