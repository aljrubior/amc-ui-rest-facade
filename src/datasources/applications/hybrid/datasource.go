package hybrid

import (
	"github.com/aljrubior/go-facade/services/hybrid"
)

func NewDatasource(hybridService hybrid.Service) Datasource {
	return Datasource{
		hybridService: hybridService,
	}
}

type Datasource struct {
	hybridService hybrid.Service
}
