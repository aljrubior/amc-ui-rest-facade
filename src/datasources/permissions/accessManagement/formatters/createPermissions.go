package formatters

import "github.com/aljrubior/go-facade/datasources/permissions/accessManagement/model"

func (t DefaultFormatter) createPermissions(resource string) model.Permission {

	return model.Permission{
		Read:   t.isRead(resource),
		Create: t.isCreate(resource),
		Modify: t.isModify(resource),
		Delete: t.isDelete(resource),
	}
}
