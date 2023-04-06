package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func DocumentGetController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.Store
	if c.Query("id") != "" {
		var Document models.BillingDocument
		if err := db.Scopes(services.Paginate(c)).First(&Document, &models.BillingDocument{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &Document
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var Document []models.BillingDocument
	if err := db.Scopes(services.Paginate(c)).First(&Document, &models.BillingDocument{ID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &Document
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DocumentPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.BillingDocument
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.Store
	var Document models.BillingDocument
	Document.Title = frm.Title
	Document.Description = frm.Description
	Document.IsActive = frm.IsActive
	if err := db.Create(&Document).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Create successfully"
	r.Data = &Document
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func DocumentPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.BillingDocument
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var Document models.BillingDocument
	db := configs.Store
	if err := db.First(&Document, &models.BillingDocument{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	Document.Title = frm.Title
	Document.Description = frm.Description
	Document.IsActive = frm.IsActive
	if err := db.Save(&Document).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Update successfully"
	r.Data = &Document
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DocumentDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var Document models.BillingDocument
	db := configs.Store
	if err := db.First(&Document, &models.BillingDocument{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	if err := db.Delete(&Document).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
