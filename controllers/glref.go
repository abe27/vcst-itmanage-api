package controllers

import (
	"fmt"

	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
	g "github.com/matoous/go-nanoid/v2"
)

func GlrefHeaderGetController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = "Get"
	if c.Query("id") != "" {
		var gl models.Glref
		if err := db.
			Preload("Corp").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Glhead").
			Preload("Book").
			Preload("Coor").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			Preload("VatCoor").
			Preload("Proj").
			Preload("DeliveryToCoor").
			First(&gl, &models.Glref{}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Data = &gl
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var gl []models.Glref
	if err := db.Scopes(services.Paginate(c)).
		Preload("Corp").
		Preload("Branch").
		Preload("Dept").
		Preload("Sect").
		Preload("Job").
		Preload("Glhead").
		Preload("Book").
		Preload("Coor").
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("VatCoor").
		Preload("Proj").
		Preload("DeliveryToCoor").
		Find(&gl).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Data = &gl
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlrefHeaderPostController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = "Post"
	var frm models.GlRefForm
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var rnn int64
	if err := db.Select("FCCODE").Where("FCREFTYPE", frm.FCREFTYPE).Model(&models.Glref{}).Count(&rnn).Error; err != nil {
		panic(err)
	}

	frm.FCCODE = fmt.Sprintf("%06d", rnn+1)
	frm.FCGID, _ = g.New(26)
	r.Data = &frm
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func GlrefHeaderPutController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Put"
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlrefHeaderDeleteController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = "Deleted"
	var glref models.Glref
	if err := db.First(&glref, &models.Glref{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := db.Delete(&glref).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	return c.Status(fiber.StatusOK).JSON(&r)
}
