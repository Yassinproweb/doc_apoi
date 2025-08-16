package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Patient struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// HashPassword hashes the patient's password
func (p *Patient) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.Password = string(hashed)
	return nil
}

// VerifyPassword checks if the provided password matches the stored hash
func (p *Patient) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
}
