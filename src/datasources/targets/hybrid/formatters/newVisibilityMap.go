package formatters

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/cluster"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) newVisibilityMap(visibilityMap cluster.VisibilityMap) *model.VisibilityMap {

	return &model.VisibilityMap{
		VisibilityMap: cluster.VisibilityMap{
			MapNodes: visibilityMap.MapNodes,
		},
	}
}
