package main

import (
	"OmarFaruk-0x01/sms-trap/app"
	"OmarFaruk-0x01/sms-trap/app/database"
	"OmarFaruk-0x01/sms-trap/app/database/migration"
	"OmarFaruk-0x01/sms-trap/app/routes"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func main() {

	if len(os.Args) < 2 {
		panic("Not enough arguments. commands: \n'migrate:up', \n'migrate:down', \n'migrate:reset', \n'serve'")
	}

	db, err := database.NewDB()

	if err != nil {
		panic(err)
	}

	echo := echo.New()

	echo.Validator = app.NewAppValidator()

	routers := []routes.Router{
		routes.NewWebRouter("", echo, db),
		routes.NewApiRouter("/api/v1", echo, db),
	}

	app := app.NewApp(echo, db, routers)

	runCommand(app, os.Args[1])
}

func runCommand(app *app.App, command string) {
	switch command {

	case "migrate:up":
		runMigration(app.Db, "up")

	case "migrate:down":
		runMigration(app.Db, "down")

	case "migrate:reset":
		runMigration(app.Db, "down")
		runMigration(app.Db, "up")

	case "serve":
		app.StartServer()

	default:
		panic("Not enough arguments. commands: \n'migrate:up', \n'migrate:down', \n'migrate:reset', \n'serve'")
	}
}

func runMigration(db *bun.DB, command string) {

	migration := migration.NewBunMigration(db)

	if command == "up" {

		migration.Up()

	} else if command == "down" {

		migration.Down()

	}

}
