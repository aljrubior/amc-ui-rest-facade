package applicationController

import "net/http"

type ApplicationController interface {
	GetApplications(w http.ResponseWriter, r *http.Request)
}
