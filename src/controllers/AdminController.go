package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kmkz-69/chatr.fr.backend/src/repositories"
)

type AdminController struct {
	fiber.Error
	fiber.Request
	fiber.Response
	adminRepository *repositories.AdminRepository
}

func create(){

}


