package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewPostAlertRequest(
	config *config.ApplicationManagerV1ConfigClient,
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
	config *config.ApplicationManagerV1ConfigClient
	bearerToken,
	orgId,
	envId string
	body []byte
}

func (request *PostAlertRequest) buildUri() string {

	protocol := request.config.Protocol
	host := request.config.Host
	port := request.config.Port
	path := fmt.Sprintf("%s%s", request.config.Path, request.config.AlertsPath)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (request *PostAlertRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(request.body))

	request.AddDefaultHeaders(req, request.orgId, request.envId, request.bearerToken)
	request.AddContentType(req, clients.ContentTypeJSON)

	return req
}
