package runtimeFabricManagement

import "github.com/aljrubior/go-facade/clients/runtimeFabricManagement/responses"

type HttpClient interface {
	GetTargets(token, orgId, envId string) (*[]responses.TargetResponse, error)
}
