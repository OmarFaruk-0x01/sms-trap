package handlers

import (
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/services"
	"OmarFaruk-0x01/sms-trap/app/views"
	"OmarFaruk-0x01/sms-trap/app/views/guide"
	"OmarFaruk-0x01/sms-trap/app/views/layouts"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GuideHandler struct {
	appConfig      *config.AppConfig
	smsTrapService *services.SmsTrapService
}

func (g *GuideHandler) ShowGuides() echo.HandlerFunc {

	return func(ctx echo.Context) error {
		phones, err := g.smsTrapService.GetPhones()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		return views.Render(guide.ShowGuides(&guide.ShowGuidesProps{
			InboxLayoutProps: &layouts.InboxLayoutProps{
				Phones: phones,
				AppLayoutProps: &layouts.AppLayoutProps{
					ActiveRoute: ctx.Request().URL.Path,
				},
			},
		}), 200, ctx)
	}
}

func NewGuideHandler(appConfig *config.AppConfig) *GuideHandler {
	return &GuideHandler{
		appConfig:      appConfig,
		smsTrapService: services.NewSmsTrapService(appConfig.Db),
	}
}
