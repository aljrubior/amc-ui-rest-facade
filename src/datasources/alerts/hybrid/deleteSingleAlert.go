package hybrid

func (t DefaultDatasource) DeleteSingleAlert(token, orgId, envId, alertId string) error {

	return t.hybridService.DeleteSingleAlert(token, orgId, envId, alertId)
}
