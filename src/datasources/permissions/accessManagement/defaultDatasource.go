package accessManagement

import (
	"github.com/aljrubior/go-facade/services/accessManagement"
)

func NewDefaultDatasource(accessManagementService accessManagement.Service) DefaultDatasource {
	return DefaultDatasource{
		accessManagementService: accessManagementService,
	}
}

type DefaultDatasource struct {
	accessManagementService accessManagement.Service
}
