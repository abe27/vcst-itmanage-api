package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func CompanyGetController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.Store
	if c.Query("id") != "" {
		var Company models.Company
		if err := db.Scopes(services.Paginate(c)).First(&Company, &models.Company{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &Company
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var Company []models.Company
	if err := db.Scopes(services.Paginate(c)).First(&Company, &models.Company{ID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &Company
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CompanyPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Company
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.Store
	var Company models.Company
	Company.Name = frm.Name
	Company.Description = frm.Description
	Company.IsActive = frm.IsActive
	if err := db.Create(&Company).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Create successfully"
	r.Data = &Company
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func CompanyPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Company
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var Company models.Company
	db := configs.Store
	if err := db.First(&Company, &models.Company{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	Company.Name = frm.Name
	Company.Description = frm.Description
	Company.IsActive = frm.IsActive
	if err := db.Save(&Company).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Update successfully"
	r.Data = &Company
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CompanyDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var Company models.Company
	db := configs.Store
	if err := db.First(&Company, &models.Company{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	if err := db.Delete(&Company).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
