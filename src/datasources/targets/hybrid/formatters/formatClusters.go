package formatters

import (
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) formatClusters() []model.Target {

	var data []model.Target

	for _, v := range *t.clusters {

		data = append(data, model.Target{
			Id:          v.Id,
			Name:        v.Name,
			Type:        v.Type,
			Details:     t.newClusterDetails(v),
			Deployments: t.newServerDeployments(v.Id),
			Status:      v.Status,
		})
	}

	return data
}
