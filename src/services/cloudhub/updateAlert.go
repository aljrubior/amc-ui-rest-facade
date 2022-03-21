package cloudhub

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"github.com/aljrubior/go-facade/controllers/alert/requests"
)

func (t DefaultService) UpdateAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.Response, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.PutSingleAlert(token, orgId, envId, alertId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
