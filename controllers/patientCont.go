package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/Yassinproweb/doc_apoi/models"
	"github.com/gofiber/fiber/v2"
)

// Guest view — no login required
func GuestDashboardController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("dashboard", fiber.Map{
			"Guest": true,
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
			Expires:  time.Now().Add(24 * time.Hour),
		})

		nameUrl := strings.ReplaceAll(strings.ToLower(name), " ", "_")
		redirectURL := fmt.Sprintf("/dashboard/%s", nameUrl)

		c.Set("HX-Redirect", redirectURL)
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
			Expires:  time.Now().Add(24 * time.Hour),
		})

		nameUrl := strings.ReplaceAll(strings.ToLower(p.Name), " ", "_")
		redirectURL := fmt.Sprintf("/dashboard/%s", nameUrl)

		c.Set("HX-Redirect", redirectURL)
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

		c.Set("HX-Redirect", "/dashboard/"+name)
		return c.SendStatus(fiber.StatusOK)
	}
}

// Logout
func LogoutPatientController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.ClearCookie("patient_email")
		c.ClearCookie("patient_name")
		return c.Redirect("/patients?mode=login")
	}
}
