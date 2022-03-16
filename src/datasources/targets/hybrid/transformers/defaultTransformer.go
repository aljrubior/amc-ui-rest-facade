package transformers

import "github.com/aljrubior/amc-ui-rest-facade/datasources/targets/hybrid/model"

func NewDefaultTransformer(applications *[]model.Target) *DefaultTransformer {
	return &DefaultTransformer{
		applications: applications,
	}
}

type DefaultTransformer struct {
	applications *[]model.Target
}
