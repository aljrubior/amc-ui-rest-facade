package formatters

import (
	responses3 "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/cluster"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/server"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/serverGroup"
)

func NewResponseFormatter(
	applications *[]responses3.ApplicationResponse,
	servers *[]server.Response,
	serverGroups *[]serverGroup.Response,
	clusters *[]cluster.Response) ResponseFormatter {

	mapApplications := make(map[int][]*responses3.ApplicationResponse)

	for _, v := range *applications {
		mapApplications[v.Target.Id] = append(mapApplications[v.Target.Id], &v)
	}

	return ResponseFormatter{
		applications: &mapApplications,
		servers:      servers,
		serverGroups: serverGroups,
		clusters:     clusters,
	}
}

type ResponseFormatter struct {
	applications *map[int][]*responses3.ApplicationResponse
	servers      *[]server.Response
	serverGroups *[]serverGroup.Response
	clusters     *[]cluster.Response
}
