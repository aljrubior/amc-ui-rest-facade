package server

import (
	"github.com/aljrubior/go-facade/services/server"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultController
)

func NewDefaultController(serverService server.Service) DefaultController {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultController{
			serverService: serverService,
		}
	}
	return *instance
}

type DefaultController struct {
	serverService server.Service
}
