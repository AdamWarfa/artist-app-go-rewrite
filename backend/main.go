package main

import (
	"artist-app-backend/internals/controllers"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/artists", controllers.GetArtists)
	app.Get("/artists/:id", controllers.GetSingleArtist)
	app.Post("/artists", controllers.PostArtist)
	app.Put("/artists/:id", controllers.UpdateArtist)
	app.Delete("/artists/:id", controllers.DeleteArtist)

	app.Listen(":3000")
}
