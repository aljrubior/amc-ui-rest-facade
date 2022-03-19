package application

import "github.com/gorilla/mux"

func (t *App) initializeRoutes() {

	t.Router = mux.NewRouter().StrictSlash(true)

	t.Router.HandleFunc("/applications", t.applicationController.GetApplications).Methods("GET")
	t.Router.HandleFunc("/targets", t.targetController.GetTargets).Methods("GET")
	t.Router.HandleFunc("/servers", t.serverController.GetServers).Methods("GET")
	t.Router.HandleFunc("/alerts", t.alertController.GetAlerts).Methods("GET")
	t.Router.HandleFunc("/alerts/{productId}", t.alertController.CreateAlert).Methods("POST")
	t.Router.HandleFunc("/alerts/{productId}/{alertId}", t.alertController.GetSingleAlert).Methods("GET")
	t.Router.HandleFunc("/alerts/{productId}/{alertId}", t.alertController.UpdateAlert).Methods("PATCH")
	t.Router.HandleFunc("/alerts/{productId}/{alertId}", t.alertController.DeleteAlert).Methods("DELETE")
	t.Router.HandleFunc("/alerts/{productId}/{alertId}/history", t.alertController.GetAlertHistory).Methods("GET")
	t.Router.HandleFunc("/alerts/{productId}/resource/{resourceId}/history", t.alertController.GetResourceAlertHistory).Methods("GET")

	t.Router.HandleFunc("/users", t.userController.GetUsers).Methods("GET")
	t.Router.HandleFunc("/permissions", t.permissionController.GetPermissions).Methods("GET")

	t.Router.HandleFunc("/ping", t.healthController.GetPing).Methods("GET")
	t.Router.HandleFunc("/info", t.healthController.GetInfo).Methods("GET")
}
