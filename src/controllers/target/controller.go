package target

import "net/http"

type Controller interface {
	GetTargets(w http.ResponseWriter, r *http.Request)
}
