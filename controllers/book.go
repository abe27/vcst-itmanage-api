package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/gofiber/fiber/v2"
)

func BookGetController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func BookPostController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func BookPutController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func BookDeleteController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
