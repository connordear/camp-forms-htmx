package main

import (
	"bytes"
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

// Simple page rendering convenience helper for now, probably will
// need a custom handler that can just call `render`
func (app *application) page(pageName string, data any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.render(w, pageName, "base", data)
	}
}

func (app *application) render(w http.ResponseWriter, templateFileName string, templateName string, data any) {

	ts, ok := app.TemplateCache[templateFileName]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", templateFileName)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	// Write to buffer first to check if there are any errors
	err := ts.ExecuteTemplate(buf, templateName, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}
