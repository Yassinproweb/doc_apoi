package main

import (
	"log"
	"time"

	"github.com/Yassinproweb/TeleMedi/data"
	"github.com/Yassinproweb/TeleMedi/routes"
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

	// Routes
	routes.DocRoutes(app)

	// Start server
	log.Fatal(app.Listen(":4300"))
}
