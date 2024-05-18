package app

import (
	"OmarFaruk-0x01/sms-trap/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

type App struct {
	echo    *echo.Echo
	db      *bun.DB
	routers []routes.Router
}

func (app *App) StartServer() {

	app.echo.Static("/static/*", "public")

	app.registerMiddlewares()

	app.registerRoutes()

	app.echo.Logger.Fatal(app.echo.Start(":1323"))
}

func (app *App) registerMiddlewares() {

	app.echo.Use(middleware.Logger())

}

func (app *App) registerRoutes() {
	for _, route := range app.routers {
		route.Register()
	}
}

func NewApp(echo *echo.Echo, db *bun.DB, routers []routes.Router) *App {
	return &App{
		echo,
		db,
		routers,
	}
}
