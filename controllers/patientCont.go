package controllers

import (
	"fmt"
	"strings"

	"github.com/Yassinproweb/doc_apoi/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Fetch all patients and render homepage
func GetPatientsController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		patients, err := models.GetAllPatients()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch patients")
		}

		// Render homepage with patients data
		return c.Render("dashboard", fiber.Map{
			"Patients": patients,
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
		sess.Set("patient_name", name) // save patient name too
		sess.Save()

		nameUrl := strings.ReplaceAll(strings.ToLower(name), " ", "_")
		redirectURL := fmt.Sprintf("/patients?name=%s", nameUrl)

		// Redirect for HTMX
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
		sess.Set("patient_name", p.Name) // save patient name too
		sess.Save()

		nameUrl := strings.ReplaceAll(strings.ToLower(p.Name), " ", "_")
		redirectURL := fmt.Sprintf("/patients?name=%s", nameUrl)

		c.Set("HX-Redirect", redirectURL)
		return c.SendStatus(fiber.StatusOK)
	}
}

func PatientRedirect(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")

		p, err := models.GetPatientByName(name)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Patient not found")
		}

		return c.Render("dashboard", fiber.Map{
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
			return c.SendStatus(fiber.StatusUnauthorized) // not logged in
		}

		name := c.FormValue("name")
		contact := c.FormValue("contact")
		district := c.FormValue("district")
		age := c.FormValue("age")

		_, err = models.EditPatient(email.(string), name, age, contact, district)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Update name in session if changed
		if name != "" {
			sess.Set("patient_name", name)
			sess.Save()
		}

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	}
}
