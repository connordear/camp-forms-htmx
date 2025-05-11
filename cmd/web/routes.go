package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/connordear/camp-forms/internal/middleware"
)

func page(app *application, pageName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fullPath := "./ui/html/pages/" + pageName

		if _, err := os.Stat(fullPath); err != nil {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./ui/html/base.tmpl",
			fullPath,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.serverError(w, err)
		}
	}
}

func getCamps(app *application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		camps, err := app.Camps.GetAll("2025")
		if err != nil {
			app.ErrorLog.Fatal(err)
		}

		for _, camp := range camps {
			app.InfoLog.Println(camp.Name)
		}

	}
}

func deleteAll(app *application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.InfoLog.Println("Resetting DB...")
		err := app.Meta.RunMigrations()
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/", 303)
	}
}

func Router(app *application) *http.ServeMux {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.Handle("GET /", middleware.Logging(page(app, "home.tmpl"), app.InfoLog))
	router.Handle("GET /reset", middleware.Logging(page(app, "reset.tmpl"), app.InfoLog))
	router.Handle("GET /camps", middleware.Logging(getCamps(app), app.InfoLog))
	router.Handle("DELETE /all", middleware.Logging(deleteAll(app), app.InfoLog))

	return router
}
