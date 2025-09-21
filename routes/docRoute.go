package routes

import (
	"github.com/Yassinproweb/TeleMedi/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DocRoutes(app *fiber.App) {
	s := session.New()
	controllers.Store = s

	app.Get("/", controllers.GetDoctorsController())
	app.Post("/doctors", controllers.RegisterDoctorController(s))
	app.Post("/doctors", controllers.LoginDoctorController(s))
	// app.Post("/patients", controllers.RegisterPatient(s))
	// app.Post("/patients", controllers.LoginPatient(s))
	app.Get("/doctor/edit", controllers.EditDoctorFormController())
	app.Post("/doctor/update", controllers.UpdateDoctorController())

	// Authentication middleware
	isAuthenticated := func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		doctorID := sess.Get("doctor_id")
		if doctorID == nil {
			return c.Redirect("/login_doctor")
		}
		return c.Next()
	}

	// Protected dashboard route
	app.Get("/dashboard", isAuthenticated, func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Session error")
		}
		name := sess.Get("doctor_name").(string)
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
