package cloudhub

import (
	"github.com/aljrubior/go-facade/clients/cloudhub"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(httpClient cloudhub.HttpClient) DefaultService {
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
	httpClient cloudhub.HttpClient
}
