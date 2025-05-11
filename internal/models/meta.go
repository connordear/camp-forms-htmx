package models

import (
	"database/sql"
	"log"

	database "github.com/connordear/camp-forms/internal/db"
)

type MetaModel struct {
	DB *sql.DB
}

func (m *MetaModel) RunMigrations() error {
	return database.RunSqlScript(m.DB, "./internal/db/migrations/init.sql")
}

func (m *MetaModel) InitDatabase(infoLog *log.Logger) error {
	var major, minor int
	err := m.DB.QueryRow("select major, minor from db_version").Scan(&major, &minor)

	if err != nil {
		infoLog.Println("No database version detected, running initialization...")
		err := m.RunMigrations()

		if err != nil {
			return err
		}
	}

	err = m.DB.QueryRow("select major, minor from db_version").Scan(&major, &minor)
	if err != nil {
		return err
	}

	infoLog.Printf("Database Version %d.%d", major, minor)

	// TODO: Check if there are new migrations available and run them
	// For now we can put everything in the init script and re run it every time
	return nil
}
