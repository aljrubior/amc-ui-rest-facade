package transformers

import "github.com/aljrubior/go-facade/datasources/permissions/accessManagement/model"

func NewDefaultTransformer(permissions *model.Permissions) *DefaultTransformer {
	return &DefaultTransformer{
		permissions: permissions,
	}
}

type DefaultTransformer struct {
	permissions *model.Permissions
}
