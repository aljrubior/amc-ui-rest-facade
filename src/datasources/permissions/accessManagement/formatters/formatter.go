package formatters

import "github.com/aljrubior/amc-ui-rest-facade/datasources/permissions/accessManagement/model"

type Formatter interface {
	Format() model.Permissions
}
