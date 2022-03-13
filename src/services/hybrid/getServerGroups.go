package hybrid

import "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/serverGroup"

func (t DefaultService) GetServerGroups(token, orgId, envId string) (*[]serverGroup.Response, error) {

	resp, err := t.httpClient.GetServerGroups(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
