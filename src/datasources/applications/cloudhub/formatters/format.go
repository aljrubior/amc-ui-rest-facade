package formatters

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/cloudhub/model"
)

func (t ResponseFormatter) Format() *[]model.Application {

	return model.NewApplicationBuilder(t.response).Build()
}
