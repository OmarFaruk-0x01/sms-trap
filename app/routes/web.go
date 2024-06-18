package routes

import (
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type WebRouter struct {
	router       *echo.Group
	appConfig    *config.AppConfig
	inboxHandler *handlers.InboxHandler
	docsHandler  *handlers.DocsHandler
	guideHandler *handlers.GuideHandler
}

func (web *WebRouter) Register() {

	web.router.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/inbox")
	})

	web.router.GET("/inbox", web.inboxHandler.ShowInbox())
	web.router.GET("/inbox/:phone", web.inboxHandler.ShowInbox())
	web.router.GET("/docs", web.docsHandler.ShowDocs())
	web.router.GET("/guides", web.guideHandler.ShowGuides())
}

func (web *WebRouter) RegisterMiddlewares() {
	web.router.Use(middleware.RemoveTrailingSlash())
}

func NewWebRouter(prefix string, appConfig *config.AppConfig) Router {
	group := appConfig.Echo.Group(prefix)

	return &WebRouter{
		router:       group,
		appConfig:    appConfig,
		inboxHandler: handlers.NewInboxHandler(appConfig),
		docsHandler:  handlers.NewDocsHanlder(appConfig),
		guideHandler: handlers.NewGuideHandler(appConfig),
	}
}
