package runtimeFabricManagement

import "github.com/aljrubior/go-facade/clients/runtimeFabricManagement/responses"

type Service interface {
	GetTargets(token, orgId, envId string) (*[]responses.TargetResponse, error)
}
