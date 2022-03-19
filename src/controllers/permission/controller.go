package permission

import "net/http"

type Controller interface {
	GetPermissions(w http.ResponseWriter, r *http.Request)
}
