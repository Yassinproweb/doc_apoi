package controllers

import (
	"database/sql"

	"github.com/Yassinproweb/TeleMedi/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPatient(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Error hashing password</p>")
		}

		// Insert doctor into data
		result, err := data.DB.Exec("INSERT INTO patients (name, email, password) VALUES (?, ?, ?)", name, email, hashedPassword)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("<p class='text-red-500 text-center'>Email already exists</p>")
		}

		// Get doctor ID
		patientID, _ := result.LastInsertId()

		// Create session
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Session error</p>")
		}
		sess.Set("patient_id", patientID)
		sess.Set("patient_name", name)
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Failed to save session</p>")
		}

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	}
}

func LoginPatient(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		var storedPassword, name string
		var patientID int
		err := data.DB.QueryRow("SELECT id, name, password FROM patients(where email = ?", email).Scan(&patientID, &name, &storedPassword)
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).SendString("<p class='text-red-500 text-center'>Invalid email or password</p>")
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>data error</p>")
		}

		// Verify password
		if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("<p class='text-red-500 text-center'>Invalid email or password</p>")
		}

		// Create session
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Session error</p>")
		}
		sess.Set("patient_id", patientID)
		sess.Set("patient_name", name)
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("<p class='text-red-500 text-center'>Failed to save session</p>")
		}

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	}
}
