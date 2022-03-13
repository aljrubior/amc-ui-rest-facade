package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
	"github.com/aljrubior/amc-ui-rest-facade/services/fabric"
	"github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance DefaultController
)

func NewDefaultService(
	hybridService hybrid.Service,
	cloudhubService cloudhub.Service,
	fabricService fabric.Service) DefaultController {

	lock.Lock()
	defer lock.Unlock()

	if &instance == nil {
		instance = DefaultController{
			hybridService:   hybridService,
			cloudhubService: cloudhubService,
			fabricService:   fabricService,
		}
	}
	return instance
}

func NewDefaultController(
	hybridService hybrid.Service,
	cloudhubService cloudhub.Service,
	fabricService fabric.Service) DefaultController {

	return DefaultController{
		hybridService:   hybridService,
		cloudhubService: cloudhubService,
		fabricService:   fabricService,
	}
}

type DefaultController struct {
	hybridService   hybrid.Service
	cloudhubService cloudhub.Service
	fabricService   fabric.Service
}
