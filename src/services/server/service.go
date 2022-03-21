package server

import "github.com/aljrubior/go-facade/model/responses/target"

type Service interface {
	GetServers(token, orgId, envId string) (*target.DataResponse, error)
}
