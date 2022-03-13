package transformers

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func NewHybridApplicationTransformer(response *[]responses.ApplicationResponse) *HybridApplicationTransformer {
	return &HybridApplicationTransformer{
		response: response,
	}
}

type HybridApplicationTransformer struct {
	response *[]responses.ApplicationResponse
}

func (t *HybridApplicationTransformer) Transform() []application.Response {
	result := make([]application.Response, len(*t.response))

	for _, v := range *t.response {
		result = append(result, v)
	}

	return result
}
