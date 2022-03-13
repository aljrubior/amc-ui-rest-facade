package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/applicationController"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/healthController"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/permissionController"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/serverController"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/targetController"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/userController"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewApp() *App {
	app := App{}
	app.Initialize()

	return &app
}

type App struct {
	Router *mux.Router

	alertController       alertController.AlertController
	applicationController applicationController.ApplicationController
	healthController      healthController.HealthController
	serverController      serverController.ServerController
	targetController      targetController.TargetController
	permissionController  permissionController.PermissionController
	userController        userController.UserController
}

func (t *App) Initialize() {
	t.healthController = healthController.NewDefaultHealthController()
	t.initializeRoutes()
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", app.Router))
}
