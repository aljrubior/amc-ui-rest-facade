package model

import (
	"github.com/aljrubior/go-facade/clients/hybrid/responses"
	"github.com/aljrubior/go-facade/clients/hybrid/responses/common"
)

func NewApplicationBuilder(response *[]responses.ApplicationResponse) ApplicationBuilder {
	return ApplicationBuilder{
		response: response,
	}
}

type ApplicationBuilder struct {
	response *[]responses.ApplicationResponse
}

func (t ApplicationBuilder) Build() *[]Application {

	var result []Application

	for _, v := range *t.response {
		result = append(result, Application{
			Id: v.Id,
			Target: Target{
				Id:     v.Target.Id,
				Type:   v.Target.Type,
				Name:   v.Target.Name,
				Status: v.Target.Status,
			},
			Artifact: Artifact{
				Name:           v.Artifact.Name,
				FileName:       v.Artifact.FileName,
				FileChecksum:   v.Artifact.FileChecksum,
				LastUpdateTime: v.Artifact.TimeUpdated,
				CreateTime:     v.TimeCreated,
			},
			LastReportedStatus: v.LastReportedStatus,
			DesiredStatus:      v.DesiredStatus,
			Details: Details{
				ServerArtifacts: t.buildServerArtifacts(v),
			},
			Uptime: v.Uptime,
		})
	}
	return &result
}

func (t ApplicationBuilder) buildServerArtifacts(response responses.ApplicationResponse) []ServerArtifact {

	var result []ServerArtifact

	mapServers := make(map[int]common.Server)

	for _, v := range response.Target.Servers {
		mapServers[v.Id] = v
	}

	for _, v := range response.ServerArtifacts {

		result = append(result, ServerArtifact{
			Target: ArtifactTarget{
				Id:     v.ServerId,
				Type:   mapServers[v.ServerId].Type,
				Name:   mapServers[v.ServerId].Name,
				Status: mapServers[v.ServerId].Status,
			},
			LastReportedStatus: response.LastReportedStatus,
			Message:            "",
		})
	}

	return result
}
