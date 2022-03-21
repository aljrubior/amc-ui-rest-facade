package health

import "net/http"

type Controller interface {
	GetPing(w http.ResponseWriter, r *http.Request)
	GetInfo(w http.ResponseWriter, r *http.Request)
}
