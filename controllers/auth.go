package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusCreated).JSON(&r)
}
