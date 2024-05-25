package routes

import (
	"OmarFaruk-0x01/sms-trap/app/handlers"
	"OmarFaruk-0x01/sms-trap/app/websocket"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type ApiRouter struct {
	router         *echo.Group
	db             *bun.DB
	hub            *websocket.Hub
	smsTrapHandler *handlers.SmsTrapHandler
}

func (api *ApiRouter) Register() {
	api.router.GET("/trap", api.smsTrapHandler.Trap())
	api.router.DELETE("/traps", api.smsTrapHandler.DeleteAll())
}

func NewApiRouter(prefix string, echo *echo.Echo, db *bun.DB, hub *websocket.Hub) *ApiRouter {
	router := echo.Group(prefix)
	smsTrapHandler := handlers.NewSmsTrapHandler(db, hub)

	return &ApiRouter{
		router,
		db,
		hub,
		smsTrapHandler,
	}
}
