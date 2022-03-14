package target

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(datasources []targets.TargetDatasource) DefaultService {

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
	datasources []targets.TargetDatasource
}
