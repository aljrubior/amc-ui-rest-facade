package application

import (
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/application"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/health"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/permission"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/server"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/target"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/user"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/applications"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/permissions"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/targets"
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

	alertController       alert.Controller
	applicationController application.Controller
	healthController      health.Controller
	serverController      server.Controller
	targetController      target.Controller
	permissionController  permission.Controller
	userController        user.Controller

	applicationDatasources map[string]applications.Datasource
	targetDatasources      map[string]targets.Datasource
	alertDatasources       map[string]alerts.Datasource
	permissionDatasource   permissions.Datasource
}

func (t *App) Initialize() {

	// Configurations
	t.initializeConfigClient()

	// Datasources
	t.initializeApplicationDatasources()
	t.initializeTargetDatasources()
	t.initializePermissionDatasource()
	t.initializeAlertDatasources()

	// Controllers
	t.initializeControllers()
	t.healthController = health.NewDefaultController()
	t.initializeRoutes()
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", app.Router))
}
