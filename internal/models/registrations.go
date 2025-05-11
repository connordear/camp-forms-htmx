package models

import "database/sql"

type Registration struct {
	ID       int
	ForCamp  int
	CampYear int
}

type RegistrationModel struct {
	DB *sql.DB
}

func (m *RegistrationModel) Add(reg Registration) (int, error) {
	sql := `INSERT INTO registrations (for_camp, camp_year)
	VALUES (?, ?)`

	result, err := m.DB.Exec(sql, reg.ForCamp, reg.CampYear)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}
