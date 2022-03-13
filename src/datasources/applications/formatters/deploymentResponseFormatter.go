package formatters

import "github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"

type DeploymentResponseFormatter struct {
	deployments *[]responses.DeploymentResponse
	targets *[]responses.Target

}

func (t *DeploymentResponseFormatter) Format() {
	var data []Deployment

	for _, v := range *t.response {

		data = append(data , Deployment{
			Id: v.Id,
			Artifact: Artifact{
				LastUpdateTime: v
			},
		})

	}

}
