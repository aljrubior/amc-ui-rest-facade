package fabric

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/fabric"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(httpClient fabric.HttpClient) DefaultService {
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
	httpClient fabric.HttpClient
}
