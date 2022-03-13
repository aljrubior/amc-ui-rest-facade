package fabric

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/fabric"
)

func NewDefaultService(httpClient fabric.HttpClient) DefaultService {
	return DefaultService{
		httpClient: httpClient,
	}
}

type DefaultService struct {
	httpClient fabric.HttpClient
}
