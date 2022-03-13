package cloudhub

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub"
)

func NewDefaultService(httpClient cloudhub.HttpClient) DefaultService {
	return DefaultService{
		httpClient: httpClient,
	}

}

type DefaultService struct {
	httpClient cloudhub.HttpClient
}
