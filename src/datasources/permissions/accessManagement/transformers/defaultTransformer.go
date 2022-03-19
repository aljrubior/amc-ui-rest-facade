package transformers

import "github.com/aljrubior/amc-ui-rest-facade/datasources/permissions/accessManagement/model"

func NewDefaultTransformer(permissions *model.Permissions) *DefaultTransformer {
	return &DefaultTransformer{
		permissions: permissions,
	}
}

type DefaultTransformer struct {
	permissions *model.Permissions
}
