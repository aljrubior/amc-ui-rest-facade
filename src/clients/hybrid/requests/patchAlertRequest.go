package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewPatchSingleAlertRequest(
	config *config.ApplicationManagerV1ConfigClient,
	bearerToken,
	orgId,
	envId,
	alertId string,
	body []byte) *PatchSingleAlertRequest {

	return &PatchSingleAlertRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
		alertId:     alertId,
		body:        body,
	}
}

type PatchSingleAlertRequest struct {
	clients.BaseHttpRequest
	config *config.ApplicationManagerV1ConfigClient
	bearerToken,
	orgId,
	envId,
	alertId string
	body []byte
}

func (t *PatchSingleAlertRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.AlertsPath, t.alertId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (request *PatchSingleAlertRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(request.body))

	request.AddDefaultHeaders(req, request.orgId, request.envId, request.bearerToken)
	request.AddContentType(req, clients.ContentTypeJSON)

	return req
}
