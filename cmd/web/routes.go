package main

import (
	"net/http"

	"github.com/connordear/camp-forms/internal/middleware"
	"github.com/connordear/camp-forms/internal/models"
)

func (app *application) getCamps() http.HandlerFunc {
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

func (app *application) deleteAll() http.HandlerFunc {
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

func (app *application) createRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			app.clientError(w, http.StatusMethodNotAllowed)
			return
		}

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

		app.render(w, "registration.tmpl", "registration", newReg)
	}

}

func (app *application) router() *http.ServeMux {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /", middleware.Logging(app.page("home.tmpl", nil), app.InfoLog))
	router.HandleFunc("GET /reset", middleware.Logging(app.page("reset.tmpl", nil), app.InfoLog))
	router.HandleFunc("GET /camps", middleware.Logging(app.getCamps(), app.InfoLog))
	router.HandleFunc("DELETE /all", middleware.Logging(app.deleteAll(), app.InfoLog))
	router.HandleFunc("POST /registrations", middleware.Logging(app.createRegistration(), app.InfoLog))

	return router
}
