package alert

import "net/http"

type Controller interface {
	GetAlerts(w http.ResponseWriter, r *http.Request)
	CreateAlert(w http.ResponseWriter, r *http.Request)
	GetSingleAlert(w http.ResponseWriter, r *http.Request)
	UpdateAlert(w http.ResponseWriter, r *http.Request)
	DeleteAlert(w http.ResponseWriter, r *http.Request)
	GetAlertHistory(w http.ResponseWriter, r *http.Request)
	GetResourceAlertHistory(w http.ResponseWriter, r *http.Request)
}
