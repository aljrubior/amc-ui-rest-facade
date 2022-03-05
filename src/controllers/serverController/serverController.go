package serverController

import "net/http"

type ServerController interface {
	GetServers(w http.ResponseWriter, r *http.Request)
}
