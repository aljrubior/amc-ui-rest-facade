package cluster

import "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/common"

type Response struct {
	Id               int             `json:"id"`
	TimeCreated      int64           `json:"timeCreated"`
	TimeUpdated      int64           `json:"timeUpdated"`
	Name             string          `json:"name"`
	Type             string          `json:"type"`
	Status           string          `json:"status"`
	MulticastEnabled bool            `json:"multicastEnabled"`
	PrimaryNodeId    int             `json:"primaryNodeId"`
	Servers          []common.Server `json:"servers"`
	VisibilityMap    VisibilityMap   `json:"visibilityMap"`
}
