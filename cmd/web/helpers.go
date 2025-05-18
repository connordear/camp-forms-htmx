package main

import (
	"fmt"
	"net/http"
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

func page(app *application, pageName string, data any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ts, ok := app.TemplateCache[pageName]
		if !ok {
			err := fmt.Errorf("the template %s does not exist", pageName)
			app.serverError(w, err)
			return
		}

		err := ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.serverError(w, err)
		}
	}
}
