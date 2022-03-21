package formatters

import (
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) formatServerGroups() []model.Target {

	var data []model.Target

	for _, v := range *t.serverGroups {

		data = append(data, model.Target{
			Id:          v.Id,
			Name:        v.Name,
			Type:        v.Type,
			Details:     t.newServerGroupDetails(v),
			Deployments: t.newServerDeployments(v.Id),
			Status:      v.Status,
		})
	}

	return data
}
