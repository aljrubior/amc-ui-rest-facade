package atlas

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/atlas/requests"
	"github.com/aljrubior/amc-ui-rest-facade/clients/atlas/responses/target"
	"io/ioutil"
	"net/http"
	"time"
)

func (t DefaultHttpClient) GetTargets(token, orgId, envId, providerId string) (*[]target.Response, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetTargetsRequest(&t.config, token, orgId, envId, providerId).Build()

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

	var response []target.Response

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
