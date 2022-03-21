package formatters

import (
	model2 "github.com/aljrubior/go-facade/datasources/applications/fabric/model"
)

func (t ResponseFormatter) Format() *[]model2.Deployment {

	var data []model2.Deployment

	for _, d := range *t.deployments {

		target, exists := t.mapTargets[d.Target.TargetId]

		if !exists {
			continue
		}

		data = append(data, model2.NewDeploymentBuilder(&d, &target).Build())
	}

	return &data
}
