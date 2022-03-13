package hybrid

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/requests"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"io/ioutil"
	"net/http"
	"time"
)

func (t DefaultHttpClient) GetAlertHistory(token, orgId, envId, alertId string) (*alerts.AlertHistoriesResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetAlertHistoryRequest(&t.config, token, orgId, envId, alertId).Build()

	resp, err := httpClient.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, t.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response alerts.AlertHistoriesResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
