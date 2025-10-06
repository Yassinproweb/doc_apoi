package models

import (
	"errors"

	"github.com/Yassinproweb/doc_apoi/data"
)

type Patient struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      string `json:"age"`
	Contact  string `json:"contact"`
	District string `json:"district"`
}

// Fetch all patients
func GetAllPatients() ([]Patient, error) {
	rows, err := data.DB.Query(`SELECT * FROM patients`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var p Patient
		if err := rows.Scan(&p.Name, &p.Email, &p.Age, &p.Contact, &p.District); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, nil
}

// Add new doctor
func AddPatient(name, email, age, contact, district string) error {
	// Check if doctor already exists
	_, err := GetPatient(email)
	if err == nil {
		return errors.New(`Email already taken`)
	}

	// Insert doctor into table
	_, err = data.DB.Exec(
		`INSERT INTO patients (name, email, age, contact, district) VALUES (?, ?, ?, ?, ?, ?)`,
		name, email, age, contact, district,
	)
	return err
}

// Fetch doctor by email (for login)
func GetPatient(email string) (*Patient, error) {
	row := data.DB.QueryRow(
		`SELECT name, email, age, contact, district FROM patients WHERE email = ?`,
		email,
	)

	var p Patient
	err := row.Scan(&p.Name, &p.Email, &p.Age, &p.Contact, &p.District)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Fetch doctor by email into struct (for editing profile)
func GetPatientByEmail(email string, p *Patient) error {
	return data.DB.QueryRow(
		`SELECT name, email, contact, district FROM patients WHERE email=?`,
		email,
	).Scan(&p.Name, &p.Email, &p.Contact, &p.District)
}

// Fetch doctor by name
func GetPatientByName(name string) (*Patient, error) {
	row := data.DB.QueryRow(
		`SELECT name, email, age, contact, district FROM patients WHERE name = ?`,
		name,
	)

	var p Patient
	err := row.Scan(&p.Name, &p.Email, &p.Age, &p.Contact, &p.District)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Edit doctor profile
func EditPatient(email, name, age, contact, district string) (*Patient, error) {
	// Get current record
	var p Patient
	err := data.DB.QueryRow(
		`SELECT name, email, age, contact, district FROM patients WHERE email=?`,
		email,
	).Scan(&p.Name, &p.Email, &p.Age, &p.Contact, &p.District)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if name != "" {
		p.Name = name
	}
	if age != "" {
		p.Age = age
	}
	if contact != "" {
		p.Contact = contact
	}
	if district != "" {
		p.District = district
	}

	// Save changes
	_, err = data.DB.Exec(
		`UPDATE patients SET name=?, age=?, contact=?, district=? WHERE email=?`,
		p.Name, p.Age, p.Contact, p.District, p.Email,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
