package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func PatientAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Cookies("patient_name")
		email := c.Cookies("patient_email")

		if name == "" || email == "" {
			return c.Redirect("/patients?mode=login")
		}

		return c.Next()
	}
}

