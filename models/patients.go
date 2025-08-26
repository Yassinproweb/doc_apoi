package models

type Patient struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
