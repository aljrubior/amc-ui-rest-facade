package cloudhub

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

func (t DefaultService) UpdateAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

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
