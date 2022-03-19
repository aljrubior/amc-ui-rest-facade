package transformers

import "github.com/aljrubior/amc-ui-rest-facade/model/responses/permission"

type Transformer interface {
	Transform() *permission.DataResponse
}
