package permission

import (
	"github.com/aljrubior/go-facade/datasources/permissions"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(datasource permissions.Datasource) DefaultService {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultService{
			datasource: datasource,
		}
	}

	return *instance
}

type DefaultService struct {
	datasource permissions.Datasource
}
