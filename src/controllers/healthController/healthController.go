package healthController

import "net/http"

type HealthController interface {
	GetPing(w http.ResponseWriter, r *http.Request)
	GetInfo(w http.ResponseWriter, r *http.Request)
}
