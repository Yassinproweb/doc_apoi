package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func PatientAuth(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}

		if sess.Get("patient_name") == nil {
			return c.Redirect("/patients/login")
		}

		return c.Next()
	}
}
