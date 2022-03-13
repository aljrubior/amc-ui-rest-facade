package hybrid

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(httpClient hybrid.HttpClient) DefaultService {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultService{
			httpClient: httpClient,
		}
	}

	return *instance
}

type DefaultService struct {
	httpClient hybrid.HttpClient
}
