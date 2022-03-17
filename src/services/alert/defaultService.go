package alert

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
	"github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *DefaultService
)

func NewDefaultService(
	hybridService hybrid.Service,
	cloudhubService cloudhub.Service) DefaultService {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DefaultService{
			hybridService:   hybridService,
			cloudhubService: cloudhubService,
		}
	}

	return *instance
}

type DefaultService struct {
	hybridService   hybrid.Service
	cloudhubService cloudhub.Service
}
