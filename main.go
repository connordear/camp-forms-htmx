package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/connordear/camp-forms/internal/config"
	"github.com/connordear/camp-forms/internal/db"
	"github.com/connordear/camp-forms/internal/handler"
	"github.com/connordear/camp-forms/internal/models"
)

func main() {
	port := flag.String("port", ":4000", "HTTP Port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	database := db.InitDatabase(infoLog, errorLog)

	app := &config.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Camps:    &models.CampModel{DB: database},
	}

	server := http.Server{
		Addr:    *port,
		Handler: handler.Routes(app),
	}

	app.InfoLog.Println("Listening on port ", *port)
	server.ListenAndServe()
}
