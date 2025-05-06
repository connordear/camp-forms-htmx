package config

import (
	"log"

	"github.com/connordear/camp-forms/internal/models"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Camps    *models.CampModel
}
