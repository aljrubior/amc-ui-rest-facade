package target

import (
	targetService "github.com/aljrubior/go-facade/services/target"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultController
)

func NewDefaultController(targetService targetService.Service) DefaultController {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultController{
			targetService: targetService,
		}
	}
	return *instance
}

type DefaultController struct {
	targetService targetService.Service
}
