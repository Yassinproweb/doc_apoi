package models

import (
	"errors"

	"github.com/Yassinproweb/doc_apoi/data"
	"golang.org/x/crypto/bcrypt"
)

type Doctor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Skill    string `json:"skill"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Avatar   string `json:"avatar"`
}

// Default avatar path
const DefaultAvatar = "/static/imgs/pngs/doc_apoi.png"

// Fetch all doctors
func GetAllDoctors() ([]Doctor, error) {
	rows, err := data.DB.Query(`SELECT name, email, password, skill, title, location, avatar FROM doctors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var d Doctor

		if err := rows.Scan(&d.Name, &d.Email, &d.Password, &d.Skill, &d.Title, &d.Location, &d.Avatar); err != nil {
			return nil, err
		}

		if d.Avatar == "" {
			d.Avatar = DefaultAvatar
		}

		doctors = append(doctors, d)
	}

	return doctors, nil
}

// Add new doctor
func AddDoctor(name, email, password, skill, title, location, avatar string) error {
	// Check if doctor already exists
	_, err := GetDoctor(email)
	if err == nil {
		return errors.New(`Email already taken`)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insert doctor into table
	_, err = data.DB.Exec(
		`INSERT INTO doctors (name, email, password, skill, title, location, avatar) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		name, email, string(hashedPassword), skill, title, location, avatar,
	)
	return err
}

// Fetch doctor by email (for login)
func GetDoctor(email string) (*Doctor, error) {
	row := data.DB.QueryRow(
		`SELECT name, email, password, skill, title, location, avatar FROM doctors WHERE email = ?`,
		email,
	)

	var d Doctor
	err := row.Scan(&d.Name, &d.Email, &d.Password, &d.Skill, &d.Title, &d.Location, &d.Avatar)
	if err != nil {
		return nil, err
	}

	if d.Avatar == "" {
		d.Avatar = DefaultAvatar
	}

	return &d, nil
}

// Fetch doctor by name
func GetDoctorByName(name string) (*Doctor, error) {
	row := data.DB.QueryRow(
		`SELECT name, email, password, skill, title, location, avatar FROM doctors WHERE name = ?`,
		name,
	)

	var d Doctor
	err := row.Scan(&d.Name, &d.Email, &d.Password, &d.Skill, &d.Title, &d.Location, &d.Avatar)
	if err != nil {
		return nil, err
	}

	if d.Avatar == "" {
		d.Avatar = DefaultAvatar
	}

	return &d, nil
}

// Fetch doctor by email (into provided struct)
func GetDoctorByEmail(email string, d *Doctor) error {
	err := data.DB.QueryRow(
		`SELECT name, email, password, skill, title, location, avatar FROM doctors WHERE email = ?`,
		email,
	).Scan(&d.Name, &d.Email, &d.Password, &d.Skill, &d.Title, &d.Location, &d.Avatar)
	if err != nil {
		return err
	}

	if d.Avatar == "" {
		d.Avatar = DefaultAvatar
	}

	return nil
}

// Edit doctor profile (name, password, etc.)
func EditDoctor(email, name, password, skill, title, location string) (*Doctor, error) {
	// Get current record
	var d Doctor
	err := data.DB.QueryRow(
		`SELECT name, email, password, skill, title, location FROM doctors WHERE email=?`,
		email,
	).Scan(&d.Name, &d.Email, &d.Password, &d.Skill, &d.Title, &d.Location)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if name != "" {
		d.Name = name
	}
	if skill != "" {
		d.Skill = skill
	}
	if title != "" {
		d.Title = title
	}
	if location != "" {
		d.Location = location
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		d.Password = string(hashedPassword)
	}

	// Save changes
	_, err = data.DB.Exec(
		`UPDATE doctors SET name=?, password=?, skill=?, title=?, location=? WHERE email=?`,
		d.Name, d.Password, d.Skill, d.Title, d.Location, d.Email,
	)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// ✅ Update doctor avatar using their email
func UpdateDoctorAvatar(email, avatarPath string) error {
	_, err := data.DB.Exec(`UPDATE doctors SET avatar = ? WHERE email = ?`, avatarPath, email)
	return err
}

// Compare hashed password with plain text
func (d *Doctor) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(password))
	return err == nil
}

