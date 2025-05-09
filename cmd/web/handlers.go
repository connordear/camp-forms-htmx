package main

import (
	"github.com/connordear/camp-forms/internal/config"
	"github.com/connordear/camp-forms/internal/middleware"
	"net/http"
)

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
	// router.Handle("GET /", middleware.Logging(getCamps(app), app.InfoLog))
	router.Handle("GET /camps", middleware.Logging(getCamps(app), app.InfoLog))

	return router
}
