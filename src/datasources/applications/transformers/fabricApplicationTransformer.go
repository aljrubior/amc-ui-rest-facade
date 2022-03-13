package transformers

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func NewFabricApplicationTransformer(response *[]responses.DeploymentResponse) *FabricApplicationTransformer {
	return &FabricApplicationTransformer{
		response: response,
	}
}

type FabricApplicationTransformer struct {
	response *[]responses.DeploymentResponse
}

func (t *FabricApplicationTransformer) Transform() []application.Response {
	result := make([]application.Response, len(*t.response))

	for _, v := range *t.response {
		result = append(result, v)
	}

	return result
}
