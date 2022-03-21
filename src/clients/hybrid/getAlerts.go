package hybrid

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/clients/hybrid/requests"
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"io/ioutil"
	"net/http"
	"time"
)

func (client DefaultHttpClient) GetAlerts(token, orgId, envId string) (*alerts.AlertsResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetAlertsRequest(&client.config, token, orgId, envId).Build()

	resp, err := httpClient.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, client.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response alerts.AlertsResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
