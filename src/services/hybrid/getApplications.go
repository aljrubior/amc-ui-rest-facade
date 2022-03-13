package hybrid

import "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses"

func (t DefaultService) GetApplications(token, orgId, envId string) (*[]responses.ApplicationResponse, error) {

	resp, err := t.httpClient.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
