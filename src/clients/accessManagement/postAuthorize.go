package accessManagement

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/clients/accessManagement/requests"
	"github.com/aljrubior/go-facade/clients/accessManagement/responses"
	"io/ioutil"
	"net/http"
	"time"
)

func (client DefaultHttpClient) PostAuthorize(token, orgId, envId string, body []byte) (*responses.AuthorizeResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewPostAuthorizeRequest(&client.config, token, orgId, envId, body).Build()

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

	var response responses.AuthorizeResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
