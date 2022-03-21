package hybrid

import (
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/formatters"
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/transformers"
	"github.com/aljrubior/go-facade/model/responses/target"
)

func (t Datasource) GetTargets(token, orgId, envId string) (*[]target.Response, error) {

	apps, err := t.hybridService.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	servers, err := t.hybridService.GetServers(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	serverGroups, err := t.hybridService.GetServerGroups(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	clusters, err := t.hybridService.GetClusters(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	targets := formatters.NewResponseFormatter(apps, servers, serverGroups, clusters).Format()

	return transformers.NewDefaultTransformer(targets).Transform(), nil
}
