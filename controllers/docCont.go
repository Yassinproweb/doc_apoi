package controllers

import (
	"fmt"
	"strings"

	"github.com/Yassinproweb/doc_apoi/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

// Fetch all doctors and render homepage
func GetDoctorsController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		doctors, err := models.GetAllDoctors()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch doctors")
		}

		// Render homepage with doctors data
		return c.Render("dashboard", fiber.Map{
			"Doctors": doctors,
		})
	}
}

// Doctor registration
func RegisterDoctorController(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		skill := c.FormValue("skill")
		title := c.FormValue("title")
		location := c.FormValue("location")

		err := models.AddDoctor(name, email, password, skill, title, location)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Registration failed: " + err.Error())
		}

		// Create session
		sess, _ := s.Get(c)
		sess.Set("doctor_email", email)
		sess.Set("doctor_name", name) // save doctor name too
		sess.Save()

		nameUrl := strings.ReplaceAll(strings.ToLower(name), " ", "_")
		redirectURL := fmt.Sprintf("/doctors?name=%s", nameUrl)

		// Redirect for HTMX
		c.Set("HX-Redirect", redirectURL)
		return c.SendStatus(fiber.StatusCreated)
	}
}

// Doctor login
func LoginDoctorController(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		d, err := models.GetDoctor(email)
		if err != nil || !d.CheckPassword(password) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Create session
		sess, _ := s.Get(c)
		sess.Set("doctor_email", d.Email)
		sess.Set("doctor_name", d.Name) // save doctor name too
		sess.Save()

		nameUrl := strings.ReplaceAll(strings.ToLower(d.Name), " ", "_")
		redirectURL := fmt.Sprintf("/doctors?name=%s", nameUrl)

		c.Set("HX-Redirect", redirectURL)
		return c.SendStatus(fiber.StatusOK)
	}
}

func DoctorRedirect(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")

		d, err := models.GetDoctorByName(name)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Doctor not found")
		}

		return c.Render("dr-profile", fiber.Map{
			"Doctor": d,
		})
	}
}

// Update doctor profile
func UpdateDoctorController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		email := sess.Get("doctor_email")
		if email == nil {
			return c.SendStatus(fiber.StatusUnauthorized) // not logged in
		}

		name := c.FormValue("name")
		skill := c.FormValue("skill")
		title := c.FormValue("title")
		location := c.FormValue("location")
		password := c.FormValue("password")

		_, err = models.EditDoctor(email.(string), name, password, skill, title, location)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Update name in session if changed
		if name != "" {
			sess.Set("doctor_name", name)
			sess.Save()
		}

		// Redirect for HTMX
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	}
}

// Get edit doctor form
func EditDoctorFormController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
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

		return c.Render("update-doctor", fiber.Map{"Doctor": d})
	}
}
