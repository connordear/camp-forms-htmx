package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	database "github.com/connordear/camp-forms/internal/db"
	"github.com/connordear/camp-forms/internal/models"
)

type application struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	TemplateCache map[string]*template.Template
	Camps         *models.CampModel
	Meta          *models.MetaModel
	Registrations *models.RegistrationModel
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

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		TemplateCache: templateCache,
		Camps:         &models.CampModel{DB: db},
		Meta:          &models.MetaModel{DB: db},
		Registrations: &models.RegistrationModel{DB: db},
	}

	if err = app.Meta.InitDatabase(infoLog); err != nil {
		errorLog.Fatal(err)
	}

	server := http.Server{
		Addr:     *port,
		Handler:  app.router(),
		ErrorLog: errorLog,
	}

	app.InfoLog.Println("Listening on port ", *port)
	server.ListenAndServe()
}
