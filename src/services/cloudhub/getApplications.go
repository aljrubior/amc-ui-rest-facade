package cloudhub

import "github.com/aljrubior/go-facade/clients/cloudhub/responses"

func (t DefaultService) GetApplications(token, org, envId string) (*[]responses.ApplicationResponse, error) {

	resp, err := t.httpClient.GetApplications(token, org, envId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
