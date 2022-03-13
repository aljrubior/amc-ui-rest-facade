package runtimeFabricManagement

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/runtimeFabricManagement/responses"
)

func (t DefaultService) GetTargets(token, orgId, envId string) (*[]responses.TargetResponse, error) {

	return t.httpClient.GetTargets(token, orgId, envId)
}
