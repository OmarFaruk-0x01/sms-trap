package routes

import (
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/websocket"
	"log"

	gorillaWS "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WebSocketRouter struct {
	router    *echo.Group
	appConfig *config.AppConfig
}

var upgrader = gorillaWS.Upgrader{}

func (wsr *WebSocketRouter) Register() {
	wsr.router.GET("", func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

		if err != nil {
			log.Fatal(err)
		}

		client := websocket.NewClient(wsr.appConfig.Hub, conn)

		client.Register()

		return nil
	})

}

func (ws *WebSocketRouter) RegisterMiddlewares() {

}

func NewWebSocketRouter(prefix string, appConfig *config.AppConfig) Router {
	group := appConfig.Echo.Group(prefix)

	return &WebSocketRouter{
		router:    group,
		appConfig: appConfig,
	}
}
