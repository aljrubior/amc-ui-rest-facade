package fabric

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultHttpClient
)

func NewDefaultHttpClient(config config.FabricConfigClient) DefaultHttpClient {

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
	config config.FabricConfigClient
}
