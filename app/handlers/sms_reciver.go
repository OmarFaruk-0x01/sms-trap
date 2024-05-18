package handlers

import (
	"OmarFaruk-0x01/sms-trap/app/models"
	"OmarFaruk-0x01/sms-trap/app/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type SmsTrapHandler struct {
	db             *bun.DB
	smsTrapService *services.SmsTrapService
}

type TrapQuery struct {
	Phones  []string `query:"phones[]" validate:"required,min=1,dive,numeric"`
	Message string   `query:"message" validate:"required"`
	Type    string   `query:"type" validate:"required,oneof=text unicode"`
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

		return c.JSON(200, query)
	}
}

func NewSmsTrapHandler(db *bun.DB) *SmsTrapHandler {
	smsTrapService := services.NewSmsTrapService(db)

	return &SmsTrapHandler{
		db,
		smsTrapService,
	}
}
