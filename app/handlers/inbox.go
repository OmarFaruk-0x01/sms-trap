package handlers

import (
	"OmarFaruk-0x01/sms-trap/app/services"
	"OmarFaruk-0x01/sms-trap/views"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type InboxHandler struct {
	db             *bun.DB
	smsTrapService *services.SmsTrapService
}

type ShowInboxParam struct {
	phone string `param:"phone"`
}

func (inbox *InboxHandler) ShowInbox() echo.HandlerFunc {
	return func(c echo.Context) error {
		traps, err := inbox.smsTrapService.GetPhones()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		activePhone := c.Param("phone")

		return views.Render(views.Home(traps, activePhone), http.StatusOK, c)
	}
}

func NewInboxHandler(db *bun.DB) *InboxHandler {
	smsTrapService := services.NewSmsTrapService(db)

	return &InboxHandler{
		db,
		smsTrapService,
	}
}
