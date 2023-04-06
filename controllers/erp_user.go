package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func ErpUserGetController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.Store
	if c.Query("id") != "" {
		var erp models.ErpUser
		if err := db.Preload("Department").First(&erp, &models.ErpUser{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &erp
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var erp []models.ErpUser
	if err := db.Scopes(services.Paginate(c)).Preload("Department").First(&erp, &models.ErpUser{ID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &erp
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ErpUserPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.ErpUser
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.Store
	var department models.Department
	if err := db.First(&department, &models.Department{Title: frm.DepartmentID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	var erp models.ErpUser
	erp.UserName = frm.UserName
	erp.Password = frm.Password
	erp.DepartmentID = department.ID
	erp.IsActive = frm.IsActive
	if err := db.Create(&erp).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Create successfully"
	r.Data = &erp
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ErpUserPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.ErpUser
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var erp models.ErpUser
	db := configs.Store
	var department models.Department
	if err := db.First(&department, &models.Department{Title: frm.DepartmentID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	if err := db.First(&erp, &models.ErpUser{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	erp.UserName = frm.UserName
	erp.Password = frm.Password
	erp.DepartmentID = department.ID
	erp.IsActive = frm.IsActive
	if err := db.Save(&erp).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Update successfully"
	r.Data = &erp
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ErpUserDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var erp models.ErpUser
	db := configs.Store
	if err := db.First(&erp, &models.ErpUser{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	if err := db.Delete(&erp).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
