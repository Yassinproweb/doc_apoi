package routes

import (
	"github.com/Yassinproweb/doc_apoi/controllers"
	"github.com/Yassinproweb/doc_apoi/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DocRoutes(app *fiber.App, s *session.Store) {
	// Share session store with controllers
	controllers.Store = s

	app.Get("/doctors/:name", middlewares.DoctorAuth(s), controllers.DoctorRedirect(s))

	// Doctor form routes
	app.Get("/doctors", func(c *fiber.Ctx) error {
		mode := c.Query("mode", "register")
		return c.Render("forms/doctors", fiber.Map{
			"Mode": mode,
		})
	})

	// Doctor auth routes
	app.Post("/doctors/register", controllers.RegisterDoctorController(s))
	app.Post("/doctors/login", controllers.LoginDoctorController(s))

	// Doctor profile management
	app.Get("/doctors/:name/edit", controllers.EditDoctorFormController())
	app.Post("/doctors/:name/update", controllers.UpdateDoctorController())

	// Logout route
	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		if err := sess.Destroy(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to logout")
		}
		return c.Redirect("/doctors?mode=login")
	})
}
