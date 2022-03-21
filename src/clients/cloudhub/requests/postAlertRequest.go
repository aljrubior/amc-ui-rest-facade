package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/go-facade/clients"
	"github.com/aljrubior/go-facade/config"
	"net/http"
)

func NewPostAlertRequest(
	config *config.CloudhubConfigClient,
	bearerToken,
	orgId,
	envId string,
	body []byte) *PostAlertRequest {

	return &PostAlertRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
		body:        body,
	}
}

type PostAlertRequest struct {
	clients.BaseHttpRequest
	config *config.CloudhubConfigClient
	bearerToken,
	orgId,
	envId string
	body []byte
}

func (t *PostAlertRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := t.config.AlertsPath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *PostAlertRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(t.body))

	t.AddDefaultHeaders(req, t.orgId, t.envId, t.bearerToken)

	t.AddContentType(req, clients.ContentTypeJSON)

	return req
}
