package models

import "database/sql"

type Registration struct {
	ID        int
	ForUser   int
	FirstName string
	LastName  string
	ForCamp   Camp
}

type RegistrationModel struct {
	DB *sql.DB
}

func (m *RegistrationModel) Get(regId int) (*Registration, error) {
	sql := `SELECT
			r.id,
			r.for_camp,
			r.camp_year,
			c.name,
			r.first_name,
			r.last_name
		FROM
			registrations r
			LEFT JOIN camp_years cy ON (r.for_camp = cy.camp_id)
			LEFT JOIN camps c ON (r.for_camp = c.id)
		WHERE
			r.id = ?
	`
	row := m.DB.QueryRow(sql, regId)

	r := &Registration{}

	err := row.Scan(&r.ID, &r.ForCamp.ID, &r.ForCamp.Year, &r.ForCamp.Name, &r.FirstName, &r.LastName)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (m *RegistrationModel) GetAll(userId int, year int) ([]*Registration, error) {
	sql := `SELECT
			r.id,
			r.for_camp,
			r.camp_year,
			c.name,
			r.first_name,
			r.last_name
		FROM
			registrations r
			LEFT JOIN camp_years cy ON (r.for_camp = cy.camp_id)
			LEFT JOIN camps c ON (r.for_camp = c.id)
		WHERE
			r.for_user = ?
			AND cy.year = ?
	`
	rows, err := m.DB.Query(sql, userId, year)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	registrations := []*Registration{}

	for rows.Next() {
		r := &Registration{}
		err := rows.Scan(
			&r.ID,
			&r.ForCamp.ID,
			&r.ForCamp.Year,
			&r.ForCamp.Name,
			&r.FirstName,
			&r.LastName,
		)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, r)
	}
	return registrations, nil

}

func (m *RegistrationModel) Add(reg *Registration) (int, error) {
	sql := `INSERT INTO registrations (for_user, for_camp, camp_year, first_name, last_name)
	VALUES (?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(sql, reg.ForUser, reg.ForCamp.ID, reg.ForCamp.Year, reg.FirstName, reg.LastName)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m *RegistrationModel) Delete(id int) error {
	sql := `DELETE FROM registrations WHERE id = ?`

	_, err := m.DB.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}
