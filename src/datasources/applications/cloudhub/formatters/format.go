package formatters

import (
	"github.com/aljrubior/go-facade/datasources/applications/cloudhub/model"
)

func (t ResponseFormatter) Format() *[]model.Application {

	return model.NewApplicationBuilder(t.response).Build()
}
