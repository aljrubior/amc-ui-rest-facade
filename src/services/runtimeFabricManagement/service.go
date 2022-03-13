package runtimeFabricManagement

import "github.com/aljrubior/amc-ui-rest-facade/clients/runtimeFabricManagement/responses"

type Service interface {
	GetTargets(token, orgId, envId string) (*[]responses.TargetResponse, error)
}
