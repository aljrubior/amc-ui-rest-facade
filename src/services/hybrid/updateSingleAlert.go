package hybrid

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert/requests"
)

func (t DefaultService) UpdateSingleAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.Response, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.PatchSingleAlert(token, orgId, envId, alertId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
