package accessManagement

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement/responses"
)

func (t DefaultService) GetMembers(token, orgId, envId string) (*[]responses.MemberResponse, error) {

	resp, err := t.httpClient.GetMembers(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
