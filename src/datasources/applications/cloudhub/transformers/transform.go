package transformers

import (
	"github.com/aljrubior/go-facade/model/responses/application"
)

func (t *DefaultTransformer) Transform() *[]application.Response {
	result := make([]application.Response, 0)

	for _, v := range *t.response {
		result = append(result, v)
	}

	return &result
}
