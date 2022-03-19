package accessManagement

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement/requests"
	"github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement/responses"
	"io/ioutil"
	"net/http"
	"time"
)

func (t DefaultHttpClient) GetMembers(token, orgId, envId string) (*responses.DataResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetMembersRequest(&t.config, token, orgId, envId).Build()

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

	var response responses.DataResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
