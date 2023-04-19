package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func OrderHeadGetController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = "ok"
	if c.Query("id") != "" {
		var order models.Orderh
		if err := db.
			Preload("Corp").
			Preload("Book").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Coor").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			Preload("Proj").
			Preload("DeliverTo").
			Preload("Payterm").
			First(&order, &models.Orderh{FCSKID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Data = &order
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("book") != "" {
		var order models.Orderh
		if err := db.
			Preload("Corp").
			Preload("Book").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Coor").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			Preload("Proj").
			Preload("DeliverTo").
			Preload("Payterm").
			First(&order, &models.Orderh{FCBOOK: c.Query("book")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Data = &order
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var order []models.Orderh
	if err := db.Scopes(services.Paginate(c)).
		Preload("Corp").
		Preload("Book").
		Preload("Branch").
		Preload("Dept").
		Preload("Sect").
		Preload("Job").
		Preload("Coor").
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("Proj").
		Preload("DeliverTo").
		Preload("Payterm").
		Find(&order).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Data = &order
	return c.Status(fiber.StatusOK).JSON(&r)
}

func OrderHeadPostController(c *fiber.Ctx) error {
	// db := WHSDb(c)
	var r models.Response
	r.Message = "Create Successful"
	var frm models.Orderh
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	// var order models.Orderh

	return c.Status(fiber.StatusOK).JSON(&r)
}

func OrderHeadPutController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func OrderHeadDeleteController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
