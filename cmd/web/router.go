package main

import (
	"html/template"
	"net/http"

	"github.com/connordear/camp-forms/internal/config"
	"github.com/connordear/camp-forms/internal/middleware"
)

func home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Use the template.ParseFiles() function to read the template file into a
		// template set. If there's an error, we log the detailed error message and use
		// the http.Error() function to send a generic 500 Internal Server Error
		// response to the user.
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/pages/home.tmpl",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		// We then use the Execute() method on the template set to write the
		// template content as the response body. The last parameter to Execute()
		// represents any dynamic data that we want to pass in, which for now we'll
		// leave as nil.
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func getCamps(app *config.Application) http.HandlerFunc {
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

func Router(app *config.Application) *http.ServeMux {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.Handle("GET /", middleware.Logging(home(app), app.InfoLog))
	router.Handle("GET /camps", middleware.Logging(getCamps(app), app.InfoLog))

	return router
}
