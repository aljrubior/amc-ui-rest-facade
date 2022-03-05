package permissionController

import "net/http"

type PermissionController interface {
	GetPermissions(w http.ResponseWriter, r *http.Request)
}
