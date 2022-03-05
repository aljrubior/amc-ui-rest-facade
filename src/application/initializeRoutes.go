package application

func (app *App) initializeRoutes() {

	app.Router.HandleFunc("/applications", app.applicationController.GetApplications).Methods("GET")

	app.Router.HandleFunc("/targets", app.targetController.GetTargets).Methods("GET")

	app.Router.HandleFunc("/servers", app.serverController.GetServers).Methods("GET")

	app.Router.HandleFunc("/permissions", app.permissionController.GetPermissions).Methods("GET")

	app.Router.HandleFunc("/users", app.userController.GetUsers).Methods("GET")

	app.Router.HandleFunc("/alerts", app.alertController.GetAlerts).Methods("GET")
	app.Router.HandleFunc("/alerts/{productId}", app.alertController.PostAlert).Methods("POST")
	app.Router.HandleFunc("/alerts/{productId}/{alertId}", app.alertController.GetAlert).Methods("GET")
	app.Router.HandleFunc("/alerts/{productId}/{alertId}", app.alertController.PatchAlert).Methods("PATCH")
	app.Router.HandleFunc("/alerts/{productId}/{alertId}", app.alertController.PutAlert).Methods("PUT")
	app.Router.HandleFunc("/alerts/{productId}/{alertId}", app.alertController.DeleteAlert).Methods("DELETE")
	app.Router.HandleFunc("/alerts/{productId}/{alertId}/history", app.alertController.GetAlertHistory).Methods("GET")
	app.Router.HandleFunc("/alerts/{productId}/resource/{resourceId}/history", app.alertController.GetResourceHistory).Methods("GET")

	app.Router.HandleFunc("/ping", app.healthController.GetPing).Methods("PUT")
	app.Router.HandleFunc("/info", app.healthController.GetInfo).Methods("DELETE")
}
