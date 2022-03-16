package formatters

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/common"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) formatServers() []model.Target {

	commonServers := make([]common.Server, 0)

	for _, v := range *t.servers {
		commonServers = append(commonServers, v.Server)
	}

	data := t.formatCommonServers(&commonServers)

	return data
}
