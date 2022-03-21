package formatters

import (
	"github.com/aljrubior/go-facade/datasources/permissions/accessManagement/model"
)

func (t DefaultFormatter) Format() *model.Permissions {

	return &model.Permissions{
		Applications: t.createPermissions(model.APPLICATIONS_RESOURCE_NAME),
		Servers:      t.createPermissions(model.SERVERS_RESOURCE_NAME),
		ServerGroups: t.createPermissions(model.SERVER_GROUPS_RESOURCE_NAME),
		Clusters:     t.createPermissions(model.CLUSTERS_RESOURCE_NAME),
		Alerts:       t.createPermissions(model.ALERTS_RESOURCE_NAME),
		Services:     t.createPermissions(model.SERVICES_NAMESPACE),
	}
}
