package permission

import (
	"github.com/aljrubior/go-facade/services/permission"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultController
)

func NewDefaultController(permissionService permission.Service) DefaultController {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultController{
			permissionService: permissionService,
		}
	}
	return *instance
}

type DefaultController struct {
	permissionService permission.Service
}
