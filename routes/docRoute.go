package routes

import (
	"github.com/Yassinproweb/TeleMedi/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DocRoutes(app *fiber.App) {
	s := session.New()
	controllers.Store = s

	app.Get("/", controllers.GetDoctors())
	app.Post("/doctors", controllers.RegisterDoctor(s))
	app.Post("/doctors", controllers.LoginDoctor(s))
	app.Post("/patients", controllers.RegisterPatient(s))
	app.Post("/patients", controllers.LoginPatient(s))
	app.Get("/doctor/edit", controllers.EditDoctorForm())
	app.Post("/doctor/update", controllers.UpdateDoctor())
}
