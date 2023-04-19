package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/gofiber/fiber/v2"
)

func OrderDetailGetController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Get successful"
	return c.Status(fiber.StatusOK).JSON(&r)
}

func OrderDetailPostController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Create successful"
	return c.Status(fiber.StatusOK).JSON(&r)
}

func OrderDetailPutController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Update successful"
	return c.Status(fiber.StatusOK).JSON(&r)
}

func OrderDetailDeleteController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Delete successful"
	return c.Status(fiber.StatusOK).JSON(&r)
}
