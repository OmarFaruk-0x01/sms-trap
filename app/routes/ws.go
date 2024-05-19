package routes

import (
	"OmarFaruk-0x01/sms-trap/app/websocket"
	"log"

	gorillaWS "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type WebSocketRouter struct {
	router *echo.Group
	db     *bun.DB
	hub    *websocket.Hub
}

var upgrader = gorillaWS.Upgrader{}

func (wsr *WebSocketRouter) Register() {
	wsr.router.GET("", func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

		if err != nil {
			log.Fatal(err)
		}

		client := websocket.NewClient(wsr.hub, conn)

		client.Register()

		return nil
	})

}

func NewWebSocketRouter(prefix string, echo *echo.Echo, db *bun.DB, hub *websocket.Hub) Router {
	group := echo.Group(prefix)

	return &WebSocketRouter{
		router: group,
		db:     db,
		hub:    hub,
	}
}
