package handler

import (
	"github.com/connordear/camp-forms/internal/config"
	"github.com/connordear/camp-forms/internal/middleware"
	"net/http"
)

func Routes(app *config.Application) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /camps", middleware.Logging(getCamps(app), app.InfoLog))

	return router
}
