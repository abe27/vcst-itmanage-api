package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func WHouseGetController(c *fiber.Ctx) error {
	var r models.Response
	db := WHSDb(c)
	if c.Query("id") != "" {
		var WHouse models.WHouse
		if err := db.Scopes(services.Paginate(c)).First(&WHouse, &models.WHouse{FCSKID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &WHouse
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var WHouse []models.WHouse
	if err := db.Scopes(services.Paginate(c)).First(&WHouse, &models.WHouse{FCSKID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &WHouse
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WHousePostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.WHouse
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	// db := WHSDb(c)
	var WHouse models.WHouse
	WHouse.FCCODE = frm.FCCODE
	WHouse.FCNAME = frm.FCNAME
	WHouse.FCNAME2 = frm.FCNAME2
	// if err := db.Create(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Create successfully"
	r.Data = &WHouse
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func WHousePutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.WHouse
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var WHouse models.WHouse
	db := WHSDb(c)
	if err := db.First(&WHouse, &models.WHouse{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	WHouse.FCCODE = frm.FCCODE
	WHouse.FCNAME = frm.FCNAME
	WHouse.FCNAME2 = frm.FCNAME2
	// if err := db.Save(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Update successfully"
	r.Data = &WHouse
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WHouseDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var WHouse models.WHouse
	db := WHSDb(c)
	if err := db.First(&WHouse, &models.WHouse{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	// if err := db.Delete(&WHouse).Error; err!= nil {
	//   r.Message = err.Error()
	//   return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
