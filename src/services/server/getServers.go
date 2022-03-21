package server

import "github.com/aljrubior/go-facade/model/responses/target"

func (t DefaultService) GetServers(token, orgId, envId string) (*target.DataResponse, error) {

	data := make([]target.Response, 0)

	for _, v := range t.datasources {
		targets, err := v.GetTargets(token, orgId, envId)

		if err != nil {
			println("Error - ", err.Error())
			continue
		}

		data = append(data, *targets...)

	}

	return target.NewDataResponse(data), nil
}
