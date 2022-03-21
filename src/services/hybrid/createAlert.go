package hybrid

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"github.com/aljrubior/go-facade/controllers/alert/requests"
)

func (t DefaultService) CreateAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.Response, error) {

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
