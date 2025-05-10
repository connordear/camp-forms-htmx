package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/connordear/camp-forms/internal/db"
	"github.com/connordear/camp-forms/internal/models"
)

type application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Camps    *models.CampModel
}

func main() {
	port := flag.String("port", ":4000", "HTTP Port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := database.OpenDb(os.Getenv("DB_PATH"))
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Camps:    &models.CampModel{DB: db},
	}

	if err = app.initDatabase(db); err != nil {
		errorLog.Fatal(err)
	}

	server := http.Server{
		Addr:     *port,
		Handler:  Router(app),
		ErrorLog: errorLog,
	}

	app.InfoLog.Println("Listening on port ", *port)
	server.ListenAndServe()
}
