package handlers

import (
	"OmarFaruk-0x01/sms-trap/app/services"
	"OmarFaruk-0x01/sms-trap/app/views"
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
			activeTab = "stats"
		}

		if activeTimeFilter == "" {
			activeTimeFilter = "today"
		}

		selectedTraps, err := inbox.smsTrapService.FindAllByPhone(activePhone)
		if err != nil {
			return c.String(404, "Not Found")
		}

		return views.Render(views.Home(&views.HomeProps{
			Phones:           phones,
			SelectedTraps:    selectedTraps,
			ActivePhone:      activePhone,
			ActiveTab:        activeTab,
			ActiveTimeFilter: activeTimeFilter,
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
