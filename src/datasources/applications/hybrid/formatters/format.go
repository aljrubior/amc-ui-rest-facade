package formatters

import (
	"github.com/aljrubior/go-facade/datasources/applications/hybrid/model"
)

func (t ResponseFormatter) Format() *[]model.Application {

	return model.NewApplicationBuilder(t.response).Build()
}
