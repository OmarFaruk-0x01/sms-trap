package main

import (
	"OmarFaruk-0x01/sms-trap/app"
	"OmarFaruk-0x01/sms-trap/app/models"
	"OmarFaruk-0x01/sms-trap/app/routes"
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {

	sqldb, err := sql.Open("mysql", "root:@/sms-trap")

	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	_, err = db.NewCreateTable().
		Model((*models.Trap)(nil)).
		IfNotExists().
		Exec(context.Background())

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

	app.StartServer()
}
