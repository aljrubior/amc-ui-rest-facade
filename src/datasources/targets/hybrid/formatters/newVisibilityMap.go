package formatters

import (
	"github.com/aljrubior/go-facade/clients/hybrid/responses/cluster"
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) newVisibilityMap(visibilityMap cluster.VisibilityMap) *model.VisibilityMap {

	return &model.VisibilityMap{
		VisibilityMap: cluster.VisibilityMap{
			MapNodes: visibilityMap.MapNodes,
		},
	}
}
