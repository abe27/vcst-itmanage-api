package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func PositionGetController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.Store
	if c.Query("id") != "" {
		var Position models.Position
		if err := db.Scopes(services.Paginate(c)).First(&Position, &models.Position{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &Position
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var Position []models.Position
	if err := db.Scopes(services.Paginate(c)).First(&Position, &models.Position{ID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &Position
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PositionPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Position
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.Store
	var Position models.Position
	Position.Title = frm.Title
	Position.Description = frm.Description
	Position.IsActive = frm.IsActive
	if err := db.Create(&Position).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Create successfully"
	r.Data = &Position
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PositionPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Position
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var Position models.Position
	db := configs.Store
	if err := db.First(&Position, &models.Position{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	Position.Title = frm.Title
	Position.Description = frm.Description
	Position.IsActive = frm.IsActive
	if err := db.Save(&Position).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Update successfully"
	r.Data = &Position
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PositionDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var Position models.Position
	db := configs.Store
	if err := db.First(&Position, &models.Position{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	if err := db.Delete(&Position).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
