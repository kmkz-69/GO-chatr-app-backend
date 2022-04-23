package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kmkz-69/chatr.fr.backend/src/repositories"
)

// func create admin username and password and password confirmation
// if username and password are empty or if password and password confirmation are not the same
// return 400
// if error occurs return 500
// if everything is ok return 201*

func CreateAdmin(c *fiber.Ctx) error {
	var admin repositories.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	
	if admin.Username == "" || admin.Password == ""  {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	
	
}

