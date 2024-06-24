package config

import (
	"OmarFaruk-0x01/sms-trap/app/websocket"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type AppConfig struct {
	Echo *echo.Echo
	Db   *bun.DB
	Hub  *websocket.Hub
	Port string
	Env  string
}

func NewAppConfig(echo *echo.Echo, db *bun.DB, hub *websocket.Hub, port string, env string) *AppConfig {
	return &AppConfig{echo, db, hub, port, env}
}
