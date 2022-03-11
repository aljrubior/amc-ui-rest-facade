package cloudhub

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub/requests"
	"github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub/responses"
	"io/ioutil"
	"net/http"
	"time"
)

func (client DefaultHttpClient) GetApplications(token string) (*[]responses.ApplicationResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetApplications(&client.config, token).Build()

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

	var response []responses.ApplicationResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
