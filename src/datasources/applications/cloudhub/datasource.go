package cloudhub

import (
	"github.com/aljrubior/amc-ui-rest-facade/services/cloudhub"
)

func NewDatasource(cloudhubService cloudhub.Service) Datasource {
	return Datasource{
		cloudhubService: cloudhubService,
	}
}

type Datasource struct {
	cloudhubService cloudhub.Service
}
