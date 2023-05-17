package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func FormulaSectionGetController(c *fiber.Ctx) error {
	var r models.Response
	db := WHSDb(c)
	var sectData []models.Sect
	if err := db.Scopes(services.Paginate(c)).Order("FCCODE").Where("FCNAME NOT LIKE '%ยกเลิก%'").Where("FCSTATUS NOT IN (?)", "I").Find(&sectData).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Data = &sectData
	return c.Status(fiber.StatusOK).JSON(&r)
}
