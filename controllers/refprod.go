package controllers

import (
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func RefProdGetController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "OK"
	db := WHSDb(c)
	if c.Query("id") != "" {
		var refProd models.Refprod
		if err := db.
			Order("FCSEQ").
			Preload("Corp").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Glhead").
			Preload("Glref.Book").
			Preload("Glref.FromWhouse").
			Preload("Glref.ToWhouse").
			Preload("Prod").
			Preload("Unit").
			Preload("UnitSTD").
			Preload("Stum").
			Preload("StumStd").
			Preload("WHouse").
			Preload("Proj").
			Preload("Gl").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			First(&refProd, &models.Refprod{}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Data = &refProd
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("glref_id") != "" {
		var refProd []models.Refprod
		if err := db.
			Order("FCSEQ").
			Preload("Corp").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Glhead").
			Preload("Glref.Book").
			Preload("Glref.FromWhouse").
			Preload("Glref.ToWhouse").
			Preload("Prod").
			Preload("Unit").
			Preload("UnitSTD").
			Preload("Stum").
			Preload("StumStd").
			Preload("WHouse").
			Preload("Proj").
			Preload("Gl").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			Find(&refProd, &models.Refprod{FCGLREF: c.Query("glref_id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Data = &refProd
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var refProd []models.Refprod
	if err := db.Scopes(services.Paginate(c)).Find(&refProd).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Data = &refProd
	return c.Status(fiber.StatusOK).JSON(&r)
}
