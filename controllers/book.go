package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func BookGetController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = "Ok"
	if c.Query("id") != "" {
		var book models.Book
		if err := db.First(&book, &models.Book{FCSKID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		r.Data = &book
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	fctype := "PO"
	if c.Query("fctype") != "" {
		fctype = c.Query("fctype")
	}

	var book []models.Book
	if err := db.Scopes(services.Paginate(c)).
		Preload("CreatedBy").
		Preload("Corp").
		Preload("WHouse").
		Preload("Accbook").
		Preload("Gmodhead").
		Preload("Branch").
		Find(&book, &models.Book{FCREFTYPE: fctype}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Data = &book
	return c.Status(fiber.StatusOK).JSON(&r)
}

func BookPostController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func BookPutController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func BookDeleteController(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
