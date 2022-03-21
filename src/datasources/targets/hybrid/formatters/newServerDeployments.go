package formatters

import (
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) newServerDeployments(serverId int) []model.Deployment {

	deployments := make([]model.Deployment, 0)

	for _, v := range (*t.applications)[serverId] {

		deployments = append(deployments, model.Deployment{
			Id: v.Id,
			Artifact: model.Artifact{
				Name:           v.Artifact.Name,
				FileName:       v.Artifact.FileName,
				FileChecksum:   v.Artifact.FileChecksum,
				LastUpdateTime: v.Artifact.TimeUpdated,
			},
			LastReportedStatus: v.LastReportedStatus,
		})
	}

	return deployments
}
