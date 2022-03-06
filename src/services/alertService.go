package services

type AlertService interface {
	GetAlerts(token, orgId, envId string)
}
