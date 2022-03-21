package runtimeFabricManagement

import (
	"github.com/aljrubior/go-facade/clients"
	"github.com/aljrubior/go-facade/config"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultHttpClient
)

func NewDefaultHttpClient(config config.RuntimeFabricManagementClientConfig) DefaultHttpClient {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultHttpClient{
			config: config,
		}
	}

	return *instance
}

type DefaultHttpClient struct {
	clients.BaseHttpClient
	config config.RuntimeFabricManagementClientConfig
}
