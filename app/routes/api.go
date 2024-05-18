package routes

import (
	"OmarFaruk-0x01/sms-trap/app/handlers"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type ApiRouter struct {
	router         *echo.Group
	db             *bun.DB
	smsTrapHandler *handlers.SmsTrapHandler
}

func (api *ApiRouter) Register() {
	api.router.GET("/trap", api.smsTrapHandler.Trap())
}

func NewApiRouter(prefix string, echo *echo.Echo, db *bun.DB) *ApiRouter {
	router := echo.Group(prefix)
	smsTrapHandler := handlers.NewSmsTrapHandler(db)

	return &ApiRouter{
		router,
		db,
		smsTrapHandler,
	}
}
