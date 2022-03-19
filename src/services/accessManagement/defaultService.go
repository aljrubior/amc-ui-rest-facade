package accessManagement

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(httpClient accessManagement.HttpClient) DefaultService {
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
	httpClient accessManagement.HttpClient
}
