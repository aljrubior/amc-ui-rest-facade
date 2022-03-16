package formatters

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/cluster"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) newClusterDetails(cluster cluster.Response) model.Details {

	return model.Details{
		Servers:          t.formatCommonServers(&cluster.Servers),
		MulticastEnabled: cluster.MulticastEnabled,
		PrimaryNodeId:    cluster.PrimaryNodeId,
		VisibilityMap:    t.newVisibilityMap(cluster.VisibilityMap),
	}
}
