package hybrid

import (
	"github.com/aljrubior/go-facade/services/hybrid"
)

func NewDefaultDatasource(hybridService hybrid.Service) DefaultDatasource {
	return DefaultDatasource{
		hybridService: hybridService,
	}
}

type DefaultDatasource struct {
	hybridService hybrid.Service
}
