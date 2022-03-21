package runtimeFabricManagement

import (
	"github.com/aljrubior/go-facade/clients/runtimeFabricManagement"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(httpClient runtimeFabricManagement.HttpClient) DefaultService {
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
	httpClient runtimeFabricManagement.HttpClient
}
