package models

import (
	"errors"

	"github.com/Yassinproweb/TeleMedi/data"
	"golang.org/x/crypto/bcrypt"
)

type Doctor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Skill    string `json:"skill"`
	Title    string `json:"title"`
}

func GetAllDoctors() ([]Doctor, error) {
	rows, err := data.DB.Query(`SELECT * FROM doctors`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var d Doctor

		if err := rows.Scan(&d.Name, &d.Email, &d.Password, &d.Skill, &d.Title); err != nil {
			return nil, err
		}

		doctors = append(doctors, d)
	}

	return doctors, err
}

func AddDoctor(name, email, password, skill, title string) error {
	_, err := GetDoctor(email)
	if err != nil {
		return errors.New(`Email already taken`)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insert doctor into data
	_, err = data.DB.Exec(`INSERT INTO doctors (name, email, password, title, skill) VALUES (?, ?, ?, ?, ?)`, name, email, string(hashedPassword), title, skill)

	return err
}

func GetDoctor(email string) (*Doctor, error) {
	row := data.DB.QueryRow(`SELECT id, name, password FROM doctors WHERE email = ?`, email)

	var d Doctor
	err := row.Scan(&d.Name, &d.Password)
	if err != nil {
		return nil, err
	}

	return &d, err
}

func GetDoctorByEmail(email string, d *Doctor) error {
	return data.DB.QueryRow(`SELECT name, email, skill, title FROM doctors WHERE email=?`, email).
		Scan(&d.Name, &d.Email, &d.Skill, &d.Title)
}

func GetDoctorByName(name string) (*Doctor, error) {
	row := data.DB.QueryRow(`SELECT * FROM doctors WHERE name = ?`, name)

	var d Doctor

	err := row.Scan(&d.Name, &d.Email, &d.Password, &d.Skill, &d.Title)
	if err != nil {
		return nil, err
	}

	return &d, err
}

func EditDoctor(email, name, password, skill, title string) (*Doctor, error) {
	// Get current record
	var d Doctor
	err := data.DB.QueryRow(`SELECT name, email, password, skill, title FROM doctors WHERE email=?`, email).
		Scan(&d.Name, &d.Email, &d.Password, &d.Skill, &d.Title)
	if err != nil {
		return nil, err
	}

	// Only update fields if provided
	if name != "" {
		d.Name = name
	}
	if skill != "" {
		d.Skill = skill
	}
	if title != "" {
		d.Title = title
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		d.Password = string(hashedPassword)
	}

	// Save changes
	_, err = data.DB.Exec(`UPDATE doctors SET name=?, password=?, skill=?, title=? WHERE email=?`,
		d.Name, d.Password, d.Skill, d.Title, d.Email)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (d *Doctor) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(password))
	return err == nil
}
