package main

import (
	"OmarFaruk-0x01/sms-trap/app"
	"OmarFaruk-0x01/sms-trap/app/config"
	"OmarFaruk-0x01/sms-trap/app/database"
	"OmarFaruk-0x01/sms-trap/app/database/migration"
	"OmarFaruk-0x01/sms-trap/app/routes"
	"OmarFaruk-0x01/sms-trap/app/websocket"
	"embed"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

var (
	dbPath  string
	port    string
	env     string
	version string = "v0.0.13"
)

//go:embed public
var publicAssets embed.FS

func init() {
	flag.StringVar(&dbPath, "db-path", "", "Path to the SQLite database file")
	flag.StringVar(&port, "port", "1290", "Port to run the server on")
	flag.Parse()
}

func main() {

	env = os.Getenv("APP_ENV")

	db, tempFile, err := database.NewSqliteDB(dbPath, env)

	if err != nil {
		panic(err)
	}

	if tempFile != "" {
		defer os.Remove(tempFile)
	}

	runMigration(db, "up")

	echo := echo.New()

	echo.Validator = app.NewAppValidator()

	echo.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "public",
		Filesystem: http.FS(publicAssets),
	}))

	echo.HideBanner = env != "dev"

	hub := websocket.NewHub()

	appConfig := &config.AppConfig{
		Echo:    echo,
		Db:      db,
		Hub:     hub,
		Port:    port,
		Env:     env,
		Version: version,
	}

	routers := []routes.Router{
		routes.NewWebRouter("", appConfig),
		routes.NewApiRouter("/api/v1", appConfig),
		routes.NewWebSocketRouter("/ws/", appConfig),
	}

	app := app.NewApp(appConfig, routers)

	go hub.Run()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func(c chan os.Signal) {
		<-c
		app.Shutdown()
		fmt.Printf("Shutting down...\n")
		os.Exit(0)
	}(c)

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
