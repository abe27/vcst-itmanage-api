package controllers

import (
	"fmt"

	"github.com/abe27/vcst/api.v1/configs"
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

		var isCompleted int64
		if err := configs.Store.Select("id").Find(&models.GlrefHistory{}, &models.GlrefHistory{GLREF: refProd.Glref.FCSKID}).Count(&isCompleted).Error; err != nil {
			refProd.Glref.FCSTATUS = false
		}
		refProd.Glref.FCSTATUS = isCompleted > 0
		r.Data = &refProd
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var refProdTable []models.Refprod
	var refProd []models.Refprod
	if c.Query("glref_id") != "" {
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

		for _, x := range refProd {
			var isCompleted int64
			if err := configs.Store.Select("id").Find(&models.GlrefHistory{}, &models.GlrefHistory{GLREF: x.Glref.FCSKID}).Count(&isCompleted).Error; err != nil {
				x.Glref.FCSTATUS = false
			}
			x.Glref.FCSTATUS = isCompleted > 0
			refProdTable = append(refProdTable, x)
		}

		r.Data = &refProdTable
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if err := db.Scopes(services.Paginate(c)).Find(&refProd).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	for _, x := range refProd {
		var isCompleted int64
		if err := configs.Store.Select("id").Find(&models.GlrefHistory{}, &models.GlrefHistory{GLREF: x.Glref.FCSKID}).Count(&isCompleted).Error; err != nil {
			x.Glref.FCSTATUS = false
		}
		x.Glref.FCSTATUS = isCompleted > 0
		refProdTable = append(refProdTable, x)
	}

	r.Data = &refProdTable
	return c.Status(fiber.StatusOK).JSON(&r)
}

func RefProdDeleteController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = fmt.Sprintf("Delete %s.", c.Params("id"))
	var refProd models.Refprod
	if err := db.First(&refProd, &models.Refprod{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	glRefID := refProd.FCGLREF

	if err := db.Delete(&refProd).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var listProd []models.Refprod
	if err := db.Find(&listProd, &models.Refprod{FCGLREF: glRefID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if len(listProd) <= 0 {
		var glref models.Glref
		if err := db.First(&glref, &models.Glref{FCSKID: refProd.FCGLREF}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		// Delete GLREF
		if err := db.Delete(&glref, &models.Glref{FCSKID: refProd.FCGLREF}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
	}
	return c.Status(fiber.StatusOK).JSON(&r)
}
