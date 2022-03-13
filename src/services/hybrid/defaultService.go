package hybrid

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
)

func NewDefaultService(httpClient hybrid.HttpClient) DefaultService {
	return DefaultService{
		httpClient: httpClient,
	}
}

type DefaultService struct {
	httpClient hybrid.HttpClient
}
