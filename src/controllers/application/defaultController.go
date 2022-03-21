package application

import (
	"github.com/aljrubior/go-facade/services/application"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultController
)

func NewDefaultController(applicationService application.Service) DefaultController {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultController{
			applicationService: applicationService,
		}
	}
	return *instance
}

type DefaultController struct {
	applicationService application.Service
}
