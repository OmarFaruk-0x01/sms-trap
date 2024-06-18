package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

// dns will be pass through parameters
func NewMySqlDB() (db *bun.DB, err error) {

	sqldb, err := sql.Open("mysql", "root:@/sms-trap")

	if err != nil {
		panic(err)
	}

	db = bun.NewDB(sqldb, mysqldialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return
}
