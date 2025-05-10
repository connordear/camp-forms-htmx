package models

import "database/sql"

type Registration struct {
	ID int
}

type RegistrationModel struct {
	DB *sql.DB
}

func (m *RegistrationModel) 
