package hybrid

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/clients/hybrid/requests"
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"io/ioutil"
	"net/http"
	"time"
)

func (t DefaultHttpClient) GetResourceAlertHistory(token, orgId, envId, resourceId string) (*alerts.ResourceAlertHistoriesResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetResourceAlertHistoryRequest(&t.config, token, orgId, envId, resourceId).Build()

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

	var response alerts.ResourceAlertHistoriesResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
