package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func PatientAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Cookies("patient_email")

		if email == "" {
			return c.Redirect("/patients?mode=login")
		}

		return c.Next()
	}
}
