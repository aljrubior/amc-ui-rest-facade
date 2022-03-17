package alert

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/alert"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultController
)

func NewDefaultController(alertService alert.Service) DefaultController {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultController{
			alertService: alertService,
		}
	}
	return *instance
}

type DefaultController struct {
	alertService alert.Service
}
