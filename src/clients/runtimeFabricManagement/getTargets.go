package runtimeFabricManagement

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/runtimeFabricManagement/requests"
	"github.com/aljrubior/amc-ui-rest-facade/clients/runtimeFabricManagement/responses"
	"io/ioutil"
	"net/http"
	"time"
)

func (t DefaultHttpClient) GetTargets(token, orgId, envId string) (*[]responses.TargetResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetTargetsRequest(&t.config, token, orgId, envId).Build()

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

	var response []responses.TargetResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
