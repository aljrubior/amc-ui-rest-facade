package fabric

import "github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"

func (t DefaultService) GetDeployments(token, orgId, envId string) (*[]responses.DeploymentResponse, error) {

	resp, err := t.httpClient.GetDeployments(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Items, nil
}
