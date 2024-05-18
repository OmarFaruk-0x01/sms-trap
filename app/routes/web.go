package routes

import (
	"OmarFaruk-0x01/sms-trap/app/handlers"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type WebRouter struct {
	router       *echo.Group
	db           *bun.DB
	inboxHandler *handlers.InboxHandler
}

func (web *WebRouter) Register() {

	web.router.GET("/", web.inboxHandler.ShowInbox())
	web.router.GET("/:phone", web.inboxHandler.ShowInbox())

}

func NewWebRouter(prefix string, echo *echo.Echo, db *bun.DB) Router {
	group := echo.Group(prefix)
	inboxHandler := handlers.NewInboxHandler(db)

	return &WebRouter{
		router:       group,
		db:           db,
		inboxHandler: inboxHandler,
	}
}
