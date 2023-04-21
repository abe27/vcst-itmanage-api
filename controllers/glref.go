package controllers

import (
	"fmt"
	"time"

	"github.com/abe27/vcst/api.v1/configs"
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

	if c.Query("fcrftype") != "" && c.Query("fddate") != "" {
		var gl []models.Glref
		if err := db.Scopes(services.Paginate(c)).
			Order("FCCODE").
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
			Where("FDDATE", c.Query("fddate")).
			Find(&gl, &models.Glref{FCRFTYPE: c.Query("fcrftype")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		r.Data = &gl
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("fcrftype") != "" {
		var gl []models.Glref
		if err := db.Scopes(services.Paginate(c)).
			Order("FCCODE").
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
			Find(&gl, &models.Glref{FCRFTYPE: c.Query("fcrftype")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
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

	// Begin transaction
	tx := db.Begin()
	var whs models.Whouse
	if err := tx.Select("FCSKID").First(&whs, &models.Whouse{FCCODE: frm.FCWHOUSE}).Error; err != nil {
		r.Message = fmt.Sprintf("%s %s", frm.FCWHOUSE, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var branch models.Branch
	if err := tx.Select("FCSKID").First(&branch, &models.Branch{FCCODE: frm.FCBRANCH}).Error; err != nil {
		r.Message = fmt.Sprintf("%s %s", frm.FCBRANCH, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var rnn int64
	if err := tx.Select("FCCODE").Where("FCREFTYPE", frm.FCREFTYPE).Model(&models.Glref{}).Count(&rnn).Error; err != nil {
		panic(err)
	}

	frm.FCCODE = fmt.Sprintf("%06d", rnn+1)
	prefix := frm.FCREFNO
	frm.FCREFNO = fmt.Sprintf("%s%s%06d", prefix, (time.Now().Format("20060102"))[3:], rnn+1)
	frm.FCGID, _ = g.New(26)

	var fcamt float64 = 0
	for _, i := range frm.REFPROD {
		fcamt += i.QTY
	}

	var glref models.Glref
	glref.FCDATASER = configs.FCDATASER
	glref.FCCODE = frm.FCCODE
	glref.FCGID = frm.FCGID
	glref.FCREFNO = frm.FCREFNO
	glref.FCREFTYPE = frm.FCREFTYPE
	glref.FCRFTYPE = frm.FCRFTYPE
	glref.FCSTEP = frm.FCSTEP
	glref.FDDATE = frm.FDDATE
	glref.FCBRANCH = branch.FCSKID
	glref.FNAMT = fcamt
	glref.FIMILLISEC = time.Now().Unix()
	glref.FTDATETIME = time.Now()
	glref.FTLASTEDIT = time.Now()
	glref.FTLASTUPD = time.Now()

	if err := tx.Create(&glref).Error; err != nil {
		// rollback the transaction in case of error
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var seq int64 = 1
	for _, i := range frm.REFPROD {
		var v models.Prod
		if err := tx.Find(&v, &models.Prod{FCCODE: i.FCCODE}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		if v.FCSKID != "" {
			var refProd models.Refprod
			refProd.FCSEQ = fmt.Sprintf("%03d", seq)
			refProd.FCDATASER = configs.FCDATASER
			refProd.FCGLREF = glref.FCSKID
			refProd.FDDATE = glref.FDDATE
			refProd.FCREFPDTYP = "P"
			refProd.FCIOTYPE = glref.FCSTEP
			refProd.FCRFTYPE = glref.FCRFTYPE
			refProd.FCREFTYPE = glref.FCREFTYPE
			refProd.FCPRODTYPE = v.FCPRTYPE
			refProd.FCCORP = v.FCCORP
			refProd.FCBRANCH = branch.FCSKID
			refProd.FCWHOUSE = whs.FCSKID
			refProd.FCSECT = v.FCSECT
			refProd.FCPROD = v.FCSKID
			refProd.FCUM = v.FCUM
			refProd.FCUMSTD = v.FCUM
			refProd.FNUMQTY = float64(i.PACK)
			refProd.FNQTY = i.QTY
			refProd.FCSTUM = v.FCUM
			refProd.FCSTUMSTD = v.FCUM
			refProd.FNSTUMQTY = float64(i.PACK)
			refProd.FTDATETIME = time.Now()
			refProd.FTLASTEDIT = time.Now()
			refProd.FTLASTUPD = time.Now()

			if err := tx.Create(&refProd).Error; err != nil {
				tx.Rollback()
				r.Message = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}

			var stock models.Stock
			tx.First(&stock, &models.Stock{FCPROD: v.FCSKID, FCWHOUSE: whs.FCSKID})
			stock.FCDATASER = configs.FCDATASER
			stock.FCCORP = v.FCCORP
			stock.FCBRANCH = branch.FCSKID
			stock.FCWHOUSE = whs.FCSKID
			stock.FCPROD = v.FCSKID
			stock.FDDATE = glref.FDDATE
			stock.FNQTY = stock.FNQTY + i.QTY
			if stock.FCSKID == "" {
				stock.FTDATETIME = time.Now()
			}
			stock.FTLASTUPD = time.Now()
			stock.FTLASTEDIT = time.Now()
			if err := tx.Save(&stock).Error; err != nil {
				tx.Rollback()
				r.Message = fmt.Sprintf("Failed transection on Stock: %v", err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}
			seq++
		}
	}

	// Commit the transaction in case success
	tx.Commit()
	// End
	msg := fmt.Sprintf("\nเปิดรับสินค้าแบบ Manual\nเลขที่: %s \nสินค้า: %d รายการ\nจำนวน: %d\nเรียบร้อยแล้ว\n%s", glref.FCREFNO, len(frm.REFPROD), int(fcamt), time.Now().Format("2006-01-02 15:04:05"))
	go services.LineNotify(msg)
	r.Data = &glref
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
