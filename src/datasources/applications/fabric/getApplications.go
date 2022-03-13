package fabric

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/fabric/formatters"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/fabric/transformers"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func (t FabricApplicationDatasource) GetApplications(token, orgId, envId string) (*[]application.Response, error) {

	deployments, err := t.fabricService.GetDeployments(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	targets, err := t.runtimeFabricManagementService.GetTargets(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	formatterResult := formatters.NewDeploymentResponseFormatter(deployments, targets).Format()

	return transformers.NewFabricApplicationTransformer(formatterResult).Transform(), nil

}
