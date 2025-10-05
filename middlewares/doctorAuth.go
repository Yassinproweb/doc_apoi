package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DoctorAuth(s *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}

		if sess.Get("doctor_name") == nil {
			return c.Redirect("/doctors/login")
		}

		return c.Next()
	}
}
