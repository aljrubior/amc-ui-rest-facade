package formatters

import (
	"github.com/aljrubior/go-facade/clients/hybrid/responses/common"
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) formatCommonServers(servers *[]common.Server) []model.Target {

	var data []model.Target

	for _, v := range *servers {

		data = append(data, t.formatSingleServer(v))
	}

	return data
}
