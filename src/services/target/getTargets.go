package target

import (
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/target"
)

func (t DefaultService) GetTargets(token, orgId, envId string) (*target.DataResponse, error) {

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
