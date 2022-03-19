package accessManagement

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement/responses"
	"github.com/aljrubior/amc-ui-rest-facade/services/accessManagement/requests"
)

func (t DefaultService) PostAuthorize(token, orgId, envId string, request requests.AuthorizeRequest) (*responses.AuthorizeResponse, error) {

	data, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	return t.httpClient.PostAuthorize(token, orgId, envId, data)
}
