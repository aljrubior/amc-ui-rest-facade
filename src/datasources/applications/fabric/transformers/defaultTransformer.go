package transformers

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/fabric/model"
)

func NewDefaultTransformer(response *[]model.Deployment) *DefaultTransformer {
	return &DefaultTransformer{
		response: response,
	}
}

type DefaultTransformer struct {
	response *[]model.Deployment
}
