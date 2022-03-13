package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/application"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/health"
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
	applicationController application.Controller
	healthController      health.Controller
	serverController      serverController.ServerController
	targetController      targetController.TargetController
	permissionController  permissionController.PermissionController
	userController        userController.UserController
}

func (t *App) Initialize() {
	t.healthController = health.NewDefaultController()
	t.initializeRoutes()
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", app.Router))
}
