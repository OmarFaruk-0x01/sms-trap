package database

import (
	"database/sql"
	"io/ioutil"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

// NewSqliteDB creates and returns a new bun.DB instance for SQLite
func NewSqliteDB(dbPath string) (db *bun.DB, tempFile string, err error) {
	var sqldb *sql.DB

	if dbPath == "" {
		// Create a temporary file
		tempFile, err = createTempFile()
		if err != nil {
			return nil, "", err
		}
		dbPath = tempFile
	}

	// Open SQLite database
	sqldb, err = sql.Open(sqliteshim.ShimName, "file:"+dbPath+"?cache=shared")
	if err != nil {
		return nil, tempFile, err
	}

	// Create a new bun.DB instance
	db = bun.NewDB(sqldb, sqlitedialect.New())

	// Add query hook for debugging
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db, tempFile, nil
}

func createTempFile() (string, error) {
	tempDir := os.TempDir()
	tempFile, err := ioutil.TempFile(tempDir, "sms-trap-*.db")
	if err != nil {
		return "", err
	}
	tempFile.Close()
	return tempFile.Name(), nil
}
