package controllers

import (
	"github.com/Yassinproweb/TeleMedi/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func RegisterDoctor(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		skill := c.FormValue("skill")
		title := c.FormValue("title")

		err := models.AddDoctor(name, email, password, skill, title)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		// Create session
		sess, _ := s.Get(c)
		sess.Set("doctor_name", name)
		sess.Save()

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusCreated)
	}
}

func LoginDoctor(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		d, err := models.GetDoctor(email)
		if err != nil || !d.CheckPassword(password) {
			return c.SendStatus(fiber.StatusNotFound)
		}

		// Create session
		sess, _ := s.Get(c)
		sess.Set("doctor_email", d.Email)

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	}
}

func UpdateDoctor(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		email := sess.Get("doctor_email")
		if email == nil {
			return c.SendStatus(fiber.StatusUnauthorized) // not logged in
		}

		name := c.FormValue("name")
		skill := c.FormValue("skill")
		title := c.FormValue("title")
		password := c.FormValue("password")

		_, err = models.EditDoctor(email.(string), name, password, skill, title)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	}
}

// Get edit doctor form
func EditDoctorForm(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		email := sess.Get("doctor_email")
		if email == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Fetch doctor (without password)
		var d models.Doctor
		err = models.GetDoctorByEmail(email.(string), &d)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.Render("update", fiber.Map{"Doctor": d})
	}
}
