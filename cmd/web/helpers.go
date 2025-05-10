package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"runtime/debug"

	database "github.com/connordear/camp-forms/internal/db"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) initDatabase(db *sql.DB) error {
	var major, minor int
	err := db.QueryRow("select major, minor from db_version").Scan(&major, &minor)

	if err != nil {
		app.InfoLog.Println("No database version detected, running initialization...")
		err := database.RunSqlScript(db, "./internal/db/migrations/init.sql")

		if err != nil {
			app.ErrorLog.Fatal(err)
		}
	}

	err = db.QueryRow("select major, minor from db_version").Scan(&major, &minor)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	app.InfoLog.Printf("Database Version %d.%d", major, minor)

	// TODO: Check if there are new migrations available and run them
	// For now we can put everything in the init script and re run it every time
	return nil
}
