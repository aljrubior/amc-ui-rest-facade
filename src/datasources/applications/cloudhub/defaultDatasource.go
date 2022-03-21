package cloudhub

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
)

func NewDefaultDatasource(cloudhubService cloudhub.Service) DefaultDatasource {
	return DefaultDatasource{
		cloudhubService: cloudhubService,
	}
}

type DefaultDatasource struct {
	cloudhubService cloudhub.Service
}
