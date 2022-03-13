package cloudhub

import "github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub/responses"

func (t DefaultService) GetApplications(token string) (*[]responses.ApplicationResponse, error) {

	resp, err := t.httpClient.GetApplications(token)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
