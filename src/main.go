package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kmkz-69/chatr.fr.backend/src/database"
)


func main() {
	database.ConnectToDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
