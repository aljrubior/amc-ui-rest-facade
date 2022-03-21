package model

import (
	"github.com/aljrubior/go-facade/clients/cloudhub/responses"
)

const (
	CloudhubTargetType = "CLOUDHUB"
)

func NewApplicationBuilder(applications *[]responses.ApplicationResponse) ApplicationBuilder {
	return ApplicationBuilder{
		applications: applications,
	}
}

type ApplicationBuilder struct {
	applications *[]responses.ApplicationResponse
}

func (t ApplicationBuilder) Build() *[]Application {

	var result []Application

	for _, v := range *t.applications {

		fileName := ""

		if v.HasFile {
			fileName = v.FileName
		}

		result = append(result, Application{
			Id: v.Domain,
			Target: Target{
				Type: CloudhubTargetType,
			},
			Artifact: Artifact{
				LastUpdateTime: v.LastUpdateTime,
				CreateTime:     nil,
				Name:           v.Domain,
				FileName:       fileName,
			},
			MuleVersion: MuleVersion{
				Version:          v.MuleVersion.Version,
				UpdateId:         v.MuleVersion.UpdateId,
				LatestUpdateId:   v.MuleVersion.LatestUpdateId,
				EndOfSupportDate: v.MuleVersion.EndOfSupportDate,
			},
			IsDeploymentWaiting: v.IsDeploymentWaiting,
			LastReportedStatus:  v.Status,
			Details: Details{
				Domain: v.FullDomain,
			},
		})
	}

	return &result
}
