package hybrid

import "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/cluster"

func (t DefaultService) GetClusters(token, orgId, envId string) (*[]cluster.Response, error) {

	resp, err := t.httpClient.GetClusters(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
