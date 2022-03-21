package cloudhub

func (t DefaultDatasource) DeleteSingleAlert(token, orgId, envId, alertId string) error {

	return t.cloudhubService.DeleteSingleAlert(token, orgId, envId, alertId)
}
