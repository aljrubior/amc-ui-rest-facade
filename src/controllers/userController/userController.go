package userController

import "net/http"

type UserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}
