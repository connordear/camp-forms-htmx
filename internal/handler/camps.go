package handler

import (
	"net/http"

	"github.com/connordear/camp-forms/internal/config"
	_ "github.com/glebarez/go-sqlite"
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
