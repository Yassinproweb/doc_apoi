package routes

import (
	"github.com/Yassinproweb/TeleMedi/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DocRoutes(app *fiber.App, s *session.Store) {
	// Share session store with controllers
	controllers.Store = s

	// Homepage
	app.Get("/", controllers.GetDoctorsController())

	// Doctor auth routes
	app.Post("/doctors/register", controllers.RegisterDoctorController(s))
	app.Post("/doctors/login", controllers.LoginDoctorController(s))

	// Doctor profile management
	app.Get("/doctor/edit", controllers.EditDoctorFormController())
	app.Post("/doctor/update", controllers.UpdateDoctorController())

	// Authentication middleware
	isAuthenticated := func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}

		// Check session using the same key as controllers
		doctorEmail := sess.Get("doctor_email")
		if doctorEmail == nil {
			return c.Redirect("/signin") // or "/login_doctor"
		}

		return c.Next()
	}

	// Protected dashboard route
	app.Get("/dashboard", isAuthenticated, func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}

		// Use doctor_email for lookup, but show doctor_name if stored
		name, _ := sess.Get("doctor_name").(string)
		if name == "" {
			// fallback if name wasn't saved in session
			name = sess.Get("doctor_email").(string)
		}

		return c.Render("dashboard", fiber.Map{
			"Name": name,
		})
	})

	// Logout route
	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		if err := sess.Destroy(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to logout")
		}
		return c.Redirect("/signin")
	})
}
