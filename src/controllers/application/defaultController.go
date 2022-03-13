package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/application"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance DefaultController
)

func NewDefaultService(applicationService application.Service) DefaultController {

	lock.Lock()
	defer lock.Unlock()

	if &instance == nil {
		instance = DefaultController{
			applicationService: applicationService,
		}
	}
	return instance
}

func NewDefaultController(applicationService application.Service) DefaultController {

	return DefaultController{
		applicationService: applicationService,
	}
}

type DefaultController struct {
	applicationService application.Service
}
