package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func DepartmentGetController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.Store
	if c.Query("id") != "" {
		var Department models.Department
		if err := db.Scopes(services.Paginate(c)).First(&Department, &models.Department{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &Department
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var Department []models.Department
	if err := db.Scopes(services.Paginate(c)).First(&Department, &models.Department{ID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &Department
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DepartmentPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Department
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.Store
	var Department models.Department
	Department.Title = frm.Title
	Department.Description = frm.Description
	Department.IsActive = frm.IsActive
	if err := db.Create(&Department).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Create successfully"
	r.Data = &Department
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func DepartmentPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Department
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var Department models.Department
	db := configs.Store
	if err := db.First(&Department, &models.Department{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	Department.Title = frm.Title
	Department.Description = frm.Description
	Department.IsActive = frm.IsActive
	if err := db.Save(&Department).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Update successfully"
	r.Data = &Department
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DepartmentDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var Department models.Department
	db := configs.Store
	if err := db.First(&Department, &models.Department{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	if err := db.Delete(&Department).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
