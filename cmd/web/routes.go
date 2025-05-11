package main

import (
	"html/template"
	"net/http"

	"github.com/connordear/camp-forms/internal/middleware"
	"github.com/connordear/camp-forms/internal/models"
)

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

func createRegistration(app *application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			app.clientError(w, http.StatusMethodNotAllowed)
			return
		}

		//fName := "John"
		//lName := "Doe"

		newReg := models.Registration{
			ForCamp:  1,
			CampYear: 2025,
		}
		id, err := app.Registrations.Add(&newReg)
		newReg.ID = id

		if err != nil {
			app.serverError(w, err)
			return
		}
		tmplFile := "./ui/html/templates/registration.tmpl"
		tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
		if err != nil {
			app.serverError(w, err)
		}

		err = tmpl.ExecuteTemplate(w, "registration", newReg)
		if err != nil {
			app.serverError(w, err)
		}

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
	router.Handle("POST /registrations", middleware.Logging(createRegistration(app), app.InfoLog))

	return router
}
