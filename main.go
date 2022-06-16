package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nr-wazid/fiber-api/database"
	"github.com/nr-wazid/fiber-api/routes"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

func welcome(c *fiber.Ctx) error {
	return c.SendString(`Welcome To Fiber Api`)
}

func setupRoutes(app *fiber.App) {
	app.Get("/main", welcome)

	// USER
	app.Post("/api/user", routes.CerateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/user/:id", routes.GetUser)
	app.Put("/api/user/:id", routes.UpdateUser)
	app.Delete("/api/user/:id", routes.Delete)
}
