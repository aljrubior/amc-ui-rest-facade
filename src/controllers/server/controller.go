package server

import "net/http"

type Controller interface {
	GetServers(w http.ResponseWriter, r *http.Request)
}
