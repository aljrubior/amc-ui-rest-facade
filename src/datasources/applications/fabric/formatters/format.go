package formatters

import "github.com/aljrubior/amc-ui-rest-facade/datasources/applications/fabric/formatters/model"

func (t DeploymentResponseFormatter) Format() *[]model.Deployment {

	var data []model.Deployment

	for _, d := range *t.deployments {

		target, exists := t.mapTargets[d.Target.TargetId]

		if !exists {
			continue
		}
		
		data = append(data, model.NewDeploymentBuilder(&d, &target).Build())
	}

	return &data
}
