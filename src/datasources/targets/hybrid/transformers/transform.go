package transformers

import (
	"github.com/aljrubior/go-facade/model/responses/target"
)

func (t *DefaultTransformer) Transform() *[]target.Response {
	result := make([]target.Response, 0)

	for _, v := range *t.applications {
		result = append(result, v)
	}

	return &result
}
