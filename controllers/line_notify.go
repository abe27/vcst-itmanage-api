package controllers

import (
	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/gofiber/fiber/v2"
)

func LineNotifyGetController(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var lineNotify models.Linenotify
		if err := configs.Store.First(&lineNotify, &models.Linenotify{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Data = &lineNotify
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var lineNotify []models.Linenotify
	if err := configs.Store.Find(&lineNotify).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Data = &lineNotify
	return c.Status(fiber.StatusOK).JSON(&r)
}

func LineNotifyPostController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Linenotify
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var line models.Linenotify
	line.Jobs = frm.Jobs
	line.Message = frm.Message
	line.Token = frm.Token
	line.IsActive = frm.IsActive
	if err := configs.Store.Create(&line).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Post"
	r.Data = &line
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func LineNotifyPutController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Linenotify
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var line models.Linenotify
	if err := configs.Store.First(&line, &models.Linenotify{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	line.Jobs = frm.Jobs
	line.Message = frm.Message
	line.Token = frm.Token
	line.IsActive = frm.IsActive
	if err := configs.Store.Create(&line).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Put"
	r.Data = &line
	return c.Status(fiber.StatusOK).JSON(&r)
}

func LineNotifyDeleteController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Linenotify
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var line models.Linenotify
	if err := configs.Store.First(&line, &models.Linenotify{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Delete(&line).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Delete"
	r.Data = &line
	return c.Status(fiber.StatusOK).JSON(&r)
}
