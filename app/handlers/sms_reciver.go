package handlers

import (
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/models"
	"OmarFaruk-0x01/sms-trap/app/services"
	"OmarFaruk-0x01/sms-trap/app/websocket"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SmsTrapHandler struct {
	appConfig      *config.AppConfig
	smsTrapService *services.SmsTrapService
}

type TrapQuery struct {
	Phones  []string `json:"phones" query:"phones[]" validate:"required,min=1,dive,numeric"`
	Message string   `json:"message" query:"message" validate:"required"`
	Label   string   `json:"label" query:"label" validate:"required,oneof=transactional promotional"`
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
				Label:   query.Label,
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

		sms.appConfig.Hub.Broadcast(websocket.NewMessage([]byte(message), nil))

		return c.JSON(200, echo.Map{
			"message": "success",
		})
	}
}

func (sms *SmsTrapHandler) DeleteAll() echo.HandlerFunc {

	return func(c echo.Context) error {

		err := sms.smsTrapService.DeleteAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(200, echo.Map{
			"message": "success",
		})
	}
}

func NewSmsTrapHandler(appConfig *config.AppConfig) *SmsTrapHandler {
	smsTrapService := services.NewSmsTrapService(appConfig.Db)

	return &SmsTrapHandler{
		appConfig,
		smsTrapService,
	}
}
