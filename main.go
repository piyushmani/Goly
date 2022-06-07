package main

import (
	"goly/handler"
	"goly/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Post("/goly", handler.CreateGoly)
	app.Get("/r/:redirect", handler.Redirect)
}

func main() {
	model.Setup()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	app.Listen(":3000")
}
