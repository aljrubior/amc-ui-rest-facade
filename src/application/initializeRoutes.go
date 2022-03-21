package application

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (t *App) initializeRoutes() {

	t.Router = mux.NewRouter().StrictSlash(true)

	t.Router.HandleFunc("/applications", t.applicationController.GetApplications).Methods(http.MethodGet)
	t.Router.HandleFunc("/targets", t.targetController.GetTargets).Methods(http.MethodGet)
	t.Router.HandleFunc("/servers", t.serverController.GetServers).Methods(http.MethodGet)
	t.Router.HandleFunc("/alerts", t.alertController.GetAlerts).Methods(http.MethodGet)
	t.Router.HandleFunc("/alerts/{productId}", t.alertController.CreateAlert).Methods(http.MethodPost)
	t.Router.HandleFunc("/alerts/{productId}/{alertId}", t.alertController.GetSingleAlert).Methods(http.MethodGet)
	t.Router.HandleFunc("/alerts/{productId}/{alertId}", t.alertController.UpdateAlert).Methods(http.MethodPatch)
	t.Router.HandleFunc("/alerts/{productId}/{alertId}", t.alertController.DeleteAlert).Methods(http.MethodDelete)
	t.Router.HandleFunc("/alerts/{productId}/{alertId}/history", t.alertController.GetAlertHistory).Methods(http.MethodGet)
	t.Router.HandleFunc("/alerts/{productId}/resource/{resourceId}/history", t.alertController.GetResourceAlertHistory).Methods(http.MethodGet)

	t.Router.HandleFunc("/users", t.userController.GetUsers).Methods(http.MethodGet)
	t.Router.HandleFunc("/permissions", t.permissionController.GetPermissions).Methods(http.MethodGet)

	t.Router.HandleFunc("/ping", t.healthController.GetPing).Methods(http.MethodGet)
	t.Router.HandleFunc("/info", t.healthController.GetInfo).Methods(http.MethodGet)
}
