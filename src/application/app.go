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

func (app *App) Initialize() {
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010"))
}
