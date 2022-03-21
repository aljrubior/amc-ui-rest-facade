package transformers

import "github.com/aljrubior/go-facade/model/responses/permission"

type Transformer interface {
	Transform() *permission.DataResponse
}
