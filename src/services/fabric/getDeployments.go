package fabric

import "github.com/aljrubior/go-facade/clients/fabric/responses"

func (t DefaultService) GetDeployments(token, orgId, envId string) (*[]responses.DeploymentResponse, error) {

	resp, err := t.httpClient.GetDeployments(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Items, nil
}
