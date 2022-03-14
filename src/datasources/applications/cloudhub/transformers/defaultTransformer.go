package transformers

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/cloudhub/model"
)

func NewDefaultTransformer(response *[]model.Application) *DefaultTransformer {
	return &DefaultTransformer{
		response: response,
	}
}

type DefaultTransformer struct {
	response *[]model.Application
}
