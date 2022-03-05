package services

type ApplicationManagerV1Service interface {
	GetAlerts(token, orgId, envId string)
}
