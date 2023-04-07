package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func ProductGetController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	prodType := "4"
	if c.Query("type") != "" {
		prodType = c.Query("type")
	}

	if c.Query("id") != "" {
		var prod models.Product
		if err := db.Scopes(services.Paginate(c)).Preload("ProductType").Preload("Unit").First(&prod, &models.Product{FCSKID: c.Query("id"), FCTYPE: prodType}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(r)
		}

		r.Message = "ok"
		r.Data = &prod
		return c.Status(fiber.StatusOK).JSON(r)
	}

	var prod []models.Product
	if err := db.Scopes(services.Paginate(c)).Preload("ProductType").Preload("Unit").Find(&prod, &models.Product{FCTYPE: prodType}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &prod
	return c.Status(fiber.StatusOK).JSON(r)
}

func ProductPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Product
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	// db := checkWHS(c)
	var prod models.Product
	prod.FCTYPE = frm.FCTYPE
	prod.FCCODE = frm.FCCODE
	prod.FCSNAME = frm.FCSNAME
	prod.FCSNAME2 = frm.FCSNAME2
	prod.FCNAME = frm.FCNAME
	prod.FCNAME2 = frm.FCNAME2
	// prod.FNAVGCOST = frm.FNAVGCOST
	// prod.FNSTDCOST = frm.FNSTDCOST
	// prod.FCSTATUS = frm.FCSTATUS
	// if err := db.Create(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Create successfully"
	r.Data = &prod
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ProductPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Product
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := WHSDb(c)
	var prod models.Product
	if err := db.First(&prod, &models.Product{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	prod.FCTYPE = frm.FCTYPE
	prod.FCCODE = frm.FCCODE
	prod.FCSNAME = frm.FCSNAME
	prod.FCSNAME2 = frm.FCSNAME2
	prod.FCNAME = frm.FCNAME
	prod.FCNAME2 = frm.FCNAME2
	// prod.FNAVGCOST = frm.FNAVGCOST
	// prod.FNSTDCOST = frm.FNSTDCOST
	// prod.FCSTATUS = frm.FCSTATUS
	// if err := db.Save(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Update successfully"
	r.Data = &prod
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ProductDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Product
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	db := WHSDb(c)
	var prod models.Product
	if err := db.First(&prod, &models.Product{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	// if err := db.Delete(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Delete successfully"
	r.Data = &prod
	return c.Status(fiber.StatusOK).JSON(&r)
}
