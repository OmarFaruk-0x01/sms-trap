package app

import (
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/routes"
	"fmt"

	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	*config.AppConfig
	Routers []routes.Router
}

func (app *App) StartServer() {

	app.registerMiddlewares()

	app.registerRoutes()

	fmt.Printf("SMS Trap Version %s\n\n", app.Version)

	app.Echo.Logger.Fatal(app.Echo.Start(":" + app.Port))

}

func (app *App) registerMiddlewares() {
	if app.Env == "dev" {
		app.Echo.Use(middleware.Logger())
	}

	app.Echo.Use(middleware.Recover())
}

func (app *App) registerRoutes() {
	for _, route := range app.Routers {
		route.RegisterMiddlewares()
		route.Register()
	}
}

func (app *App) Shutdown() {
	app.Echo.Close()
	app.Db.Close()
}

func NewApp(config *config.AppConfig, routers []routes.Router) *App {
	return &App{
		config,
		routers,
	}
}
