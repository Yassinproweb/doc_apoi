package main

import (
	"log"

	"github.com/Yassinproweb/doc_apoi/data"
	"github.com/Yassinproweb/doc_apoi/routes"
	"github.com/Yassinproweb/doc_apoi/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	data.InitDB()
	defer data.DB.Close()

	engine := html.New("./views", ".html")
	engine.AddFunc("normalize", utils.NormalizeName)
	engine.AddFunc("capitalize", utils.Capitalize)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	// Homepage
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	// Routes
	routes.DocRoutes(app)
	routes.PatRoutes(app)

	log.Fatal(app.Listen("0.0.0.0:4002"))
}
