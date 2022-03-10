package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetResourceAlertHistoryRequest(
	config *config.ApplicationManagerV1ConfigClient,
	bearerToken,
	orgId,
	envId,
	resourceId string) *GetResourceAlertHistoryRequest {

	return &GetResourceAlertHistoryRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
		resourceId:  resourceId,
	}
}

type GetResourceAlertHistoryRequest struct {
	clients.BaseHttpRequest
	config *config.ApplicationManagerV1ConfigClient
	bearerToken,
	orgId,
	envId,
	resourceId string
}

func (t *GetResourceAlertHistoryRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.ResourceAlertHistoryPath, t.resourceId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (request *GetResourceAlertHistoryRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	request.AddDefaultHeaders(req, request.orgId, request.envId, request.bearerToken)

	return req
}
