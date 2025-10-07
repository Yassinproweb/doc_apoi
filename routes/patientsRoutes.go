package routes

import (
	"github.com/Yassinproweb/doc_apoi/controllers"
	"github.com/Yassinproweb/doc_apoi/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func PatRoutes(app *fiber.App, s *session.Store) {
	// Share session store with controllers
	controllers.Store = s

	app.Get("/dashboard", controllers.GuestDashboardController(s))
	app.Get("/dashboard/:name", controllers.PatientDashboardController(s))
	app.Get("/dashboard/:name", middlewares.PatientAuth(s), controllers.PatientRedirect(s))

	// Patient form routes
	app.Get("/patients", func(c *fiber.Ctx) error {
		mode := c.Query("mode", "register")
		return c.Render("forms/patients", fiber.Map{
			"Mode": mode,
		})
	})

	// Patient auth routes
	app.Post("/patients/register", controllers.RegisterPatientController(s))
	app.Post("/patients/login", controllers.LoginPatientController(s))

	// Patient profile management
	// app.Get("/patients/:name/edit", controllers.EditPatientFormController())
	app.Post("/dashboard/:name/update", controllers.UpdatePatientController())

	// Logout route
	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		if err := sess.Destroy(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to logout")
		}
		return c.Redirect("/patients?mode=login")
	})
}
