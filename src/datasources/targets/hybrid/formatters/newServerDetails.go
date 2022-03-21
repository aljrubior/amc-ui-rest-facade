package formatters

import (
	"github.com/aljrubior/go-facade/clients/hybrid/responses/common"
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) newServerDetails(server common.Server) model.Details {

	runtimeVersion := server.GatewayVersion

	if server.ServerType == MuleServerType {
		runtimeVersion = server.MuleVersion
	}

	return model.Details{
		RuntimeVersion:      runtimeVersion,
		Type:                server.ServerType,
		AgentVersion:        server.AgentVersion,
		Addresses:           t.newServerAddresses(server),
		CurrentClusteringIp: server.CurrentClusteringIp,
	}
}
