package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func ProductTypeGetController(c *fiber.Ctx) error {
	var r models.Response
	db := WHSDb(c)
	if c.Query("id") != "" {
		var productType models.ProductType
		if err := db.Scopes(services.Paginate(c)).First(&productType, &models.ProductType{FCSKID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &productType
		return c.Status(fiber.StatusOK).JSON(r)
	}

	var productType []models.ProductType
	if err := db.Scopes(services.Paginate(c)).Find(&productType).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &productType
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ProductTypePostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.ProductType
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	// db := WHSDb(c)
	var productType models.ProductType
	productType.FCCODE = frm.FCCODE
	productType.FCNAME = frm.FCNAME
	productType.FCNAME2 = frm.FCNAME2
	// if err := db.Create(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Create successfully"
	r.Data = &productType
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ProductTypePutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.ProductType
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var productType models.ProductType
	db := WHSDb(c)
	if err := db.First(&productType, &models.ProductType{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	productType.FCCODE = frm.FCCODE
	productType.FCNAME = frm.FCNAME
	productType.FCNAME2 = frm.FCNAME2
	// if err := db.Save(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Update successfully"
	r.Data = &productType
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ProductTypeDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var productType models.ProductType
	db := WHSDb(c)
	if err := db.First(&productType, &models.ProductType{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	// if err := db.Delete(&productType).Error; err!= nil {
	//   r.Message = err.Error()
	//   return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
