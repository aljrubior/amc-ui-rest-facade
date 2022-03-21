package cloudhub

import (
	"github.com/aljrubior/go-facade/services/cloudhub"
)

func NewDatasource(cloudhubService cloudhub.Service) Datasource {
	return Datasource{
		cloudhubService: cloudhubService,
	}
}

type Datasource struct {
	cloudhubService cloudhub.Service
}
