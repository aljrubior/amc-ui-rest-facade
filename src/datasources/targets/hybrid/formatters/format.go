package formatters

import (
	model2 "github.com/aljrubior/amc-ui-rest-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) Format() *[]model2.Target {

	var result []model2.Target

	data := t.formatServers()

	for _, v := range data {
		result = append(result, v)
	}

	data = t.formatServerGroups()

	for _, v := range data {
		result = append(result, v)
	}

	data = t.formatClusters()

	for _, v := range data {
		result = append(result, v)
	}

	return &result
}
