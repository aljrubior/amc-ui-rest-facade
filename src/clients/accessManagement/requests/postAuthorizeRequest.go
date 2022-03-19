package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewPostAuthorizeRequest(
	config *config.AccessManagementConfigClient,
	bearerToken,
	orgId,
	envId string,
	body []byte) *PostAuthorizeRequest {

	return &PostAuthorizeRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
		body:        body,
	}
}

type PostAuthorizeRequest struct {
	clients.BaseHttpRequest
	config *config.AccessManagementConfigClient
	bearerToken,
	orgId,
	envId string
	body []byte
}

func (t *PostAuthorizeRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := t.config.AuthorizePath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *PostAuthorizeRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(t.body))

	t.AddDefaultHeaders(req, t.orgId, t.envId, t.bearerToken)
	t.AddContentType(req, clients.ContentTypeJSON)

	return req
}
