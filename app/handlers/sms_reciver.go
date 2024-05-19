package handlers

import (
	"OmarFaruk-0x01/sms-trap/app/models"
	"OmarFaruk-0x01/sms-trap/app/services"
	"OmarFaruk-0x01/sms-trap/app/websocket"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type SmsTrapHandler struct {
	db             *bun.DB
	hub            *websocket.Hub
	smsTrapService *services.SmsTrapService
}

type TrapQuery struct {
	Phones  []string `json:"phones" query:"phones[]" validate:"required,min=1,dive,numeric"`
	Message string   `json:"message" query:"message" validate:"required"`
	Type    string   `json:"type" query:"type" validate:"required,oneof=text unicode"`
}

type TrapMessage struct {
	Query               *TrapQuery `json:"query"`
	Message             string     `json:"message"`
	TriggerNotification bool       `json:"trigger_notification"`
}

func (sms *SmsTrapHandler) Trap() echo.HandlerFunc {
	return func(c echo.Context) error {
		var query TrapQuery

		err := c.Bind(&query)

		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		err = c.Validate(query)

		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err)
		}

		traps := []*models.TrapDTO{}

		for _, phone := range query.Phones {
			traps = append(traps, &models.TrapDTO{
				Phone:   phone,
				Message: query.Message,
				Type:    query.Type,
			})
		}

		err = sms.smsTrapService.CreateMany(traps)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		message, err := json.Marshal(&TrapMessage{
			Query:               &query,
			Message:             "success",
			TriggerNotification: true,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		sms.hub.Broadcast(websocket.NewMessage([]byte(message), nil))

		return c.JSON(200, echo.Map{
			"message": "success",
		})
	}
}

func NewSmsTrapHandler(db *bun.DB, hub *websocket.Hub) *SmsTrapHandler {
	smsTrapService := services.NewSmsTrapService(db)

	return &SmsTrapHandler{
		db,
		hub,
		smsTrapService,
	}
}
