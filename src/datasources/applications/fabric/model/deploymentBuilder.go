package model

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"
	responses2 "github.com/aljrubior/amc-ui-rest-facade/clients/runtimeFabricManagement/responses"
	"strings"
)

func NewDeploymentBuilder(
	deployment *responses.DeploymentResponse,
	target *responses2.TargetResponse) DeploymentBuilder {

	return DeploymentBuilder{
		deployment: deployment,
		target:     target,
	}

}

type DeploymentBuilder struct {
	deployment *responses.DeploymentResponse
	target     *responses2.TargetResponse
}

func (t DeploymentBuilder) Build() Deployment {

	targetName := t.target.Name

	if targetName == "" {
		targetName = "Runtime fabric"
	}

	return Deployment{
		Id: t.deployment.Id,
		Target: Target{
			Type:     t.deployment.Target.Provider,
			Provider: t.deployment.Target.Provider,
			Id:       t.target.Id,
			Name:     targetName,
			Status:   strings.ToUpper(t.target.Status),
		},
		Artifact: Artifact{
			LastUpdateTime: t.deployment.LastModifiedDate,
			CreateTime:     t.deployment.CreationDate,
			Name:           t.deployment.Name,
		},
		LastReportedStatus: t.deployment.Status,
		Details:            map[string]string{},
		Application: Application{
			Status: t.deployment.Application.Status,
		},
	}
}
