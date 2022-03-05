package targetController

import "net/http"

type TargetController interface {
	GetTargets(w http.ResponseWriter, r *http.Request)
}
