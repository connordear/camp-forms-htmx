package main

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/connordear/camp-forms/internal/middleware"
	"github.com/connordear/camp-forms/internal/models"
	"github.com/connordear/camp-forms/ui/components"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
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
		newReg := models.Registration{
			ForUser:   1,
			FirstName: "Leon",
			LastName:  "Draisaitl",
			ForCamp: models.Camp{
				ID:   1,
				Year: "2025",
				Name: "New Camp",
			},
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

func (app *application) deleteRegistration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	regId, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		app.serverError(w, err)
	}

	app.InfoLog.Printf("Deleting Registration with ID: %d\n", regId)

	app.Registrations.Delete(regId)
}

func (app *application) home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		regs, err := app.Registrations.GetAll(1, 2025)
		if err != nil {
			app.ErrorLog.Fatal(err)
		}
		components.HomePage(regs).Render(r.Context(), w)
		// app.page("home.tmpl", regs)(w, r)

	}
}

func (app *application) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		components.Hello("John").Render(r.Context(), w)
	}
}

func (app *application) router() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home())
	router.HandlerFunc(http.MethodGet, "/reset", app.page("reset.tmpl", nil))
	router.HandlerFunc(http.MethodGet, "/camps", app.getCamps())
	router.Handler(http.MethodGet, "/hello", templ.Handler(components.Hello("John")))
	router.HandlerFunc(http.MethodDelete, "/all", app.deleteAll())
	router.HandlerFunc(http.MethodPost, "/registrations", app.createRegistration())
	router.DELETE("/registrations/:id", app.deleteRegistration)

	standard := alice.New(middleware.RecoverPanic(app.ErrorLog), middleware.Logging(app.InfoLog), middleware.SecureHeaders)

	return standard.Then(router)
}
