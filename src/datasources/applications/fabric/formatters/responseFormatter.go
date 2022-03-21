package formatters

import (
	"github.com/aljrubior/go-facade/clients/fabric/responses"
	responses2 "github.com/aljrubior/go-facade/clients/runtimeFabricManagement/responses"
)

func NewResponseFormatter(
	deployments *[]responses.DeploymentResponse,
	targets *[]responses2.TargetResponse) ResponseFormatter {

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

	return ResponseFormatter{
		deployments: deployments,
		mapTargets:  mapTargets,
	}
}

type ResponseFormatter struct {
	deployments *[]responses.DeploymentResponse
	mapTargets  map[string]responses2.TargetResponse
}
