package formatters

import "github.com/aljrubior/go-facade/datasources/permissions/accessManagement/model"

type Formatter interface {
	Format() model.Permissions
}
