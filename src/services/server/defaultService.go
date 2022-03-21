package server

import (
	"github.com/aljrubior/go-facade/datasources/targets"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(datasources []targets.Datasource) DefaultService {

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
	datasources []targets.Datasource
}
