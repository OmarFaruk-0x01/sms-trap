package routes

import (
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/handlers"

	"github.com/labstack/echo/v4"
)

type ApiRouter struct {
	router         *echo.Group
	smsTrapHandler *handlers.SmsTrapHandler
	appConfig      *config.AppConfig
}

func (api *ApiRouter) Register() {
	api.router.GET("/trap", api.smsTrapHandler.Trap())
	api.router.DELETE("/traps", api.smsTrapHandler.DeleteAll())
}

func (api *ApiRouter) RegisterMiddlewares() {

}

func NewApiRouter(prefix string, appConfig *config.AppConfig) *ApiRouter {
	router := appConfig.Echo.Group(prefix)
	smsTrapHandler := handlers.NewSmsTrapHandler(appConfig)

	return &ApiRouter{
		router,
		smsTrapHandler,
		appConfig,
	}
}
