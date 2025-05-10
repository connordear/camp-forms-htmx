package database

import (
	"database/sql"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

func RunSqlScript(db *sql.DB, sqlFilePath string) error {
	sqlBytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return err
	}
	sqlScript := string(sqlBytes)

	_, err = db.Exec(sqlScript)
	if err != nil {
		return err
	}
	return nil
}

func OpenDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
