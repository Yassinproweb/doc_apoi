package routes

import (
	"github.com/Yassinproweb/doc_apoi/controllers"
	"github.com/Yassinproweb/doc_apoi/middlewares"
	"github.com/gofiber/fiber/v2"
)

func DocRoutes(app *fiber.App) {
	app.Get("/doctors/:name", middlewares.DoctorAuth(), controllers.DoctorRedirect())

	// Doctor form
	app.Get("/doctors", func(c *fiber.Ctx) error {
		mode := c.Query("mode", "register")
		return c.Render("forms/doctors", fiber.Map{
			"Mode": mode,
		})
	})

	// Auth
	app.Post("/doctors/register", controllers.RegisterDoctorController())
	app.Post("/doctors/login", controllers.LoginDoctorController())

	// Profile management
	app.Get("/doctors/:name/edit", middlewares.DoctorAuth(), controllers.EditDoctorFormController())
	app.Post("/doctors/:name/update", middlewares.DoctorAuth(), controllers.UpdateDoctorController())

	// Logout
	app.Get("/logout", controllers.LogoutDoctorController())
}
