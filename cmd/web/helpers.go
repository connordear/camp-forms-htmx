package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"runtime/debug"
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
