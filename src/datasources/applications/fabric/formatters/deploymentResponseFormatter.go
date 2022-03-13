package formatters

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"
	responses2 "github.com/aljrubior/amc-ui-rest-facade/clients/runtimeFabricManagement/responses"
)

func NewDeploymentResponseFormatter(
	deployments *[]responses.DeploymentResponse,
	targets *[]responses2.TargetResponse) DeploymentResponseFormatter {

	if targets == nil {
		targets = new([]responses2.TargetResponse)
	}

	if deployments == nil {
		deployments = new([]responses.DeploymentResponse)
	}

	mapTargets := make(map[string]responses2.TargetResponse)

	for _, t := range *targets {
		mapTargets[t.Id] = t
	}

	return DeploymentResponseFormatter{
		deployments: deployments,
		mapTargets:  mapTargets,
	}
}

type DeploymentResponseFormatter struct {
	deployments *[]responses.DeploymentResponse
	mapTargets  map[string]responses2.TargetResponse
}
