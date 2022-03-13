package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(datasources []datasources.ApplicationDatasource) DefaultService {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultService{
			datasources: datasources,
		}
	}

	return *instance
}

type DefaultService struct {
	datasources []datasources.ApplicationDatasource
}
