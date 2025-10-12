package controllers

import (
	"github.com/Yassinproweb/doc_apoi/models"
	"github.com/Yassinproweb/doc_apoi/utils"
	"github.com/gofiber/fiber/v2"
)

// Guest view — no login required
func GuestDashboardController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Cookies("patient_email")
		if email != "" {
			p, err := models.GetPatient(email)
			if err == nil && p.Name != "" {
				c.Set("HX-Redirect", utils.URLer("dashboard", utils.NormalizeName(p.Name)))
				return c.SendStatus(fiber.StatusOK)
			}
		}

		// Render guest dashboard
		doctors, _ := models.GetAllDoctors()
		return c.Render("dashboard", fiber.Map{
			"Guest":   true,
			"Doctors": doctors,
		})
	}
}

// Patient view — login required
func PatientDashboardController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paramName := c.Params("name")
		email := c.Cookies("patient_email")

		if email == "" {
			c.Set("HX-Redirect", "/dashboard")
			return c.SendStatus(fiber.StatusOK)
		}

		p, err := models.GetPatient(email)
		if err != nil {
			c.Set("HX-Redirect", "/dashboard")
			return c.SendStatus(fiber.StatusOK)
		}

		eSlug := utils.NormalizeName(p.Name)
		if eSlug != paramName {
			c.Set("HX-Redirect", utils.URLer("dashboard", eSlug))
			return c.SendStatus(fiber.StatusOK)
		}

		doctors, err := models.GetAllDoctors()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to load doctors")
		}

		return c.Render("dashboard", fiber.Map{
			"Guest":    false,
			"Name":     p.Name,
			"District": p.District,
			"Contact":  p.Contact,
			"Doctors":  doctors,
		})
	}
}

// Patient registration
func RegisterPatientController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		age := c.FormValue("age")
		contact := c.FormValue("contact")
		district := c.FormValue("district")

		if err := models.AddPatient(name, email, age, contact, district); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Registration failed: " + err.Error())
		}

		c.Cookie(&fiber.Cookie{
			Name:     "patient_email",
			Value:    email,
			HTTPOnly: true,
			SameSite: "Lax", // protect against CSRF
			Secure:   false, // only send over HTTPS (use false in localhost)
			MaxAge:   60 * 60 * 24 * 7,
		})

		c.Set("HX-Redirect", utils.URLer("dashboard", utils.NormalizeName(name)))
		return c.SendStatus(fiber.StatusCreated)
	}
}

// Patient login
func LoginPatientController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		contact := c.FormValue("contact")
		district := c.FormValue("district")

		p, err := models.GetPatient(email)
		if err != nil || contact != p.Contact || district != p.District {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "patient_email",
			Value:    p.Email,
			HTTPOnly: true,
			SameSite: "Lax", // protect against CSRF
			Secure:   false, // only send over HTTPS (use false in localhost)
			MaxAge:   60 * 60 * 24 * 7,
		})

		n := utils.NormalizeName(p.Name)
		c.Set("HX-Redirect", utils.URLer("dashboard", n))
		return c.SendStatus(fiber.StatusOK)
	}
}

// Update patient profile
func UpdatePatientController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Cookies("patient_email")
		if email == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		name := c.FormValue("name")
		contact := c.FormValue("contact")
		district := c.FormValue("district")
		age := c.FormValue("age")

		_, err := models.EditPatient(email, name, age, contact, district)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update profile")
		}

		c.Set("HX-Redirect", utils.URLer("dashboard", utils.NormalizeName(name)))
		return c.SendStatus(fiber.StatusOK)
	}
}

// Logout
func LogoutPatientController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.ClearCookie("patient_email")
		return c.Redirect("/patients?mode=login")
	}
}
