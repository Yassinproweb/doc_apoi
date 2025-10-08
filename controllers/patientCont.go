package controllers

import (
	"fmt"
	"strings"

	"github.com/Yassinproweb/doc_apoi/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Guest view — no login required
func GuestDashboardController(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, _ := s.Get(c)
		name := sess.Get("patient_name")

		// If user is already logged in, redirect to their personal dashboard
		if name != nil {
			c.Set("HX-Redirect", "/dashboard/"+name.(string))
			return c.SendStatus(fiber.StatusOK)
		}

		// Otherwise show guest dashboard
		return c.Render("dashboard", fiber.Map{
			"Guest": true,
		})
	}
}

// Authenticated patient view
func PatientDashboardController(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, _ := s.Get(c)
		email := sess.Get("patient_email")
		name := sess.Get("patient_name")

		// If session missing, fallback to guest view
		if email == nil || name == nil {
			return c.Redirect("/dashboard")
		}

		// Make sure the URL matches the session name
		paramName := c.Params("name")
		if paramName != name.(string) {
			return c.Redirect("/dashboard/" + name.(string))
		}

		// Fetch patient info
		p, err := models.GetPatientByName(name.(string))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to load dashboard")
		}

		doctors, err := models.GetAllDoctors()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to load doctors")
		}

		return c.Render("dashboard", fiber.Map{
			"Guest":   false,
			"Patient": p,
			"Doctors": doctors,
		})
	}
}

// Patient registration
func RegisterPatientController(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		age := c.FormValue("age")
		contact := c.FormValue("contact")
		district := c.FormValue("district")

		err := models.AddPatient(name, email, age, contact, district)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Registration failed: " + err.Error())
		}

		// Create session
		sess, _ := s.Get(c)
		sess.Set("patient_email", email)
		sess.Set("patient_name", name)
		sess.Save()

		nameUrl := strings.ReplaceAll(strings.ToLower(name), " ", "_")
		redirectURL := fmt.Sprintf("/dashboard/%s", nameUrl)

		c.Set("HX-Redirect", redirectURL)
		return c.SendStatus(fiber.StatusCreated)
	}
}

// Patient login
func LoginPatientController(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		contact := c.FormValue("contact")
		district := c.FormValue("district")

		p, err := models.GetPatient(email)
		if err != nil || contact != p.Contact || district != p.District {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Create session
		sess, _ := s.Get(c)
		sess.Set("patient_email", p.Email)
		sess.Set("patient_name", p.Name)
		sess.Save()

		nameUrl := strings.ReplaceAll(strings.ToLower(p.Name), " ", "_")
		redirectURL := fmt.Sprintf("/dashboard/%s", nameUrl)

		c.Set("HX-Redirect", redirectURL)
		return c.SendStatus(fiber.StatusOK)
	}
}

// View individual patient dashboard
func PatientRedirect(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")

		p, err := models.GetPatientByName(name)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Patient not found")
		}

		return c.Render("patient", fiber.Map{
			"Patient": p,
		})
	}
}

// Update patient profile
func UpdatePatientController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		email := sess.Get("patient_email")
		if email == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		name := c.FormValue("name")
		contact := c.FormValue("contact")
		district := c.FormValue("district")
		age := c.FormValue("age")

		_, err = models.EditPatient(email.(string), name, age, contact, district)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if name != "" {
			sess.Set("patient_name", name)
			sess.Save()
		}

		c.Set("HX-Redirect", "/patients/"+name)
		return c.SendStatus(fiber.StatusOK)
	}
}
