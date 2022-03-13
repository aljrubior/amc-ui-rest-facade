package transformers

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub/responses"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func NewCloudhubApplicationTransformer(response *[]responses.ApplicationResponse) *CloudhubApplicationTransformer {
	return &CloudhubApplicationTransformer{
		response: response,
	}
}

type CloudhubApplicationTransformer struct {
	response *[]responses.ApplicationResponse
}

func (t *CloudhubApplicationTransformer) Transform() []application.Response {
	result := make([]application.Response, len(*t.response))

	for _, v := range *t.response {
		result = append(result, v)
	}

	return result
}
