package cloudhub

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert/requests"
)

func (t DefaultService) CreateAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.PostAlert(token, orgId, envId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
