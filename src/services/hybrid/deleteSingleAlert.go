package hybrid

func (t DefaultService) DeleteSingleAlert(token, orgId, envId, alertId string) error {

	return t.httpClient.DeleteSingleAlert(token, orgId, envId, alertId)
}
