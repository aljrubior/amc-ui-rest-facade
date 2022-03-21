package application

import (
	"github.com/aljrubior/go-facade/controllers/alert"
	"github.com/aljrubior/go-facade/controllers/application"
	"github.com/aljrubior/go-facade/controllers/health"
	"github.com/aljrubior/go-facade/controllers/permission"
	"github.com/aljrubior/go-facade/controllers/server"
	"github.com/aljrubior/go-facade/controllers/target"
	"github.com/aljrubior/go-facade/controllers/user"
	"github.com/aljrubior/go-facade/datasources/alerts"
	"github.com/aljrubior/go-facade/datasources/applications"
	"github.com/aljrubior/go-facade/datasources/permissions"
	"github.com/aljrubior/go-facade/datasources/targets"
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
