package main

import (
	"log"

	"github.com/Yassinproweb/doc_apoi/data"
	"github.com/Yassinproweb/doc_apoi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	data.InitDB()
	defer data.DB.Close()

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/static", "./static")

	// Homepage
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	// Routes
	routes.DocRoutes(app)
	routes.PatRoutes(app)

	log.Fatal(app.Listen("0.0.0.0:4321"))
}

