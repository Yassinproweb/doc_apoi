package routes

import (
	"github.com/Yassinproweb/doc_apoi/controllers"
	"github.com/gofiber/fiber/v2"
)

func PatRoutes(app *fiber.App) {
	app.Get("/dashboard", controllers.GuestDashboardController())
	app.Get("/dashboard", controllers.GetDoctorsController())

	// Patient form route
	app.Get("/patients", func(c *fiber.Ctx) error {
		mode := c.Query("mode", "register")
		return c.Render("forms/patients", fiber.Map{
			"Mode": mode,
		})
	})

	// Auth routes
	app.Post("/patients/register", controllers.RegisterPatientController())
	app.Post("/patients/login", controllers.LoginPatientController())

	// Profile update
	app.Post("/dashboard/:name/update", controllers.UpdatePatientController())

	// Logout
	app.Get("/logout", controllers.LogoutPatientController())
}
