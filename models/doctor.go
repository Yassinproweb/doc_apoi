package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Doctor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Skill    string `json:"skill"`
	Title    string `json:"title"`
}

// HashPassword hashes the doctor's password
func (d *Doctor) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	d.Password = string(hashed)
	return nil
}

// VerifyPassword checks if the provided password matches the stored hash
func (d *Doctor) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(password))
}
