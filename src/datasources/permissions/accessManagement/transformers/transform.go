package transformers

import (
	"github.com/aljrubior/go-facade/model/responses/permission"
)

func (t *DefaultTransformer) Transform() permission.Response {

	return t.permissions
}
