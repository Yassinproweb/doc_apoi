package main

import (
	"log"
	"time"

	"github.com/Yassinproweb/meet_doc.git/data"
	"github.com/Yassinproweb/meet_doc.git/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize database
	data.InitDB()
	defer data.DB.Close()

	// Fiber app with template engine
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	// Serve static files
	app.Static("/static", "./static")

	// Single session store for the whole app
	s := session.New(session.Config{
		Expiration:     24 * time.Hour,
		CookieHTTPOnly: true,
		CookieSecure:   false, // set to true when using HTTPS in production
	})

	// Routes (pass session store)
	routes.DocRoutes(app, s)

	// Start server
	log.Fatal(app.Listen("0.0.0.0:4320"))
}
