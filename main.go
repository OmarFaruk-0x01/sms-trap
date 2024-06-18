package main

import (
	"OmarFaruk-0x01/sms-trap/app"
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/database"
	"OmarFaruk-0x01/sms-trap/app/database/migration"
	"OmarFaruk-0x01/sms-trap/app/routes"
	"OmarFaruk-0x01/sms-trap/app/websocket"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

var (
	dbPath string
	port   string
)

func init() {
	flag.StringVar(&dbPath, "db-path", "", "Path to the SQLite database file")
	flag.StringVar(&port, "port", "1290", "Port to run the server on")
	flag.Parse()
}

func main() {

	db, tempFile, err := database.NewSqliteDB(dbPath)

	if err != nil {
		panic(err)
	}

	if tempFile != "" {
		defer os.Remove(tempFile)
	}

	runMigration(db, "up")

	echo := echo.New()
	echo.Validator = app.NewAppValidator()
	hub := websocket.NewHub()

	appConfig := config.NewAppConfig(echo, db, hub, port)

	routers := []routes.Router{
		routes.NewWebRouter("", appConfig),
		routes.NewApiRouter("/api/v1", appConfig),
		routes.NewWebSocketRouter("/ws/", appConfig),
	}

	app := app.NewApp(appConfig, routers)

	go hub.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Printf("Shutting down...\n")
		app.Shutdown()
		os.Exit(0)
	}()

	app.StartServer()
}

func runMigration(db *bun.DB, command string) {

	migration := migration.NewBunMigration(db)

	if command == "up" {

		migration.Up()

	} else if command == "down" {

		migration.Down()

	}

}
