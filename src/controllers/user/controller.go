package user

import "net/http"

type Controller interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}
