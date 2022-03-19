package user

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/accessManagement"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultController
)

func NewDefaultController(csService accessManagement.Service) DefaultController {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultController{
			csService: csService,
		}
	}
	return *instance
}

type DefaultController struct {
	csService accessManagement.Service
}
