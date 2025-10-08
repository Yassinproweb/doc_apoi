package middlewares

import "github.com/gofiber/fiber/v2"

func DoctorAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Cookies("doctor_email")
		if email == "" {
			return c.Redirect("/doctors?mode=login")
		}
		return c.Next()
	}
}

