package handlers

import (
	"OmarFaruk-0x01/sms-trap/app/services"
	"OmarFaruk-0x01/sms-trap/app/views"
	"OmarFaruk-0x01/sms-trap/app/views/layouts"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type InboxHandler struct {
	db             *bun.DB
	smsTrapService *services.SmsTrapService
}

func (inbox *InboxHandler) ShowInbox() echo.HandlerFunc {
	return func(c echo.Context) error {
		phones, err := inbox.smsTrapService.GetPhones()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		activePhone := c.Param("phone")
		activeTab := c.QueryParam("tab")
		activeTimeFilter := c.QueryParam("time")

		if activeTab == "" {
			activeTab = "sms-list"
		}

		if activeTimeFilter == "" {
			activeTimeFilter = "today"
		}

		if activePhone != "" {

			_, err := inbox.smsTrapService.FindByPhone(activePhone)

			if err != nil {
				return c.Redirect(http.StatusFound, "/inbox")
			}
		}

		selectedTraps, err := inbox.smsTrapService.FindAllByPhone(activePhone)
		if err != nil {
			return c.String(404, "Not Found")
		}

		return views.Render(views.Inbox(&views.InboxProps{
			Phones:           phones,
			SelectedTraps:    selectedTraps,
			ActivePhone:      activePhone,
			ActiveTab:        activeTab,
			ActiveTimeFilter: activeTimeFilter,
			AppLayoutProps: &layouts.AppLayoutProps{
				ActiveRoute: c.Request().URL.Path,
			},
		}), http.StatusOK, c)
	}
}

func NewInboxHandler(db *bun.DB) *InboxHandler {
	smsTrapService := services.NewSmsTrapService(db)

	return &InboxHandler{
		db,
		smsTrapService,
	}
}
