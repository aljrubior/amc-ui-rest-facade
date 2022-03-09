package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewDeleteSingleAlertRequest(
	config *config.ApplicationManagerV1ConfigClient,
	bearerToken,
	orgId,
	envId,
	alertId string) *DeleteSingleAlertRequest {

	return &DeleteSingleAlertRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
		alertId:     alertId,
	}
}

type DeleteSingleAlertRequest struct {
	clients.BaseHttpRequest
	config *config.ApplicationManagerV1ConfigClient
	bearerToken,
	orgId,
	envId,
	alertId string
}

func (t *DeleteSingleAlertRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.AlertPath, t.alertId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *DeleteSingleAlertRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodDelete, uri, nil)

	t.AddDefaultHeaders(req, t.orgId, t.envId, t.bearerToken)

	return req
}
