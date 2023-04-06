package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func WhsGetController(c *fiber.Ctx) error {
	var r models.Response
	db := configs.Store
	if c.Query("id") != "" {
		var Whs models.Whs
		if err := db.Scopes(services.Paginate(c)).First(&Whs, &models.Whs{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &Whs
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var Whs []models.Whs
	if err := db.Scopes(services.Paginate(c)).First(&Whs, &models.Whs{ID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &Whs
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WhsPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Whs
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := configs.Store
	var Whs models.Whs
	Whs.Name = frm.Name
	Whs.Description = frm.Description
	Whs.IpAddress = frm.IpAddress
	Whs.OnPort = frm.OnPort
	Whs.User = frm.User
	Whs.Passord = frm.Passord
	Whs.DBName = frm.DBName
	if err := db.Create(&Whs).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Create successfully"
	r.Data = &Whs
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func WhsPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Whs
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var Whs models.Whs
	db := configs.Store
	if err := db.First(&Whs, &models.Whs{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	Whs.Name = frm.Name
	Whs.Description = frm.Description
	Whs.IpAddress = frm.IpAddress
	Whs.OnPort = frm.OnPort
	Whs.User = frm.User
	Whs.Passord = frm.Passord
	Whs.DBName = frm.DBName
	if err := db.Save(&Whs).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Update successfully"
	r.Data = &Whs
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WhsDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var Whs models.Whs
	db := configs.Store
	if err := db.First(&Whs, &models.Whs{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	if err := db.Delete(&Whs).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
