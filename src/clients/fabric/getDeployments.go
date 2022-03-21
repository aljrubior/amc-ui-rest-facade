package fabric

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/clients/fabric/requests"
	"github.com/aljrubior/go-facade/clients/fabric/responses"
	"io/ioutil"
	"net/http"
	"time"
)

func (t DefaultHttpClient) GetDeployments(token, orgId, envId string) (*responses.DeploymentsResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetDeploymentsRequest(&t.config, token, orgId, envId).Build()

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

	var response responses.DeploymentsResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
