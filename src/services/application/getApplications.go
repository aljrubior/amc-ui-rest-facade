package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/application"
)

func (t DefaultService) GetApplications(token, orgId, envId string) (*application.DataResponse, error) {

	data := make([]application.Response, 0)

	for _, v := range t.datasources {
		resp, err := v.GetApplications(token, orgId, envId)

		if err != nil {
			println("Error - ", err.Error())
			continue
		}

		data = append(data, *resp.Data)

	}

	return application.NewDataResponse(&data), nil
}
