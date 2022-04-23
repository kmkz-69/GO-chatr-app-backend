package middlewares

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/kmkz-69/chatr.fr.backend/src/repositories"

)

// AuthMiddleware is the middleware for authentication
func AuthMiddleware(c *fiber.Ctx) error {
	accessToken := c.Get("access_token")
	if accessToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}



	c.Locals("admin")
	return nil
}