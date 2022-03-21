package application

import "net/http"

type Controller interface {
	GetApplications(w http.ResponseWriter, r *http.Request)
}
