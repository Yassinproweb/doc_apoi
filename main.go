package main

import (
	"log"
	"time"

	"github.com/Yassinproweb/TeleMedi/controllers"
	"github.com/Yassinproweb/TeleMedi/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	data.InitDB()
	defer data.DB.Close()

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/", "./static")

	s := session.New(session.Config{
		Expiration:     24 * time.Hour,
		CookieHTTPOnly: true,
		CookieSecure:   false, // Set to true in production with HTTPS
	})

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

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/view", func(c *fiber.Ctx) error {
		return c.Render("dashboard", nil)
	})

	// doctor registration
	app.Post("/register_doctor", controllers.RegisterDoctor(s))

	// patient registration
	app.Post("/register_patient", controllers.RegisterPatient(s))

	// doctor login
	app.Post("/login_doctor", controllers.LoginDoctor(s))

	// patient login
	app.Post("/login_patient", controllers.LoginPatient(s))

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

	// Start server
	log.Fatal(app.Listen(":4800"))
}
