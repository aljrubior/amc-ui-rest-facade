package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetAlertHistoryRequest(
	config *config.HybridConfigClient,
	bearerToken,
	orgId,
	envId,
	alertId string) *GetAlertHistoryRequest {

	return &GetAlertHistoryRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
		alertId:     alertId,
	}
}

type GetAlertHistoryRequest struct {
	clients.BaseHttpRequest
	config *config.HybridConfigClient
	bearerToken,
	orgId,
	envId,
	alertId string
}

func (t *GetAlertHistoryRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.AlertHistoryPath, t.alertId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (request *GetAlertHistoryRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	request.AddDefaultHeaders(req, request.orgId, request.envId, request.bearerToken)

	return req
}
