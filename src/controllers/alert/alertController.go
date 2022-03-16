package alertController

import "net/http"

type AlertController interface {
	GetAlerts(w http.ResponseWriter, r *http.Request)
	PostAlert(w http.ResponseWriter, r *http.Request)
	GetAlert(w http.ResponseWriter, r *http.Request)
	PatchAlert(w http.ResponseWriter, r *http.Request)
	PutAlert(w http.ResponseWriter, r *http.Request)
	DeleteAlert(w http.ResponseWriter, r *http.Request)
	GetAlertHistory(w http.ResponseWriter, r *http.Request)
	GetResourceHistory(w http.ResponseWriter, r *http.Request)
}
