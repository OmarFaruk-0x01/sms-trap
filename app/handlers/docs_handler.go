package handlers

import (
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/services"
	"OmarFaruk-0x01/sms-trap/app/views"
	"OmarFaruk-0x01/sms-trap/app/views/docs"
	"OmarFaruk-0x01/sms-trap/app/views/layouts"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DocsHandler struct {
	appConfig      *config.AppConfig
	smsTrapService *services.SmsTrapService
}

func (doc *DocsHandler) ShowDocs() echo.HandlerFunc {

	return func(ctx echo.Context) error {

		phones, err := doc.smsTrapService.GetPhones()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		return views.Render(docs.DocsPage(&docs.DocsPageProps{
			Port: doc.appConfig.Port,
			InboxLayoutProps: &layouts.InboxLayoutProps{
				Phones:      phones,
				ActivePhone: "",
				AppLayoutProps: &layouts.AppLayoutProps{
					ActiveRoute: ctx.Request().URL.Path,
				},
			},
		}), 200, ctx)
	}
}

func NewDocsHanlder(appConfig *config.AppConfig) *DocsHandler {
	return &DocsHandler{
		appConfig:      appConfig,
		smsTrapService: services.NewSmsTrapService(appConfig.Db),
	}
}
