package transformers

import (
	"github.com/aljrubior/go-facade/datasources/applications/hybrid/model"
)

func NewDefaultTransformer(applications *[]model.Application) *DefaultTransformer {
	return &DefaultTransformer{
		applications: applications,
	}
}

type DefaultTransformer struct {
	applications *[]model.Application
}
