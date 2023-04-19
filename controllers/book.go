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
	// db := WHSDb(c)
	var r models.Response
	r.Message = "Create Successfully!"
	var frm models.Book
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var b models.Book
	b.FCACCBOOK = frm.FCACCBOOK
	b.FCATSTEP = frm.FCATSTEP
	b.FCBRANCH = frm.FCBRANCH
	b.FCCODE = frm.FCCODE
	b.FCCORP = frm.FCCORP
	b.FCCORRECTB = frm.FCCORRECTB
	b.FCCREATEAP = frm.FCCREATEAP
	b.FCCREATEBY = frm.FCCREATEBY
	b.FCEAFTERR = frm.FCEAFTERR
	b.FCFCHR = frm.FCFCHR
	b.FCGMODHEAD = frm.FCGMODHEAD
	b.FCNAME = frm.FCNAME
	b.FCNAME2 = frm.FCNAME2
	b.FCREFTYPE = frm.FCREFTYPE
	b.FCTYPELINK = frm.FCTYPELINK
	b.FCWHOUSE = frm.FCWHOUSE
	b.FNVATRATE = frm.FNVATRATE

	r.Data = &b
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func BookPutController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = "Update Successfully!"
	var frm models.Book
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var b models.Book
	if err := db.First(&b, &models.Book{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	b.FCACCBOOK = frm.FCACCBOOK
	b.FCATSTEP = frm.FCATSTEP
	b.FCBRANCH = frm.FCBRANCH
	b.FCCODE = frm.FCCODE
	b.FCCORP = frm.FCCORP
	b.FCCORRECTB = frm.FCCORRECTB
	b.FCCREATEAP = frm.FCCREATEAP
	b.FCCREATEBY = frm.FCCREATEBY
	b.FCEAFTERR = frm.FCEAFTERR
	b.FCFCHR = frm.FCFCHR
	b.FCGMODHEAD = frm.FCGMODHEAD
	b.FCNAME = frm.FCNAME
	b.FCNAME2 = frm.FCNAME2
	b.FCREFTYPE = frm.FCREFTYPE
	b.FCTYPELINK = frm.FCTYPELINK
	b.FCWHOUSE = frm.FCWHOUSE
	b.FNVATRATE = frm.FNVATRATE

	r.Data = &b
	return c.Status(fiber.StatusOK).JSON(&r)
}

func BookDeleteController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = "Delete Successfully!"
	var b models.Book
	if err := db.First(&b, &models.Book{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := db.Delete(&b).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	return c.Status(fiber.StatusOK).JSON(&r)
}
