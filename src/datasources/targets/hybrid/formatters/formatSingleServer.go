package formatters

import (
	"github.com/aljrubior/go-facade/clients/hybrid/responses/common"
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) formatSingleServer(server common.Server) model.Target {
	return model.Target{
		Id:          server.Id,
		Name:        server.Name,
		Type:        server.Type,
		Details:     t.newServerDetails(server),
		Deployments: t.newServerDeployments(server.Id),
		Status:      server.Status,
	}
}
