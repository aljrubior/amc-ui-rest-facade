package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(datasources map[string]applications.Datasource) DefaultService {

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
	datasources map[string]applications.Datasource
}
