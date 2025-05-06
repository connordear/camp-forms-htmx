package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

func RunSqlScript(errLog *log.Logger, db *sql.DB, sqlFilePath string) {
	sqlBytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		errLog.Fatalf("Error reading SQL file '%s': %v", sqlFilePath, err)
	}
	sqlScript := string(sqlBytes)

	_, err = db.Exec(sqlScript)
	if err != nil {
		errLog.Fatalf("Error executing SQL script from '%s': %v", sqlFilePath, err)
	}
}

func InitDatabase(infoLog *log.Logger, errLog *log.Logger) *sql.DB {
	db, err := sql.Open("sqlite", os.Getenv("DB_PATH"))

	if err != nil {
		errLog.Fatal(err)
	}

	var major, minor int
	err = db.QueryRow("select major, minor from db_version").Scan(&major, &minor)

	if err != nil {
		infoLog.Println("No database version detected, running initialization...")
		RunSqlScript(errLog, db, "./internal/db/migrations/init.sql")
	}

	err = db.QueryRow("select major, minor from db_version").Scan(&major, &minor)
	if err != nil {
		errLog.Fatal(err)
	}

	infoLog.Printf("Database Version %d.%d", major, minor)

	// TODO: Check if there are new migrations available and run them
	// For now we can put everything in the init script and re run it every time
	return db
}
