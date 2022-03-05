package clients

import (
	"fmt"
	"net/http"
)

const (
	ContentTypeJSON = "application/json"
)

type BaseHttpRequest struct {
}

func (_ *BaseHttpRequest) GetBearerTokenValue(token string) string {
	return fmt.Sprintf("%s %s", "bearer", token)
}

func (_ *BaseHttpRequest) AddContentType(req *http.Request, contentType string) {

	req.Header.Add("Content-Type", contentType)
}

func (this *BaseHttpRequest) AddAuthorizationHeader(req *http.Request, token string) {

	req.Header.Add("Authorization", this.GetBearerTokenValue(token))
}

func (this *BaseHttpRequest) AddDefaultHeaders(req *http.Request, orgId, envId, token string) {

	req.Header.Add("Authorization", this.GetBearerTokenValue(token))
	req.Header.Add("X-ANYPNT-ENV-ID", envId)
	req.Header.Add("X-ANYPNT-ORG-ID", orgId)
}
