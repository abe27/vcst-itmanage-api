package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func SectionGetController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.Store
	if c.Query("id") != "" {
		var Section models.Section
		if err := db.Scopes(services.Paginate(c)).First(&Section, &models.Section{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &Section
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var Section []models.Section
	if err := db.Scopes(services.Paginate(c)).First(&Section, &models.Section{ID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &Section
	return c.Status(fiber.StatusOK).JSON(&r)
}

func SectionPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Section
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.Store
	var Section models.Section
	Section.Title = frm.Title
	Section.Description = frm.Description
	Section.IsActive = frm.IsActive
	if err := db.Create(&Section).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Create successfully"
	r.Data = &Section
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func SectionPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Section
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var Section models.Section
	db := configs.Store
	if err := db.First(&Section, &models.Section{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	Section.Title = frm.Title
	Section.Description = frm.Description
	Section.IsActive = frm.IsActive
	if err := db.Save(&Section).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Update successfully"
	r.Data = &Section
	return c.Status(fiber.StatusOK).JSON(&r)
}

func SectionDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var Section models.Section
	db := configs.Store
	if err := db.First(&Section, &models.Section{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	if err := db.Delete(&Section).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
