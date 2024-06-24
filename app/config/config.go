package config

import (
	"OmarFaruk-0x01/sms-trap/app/websocket"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type AppConfig struct {
	Echo    *echo.Echo
	Db      *bun.DB
	Hub     *websocket.Hub
	Port    string
	Env     string
	Version string
}
