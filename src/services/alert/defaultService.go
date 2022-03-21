package alert

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/alerts"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(datasources map[string]alerts.Datasource) DefaultService {

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
	datasources map[string]alerts.Datasource
}
