package transformers

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications/fabric/formatters/model"
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func NewFabricApplicationTransformer(response *[]model.Deployment) *FabricApplicationTransformer {
	return &FabricApplicationTransformer{
		response: response,
	}
}

type FabricApplicationTransformer struct {
	response *[]model.Deployment
}

func (t FabricApplicationTransformer) Transform() *[]application.Response {
	result := make([]application.Response, 0)

	for _, v := range *t.response {

		println(fmt.Sprintf("%v", v))
		result = append(result, v)
	}

	return &result
}
