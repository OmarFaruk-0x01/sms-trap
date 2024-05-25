package app

import (
	"OmarFaruk-0x01/sms-trap/app/routes"
	"OmarFaruk-0x01/sms-trap/app/websocket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

type App struct {
	Echo    *echo.Echo
	Db      *bun.DB
	Routers []routes.Router
	Hub     *websocket.Hub
}

func (app *App) StartServer() {

	app.Echo.Static("/static/*", "public")

	app.registerMiddlewares()

	app.registerRoutes()

	app.Echo.Logger.Fatal(app.Echo.Start(":1323"))
}

func (app *App) registerMiddlewares() {

	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())
	app.Echo.Pre(middleware.RemoveTrailingSlash())

}

func (app *App) registerRoutes() {
	for _, route := range app.Routers {
		route.Register()
	}
}

func NewApp(echo *echo.Echo, db *bun.DB, routers []routes.Router, hub *websocket.Hub) *App {
	return &App{
		echo,
		db,
		routers,
		hub,
	}
}
