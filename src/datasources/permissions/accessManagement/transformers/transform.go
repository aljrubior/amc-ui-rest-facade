package transformers

import (
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/permission"
)

func (t *DefaultTransformer) Transform() permission.Response {

	return t.permissions
}
