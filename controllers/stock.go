package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func StockGetController(c *fiber.Ctx) error {
	var r models.Response
	db := WHSDb(c)
	if c.Query("prod_id") != "" {
		var Stock models.Stock
		if err := db.Preload("Product.ProductType").Preload("Product.Unit").Preload("WHouse").First(&Stock, &models.Stock{FCPROD: c.Query("prod_id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(r)
		}

		r.Message = "ok"
		r.Data = &Stock
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("id") != "" {
		var Stock models.Stock
		if err := db.Preload("Product.ProductType").Preload("Product.Unit").Preload("WHouse").First(&Stock, &models.Stock{FCSKID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(r)
		}

		r.Message = "ok"
		r.Data = &Stock
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var Stock []models.Stock
	if err := db.Scopes(services.Paginate(c)).Preload("Product.ProductType").Preload("Product.Unit").Preload("WHouse").Find(&Stock, &models.Stock{FCSKID: c.Query("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "ok"
	r.Data = &Stock
	return c.Status(fiber.StatusOK).JSON(&r)
}

func StockPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Stock
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	// db := WHSDb(c)
	var Stock models.Stock
	Stock.FCCORP = frm.FCCORP
	Stock.FCPROD = frm.FCPROD
	Stock.FCWHOUSE = frm.FCWHOUSE
	Stock.FNQTY = frm.FNQTY
	Stock.FNPRICE = frm.FNPRICE
	Stock.FNCOST = frm.FNCOST
	// if err := db.Create(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Create successfully"
	r.Data = &Stock
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func StockPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Stock
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	var Stock models.Stock
	db := WHSDb(c)
	if err := db.First(&Stock, &models.Stock{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}
	Stock.FCCORP = frm.FCCORP
	Stock.FCPROD = frm.FCPROD
	Stock.FCWHOUSE = frm.FCWHOUSE
	Stock.FNQTY = frm.FNQTY
	Stock.FNPRICE = frm.FNPRICE
	Stock.FNCOST = frm.FNCOST
	// if err := db.Save(&prod).Error; err != nil {
	// 	r.Message = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }

	r.Message = "Update successfully"
	r.Data = &Stock
	return c.Status(fiber.StatusOK).JSON(&r)
}

func StockDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var Stock models.Stock
	db := WHSDb(c)
	if err := db.First(&Stock, &models.Stock{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	// if err := db.Delete(&Stock).Error; err!= nil {
	//   r.Message = err.Error()
	//   return c.Status(fiber.StatusInternalServerError).JSON(r)
	// }
	r.Message = "Delete successfully"
	return c.Status(fiber.StatusOK).JSON(&r)
}
