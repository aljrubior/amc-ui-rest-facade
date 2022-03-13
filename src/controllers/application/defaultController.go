package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
	"github.com/aljrubior/amc-ui-rest-facade/services/fabric"
	"github.com/aljrubior/amc-ui-rest-facade/services/hybrid"
)

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
