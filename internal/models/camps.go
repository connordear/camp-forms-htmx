package models

import (
	"database/sql"
	"time"
)

type Camp struct {
	ID   int
	Name string
	Year string
}

type CampModel struct {
	DB *sql.DB
}

func (m *CampModel) GetAll(year string) ([]*Camp, error) {
	sql := "SELECT c.id, c.name, cy.year FROM camps c, camp_years cy LEFT JOIN camps ON c.id = cy.camp_id;"

	if year == "" {
		year = time.Now().Format("2006")
	}

	rows, err := m.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	camps := []*Camp{}

	for rows.Next() {
		camp := &Camp{}

		err = rows.Scan(&camp.ID, &camp.Name, &camp.Year)

		camps = append(camps, camp)
	}

	return camps, nil
}
