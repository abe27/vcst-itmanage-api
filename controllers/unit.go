package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func UnitGetController(c *fiber.Ctx) error {
	var r models.Response
	db := WHSDb(c)
	if c.Query("id") != "" {
		var unit models.Unit
		if err := db.Scopes(services.Paginate(c)).First(&unit, &models.Unit{FCSKID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &unit
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var unit []models.Unit
	if err := db.Scopes(services.Paginate(c)).First(&unit, &models.Unit{FCSKID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &unit
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UnitPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Unit
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	// db := WHSDb(c)
	var Unit models.Unit
	Unit.FCCODE = frm.FCCODE
	Unit.FCNAME = frm.FCNAME
	Unit.FCNAME2 = frm.FCNAME2
	// if err := db.Create(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Create successfully"
	r.Data = &Unit
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func UnitPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Unit
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var Unit models.Unit
	db := WHSDb(c)
	if err := db.First(&Unit, &models.Unit{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	Unit.FCCODE = frm.FCCODE
	Unit.FCNAME = frm.FCNAME
	Unit.FCNAME2 = frm.FCNAME2
	// if err := db.Save(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Update successfully"
	r.Data = &Unit
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UnitDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var unit models.Unit
	db := WHSDb(c)
	if err := db.First(&unit, &models.Unit{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	// if err := db.Delete(&Unit).Error; err!= nil {
	//   r.Message = err.Error()
	//   return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
