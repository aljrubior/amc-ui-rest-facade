package hybrid

import "github.com/aljrubior/go-facade/clients/hybrid/responses/server"

func (t DefaultService) GetServers(token, orgId, envId string) (*[]server.Response, error) {

	resp, err := t.httpClient.GetServers(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
