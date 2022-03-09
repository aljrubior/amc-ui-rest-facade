package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewPutSingleAlertRequest(
	config *config.CloudhubConfigClient,
	bearerToken,
	orgId,
	envId,
	alertId string,
	body []byte) *PutSingleAlertRequest {

	return &PutSingleAlertRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
		alertId:     alertId,
		body:        body,
	}
}

type PutSingleAlertRequest struct {
	clients.BaseHttpRequest
	config *config.CloudhubConfigClient
	bearerToken,
	orgId,
	envId,
	alertId string
	body []byte
}

func (t *PutSingleAlertRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.AlertPath, t.alertId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *PutSingleAlertRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(t.body))

	t.AddDefaultHeaders(req, t.orgId, t.envId, t.bearerToken)

	t.AddContentType(req, clients.ContentTypeJSON)

	return req
}
