package controllers

import (
	"fmt"
	"strings"
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
		var gl models.GlrefTable
		if err := db.
			Preload("Corp").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Glhead").
			Preload("Book").
			Preload("Coor").
			Preload("FromWhouse").
			Preload("ToWhouse").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			Preload("VatCoor").
			Preload("Proj").
			Preload("DeliveryToCoor").
			First(&gl, &models.Glref{}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		var isComplete int64
		if err := configs.Store.Select("id").First(&models.GlrefHistory{}, &models.GlrefHistory{FCSKID: gl.FCSKID}).Count(&isComplete).Error; err != nil {
			gl.FCSTATUS = false
		}
		gl.FCSTATUS = isComplete > 0
		r.Data = &gl
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var glChecked []models.GlrefTable
	var gl []models.GlrefTable
	if c.Query("filterGlrefNo") != "" && c.Query("fddate") != "" {
		if err := db.Scopes(services.Paginate(c)).
			Order("FCLUPDAPP DESC,FCCODE").
			Preload("Corp").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Glhead").
			Preload("Book").
			Preload("Coor").
			Preload("FromWhouse").
			Preload("ToWhouse").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			Preload("VatCoor").
			Preload("Proj").
			Preload("DeliveryToCoor").
			Where("FNAMT > ?", 0).
			Where("FDDATE", c.Query("fddate")).
			Where("FCREFNO LIKE ?", "%"+strings.ToUpper(c.Query("filterGlrefNo"))+"%").
			Find(&gl).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		for _, x := range gl {
			var isComplete int64
			if err := configs.Store.Select("id").First(&models.GlrefHistory{}, &models.GlrefHistory{FCSKID: x.FCSKID}).Count(&isComplete).Error; err != nil {
				x.FCSTATUS = false
			}
			x.FCSTATUS = isComplete > 0
			glChecked = append(glChecked, x)
		}
		r.Data = &glChecked
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("fcrftype") != "" {
		filterDate := time.Now().Format("2006-01-02")
		if c.Query("fddate") != "" {
			filterDate = c.Query("fddate")
		}

		if err := db.Scopes(services.Paginate(c)).
			Order("FCLUPDAPP DESC,FCCODE").
			Preload("Corp").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Glhead").
			Preload("Book").
			Preload("Coor").
			Preload("FromWhouse").
			Preload("ToWhouse").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			Preload("VatCoor").
			Preload("Proj").
			Preload("DeliveryToCoor").
			Where("FNAMT > ?", 0).
			Where("FDDATE", filterDate).
			Find(&gl, &models.Glref{FCRFTYPE: c.Query("fcrftype")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		for _, x := range gl {
			var isComplete int64
			if err := configs.Store.Select("id").First(&models.GlrefHistory{}, &models.GlrefHistory{FCSKID: x.FCSKID}).Count(&isComplete).Error; err != nil {
				x.FCSTATUS = false
			}
			x.FCSTATUS = isComplete > 0
			glChecked = append(glChecked, x)
		}
		r.Data = &glChecked
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("fddate") != "" {
		filterDate := time.Now().Format("2006-01-02")
		if c.Query("fddate") != "" {
			filterDate = c.Query("fddate")
		}
		if err := db.Scopes(services.Paginate(c)).
			Order("FCLUPDAPP DESC,FCCODE").
			Preload("Corp").
			Preload("Branch").
			Preload("Dept").
			Preload("Sect").
			Preload("Job").
			Preload("Glhead").
			Preload("Book").
			Preload("Coor").
			Preload("FromWhouse").
			Preload("ToWhouse").
			Preload("CreatedBy").
			Preload("UpdatedBy").
			Preload("VatCoor").
			Preload("Proj").
			Preload("DeliveryToCoor").
			Where("FNAMT > ?", 0).
			Where("FDDATE", filterDate).
			Find(&gl).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		for _, x := range gl {
			var isComplete int64
			if err := configs.Store.Select("id").First(&models.GlrefHistory{}, &models.GlrefHistory{FCSKID: x.FCSKID}).Count(&isComplete).Error; err != nil {
				x.FCSTATUS = false
			}
			x.FCSTATUS = isComplete > 0
			glChecked = append(glChecked, x)
		}
		r.Data = &glChecked
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if err := db.Scopes(services.Paginate(c)).
		Order("FCLUPDAPP DESC,FCCODE").
		Preload("Corp").
		Preload("Branch").
		Preload("Dept").
		Preload("Sect").
		Preload("Job").
		Preload("Glhead").
		Preload("Book").
		Preload("Coor").
		Preload("FromWhouse").
		Preload("ToWhouse").
		Preload("CreatedBy").
		Preload("UpdatedBy").
		Preload("VatCoor").
		Preload("Proj").
		Preload("DeliveryToCoor").
		Where("FNAMT > ?", 0).
		Find(&gl).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	for _, x := range gl {
		var isComplete int64
		if err := configs.Store.Select("id").First(&models.GlrefHistory{}, &models.GlrefHistory{FCSKID: x.FCSKID}).Count(&isComplete).Error; err != nil {
			x.FCSTATUS = false
		}
		x.FCSTATUS = isComplete > 0
		glChecked = append(glChecked, x)
	}
	r.Data = &glChecked
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GlrefHeaderPostController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Post"
	var frm models.GlRefForm
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	user_id, er := services.ValidateToken(token)
	if er != nil {
		r.Message = "Token is Expired"
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}

	var user models.User
	if err := configs.Store.Select("DepartmentID").Preload("Department").First(&user, &models.User{ID: fmt.Sprintf("%s", user_id)}).Error; err != nil {
		r.Message = fmt.Sprintf("%s %s", frm.FCWHOUSE, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	db := WHSDb(c)
	// Begin transaction
	tx := db.Begin()
	var sec models.Sect
	if err := tx.First(&sec, &models.Sect{FCDEPT: user.Department.Code}).Error; err != nil {
		r.Message = fmt.Sprintf("%s %s", user.Department.Code, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// var whs models.Whouse
	// if err := tx.Select("FCSKID").First(&whs, &models.Whouse{FCCODE: frm.FCWHOUSE}).Error; err != nil {
	// 	r.Message = fmt.Sprintf("%s %s", frm.FCWHOUSE, err.Error())
	// 	return c.Status(fiber.StatusNotFound).JSON(&r)
	// }

	var branch models.Branch
	if err := tx.Select("FCSKID").First(&branch, &models.Branch{FCCODE: frm.FCBRANCH}).Error; err != nil {
		r.Message = fmt.Sprintf("%s %s", frm.FCBRANCH, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var book models.Book
	if err := tx.Preload("RefType").First(&book, &models.Book{FCCODE: frm.FCBOOK}).Error; err != nil {
		r.Message = fmt.Sprintf("%s %s", frm.FCBOOK, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var rnn int64
	if err := tx.Select("FCCODE").Where("FCCODE LIKE ?", (time.Now().Format("20060102"))[2:6]+"%").Model(&models.Glref{}).Count(&rnn).Error; err != nil {
		panic(err)
	}

	frm.FCCODE = fmt.Sprintf("%s%04d", (time.Now().Format("20060102"))[2:6], (rnn + 1))
	frm.FCREFNO = strings.ReplaceAll(fmt.Sprintf("%s%s", book.FCPREFIX, frm.FCCODE), " ", "")
	frm.FCGID, _ = g.New(26)

	var fcamt float64 = 0
	for _, i := range frm.REFPROD {
		fcamt += i.QTY
	}

	var glref models.Glref
	glref.FCLUPDAPP = "$0"
	glref.FCDATASER = configs.FCDATASER
	glref.FCCODE = frm.FCCODE
	glref.FCGID = frm.FCGID
	glref.FCREFNO = frm.FCREFNO
	glref.FCREFTYPE = book.RefType.FCCODE
	glref.FCLUPDAPP = "$0"
	glref.FCRFTYPE = book.RefType.FCRFTYPE
	glref.FCSTEP = frm.FCSTEP
	glref.FDDATE = frm.FDDATE
	glref.FCBRANCH = branch.FCSKID
	glref.FNAMT = fcamt
	glref.FCJOB = book.FCJOB
	glref.FCCOOR = sec.FCCODE
	glref.FCCORP = book.FCCORP
	glref.FCDEPT = sec.FCDEPT
	glref.FCSECT = sec.FCSKID
	glref.FCBOOK = book.FCSKID

	// Form WHS
	var glWhs models.WHouse
	if err := tx.Select("FCSKID").First(&glWhs, &models.WHouse{FCCODE: "YYY"}).Error; err != nil {
		tx.Rollback()
		r.Message = fmt.Sprintf("Not Found YYY := %s", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	glref.FCFRWHOUSE = glWhs.FCSKID

	if book.FCWHOUSE != "" {
		glref.FCTOWHOUSE = book.FCWHOUSE
	} else {
		var glToWhs models.WHouse
		if err := tx.Select("FCSKID").First(&glToWhs, &models.WHouse{FCCODE: "003"}).Error; err != nil {
			tx.Rollback()
			r.Message = fmt.Sprintf("Not Found 003 := %s", err.Error())
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		glref.FCTOWHOUSE = glToWhs.FCSKID
	}

	if frm.FCREMARK != "" {
		// glref.FMMEMDATA = fmt.Sprintf("Rem%sRem", frm.FCREMARK)
		glref.FMMEMDATA = frm.FCREMARK
	}
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

		if v.FCSKID == "" {
			tx.Rollback()
			r.Message = "Notfound Prod"
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		fcWhs := glref.FCFRWHOUSE
		if frm.FCSTEP == "I" {
			fcWhs = glref.FCTOWHOUSE
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
			refProd.FCCORP = book.FCCORP
			refProd.FCBRANCH = branch.FCSKID
			refProd.FCWHOUSE = fcWhs
			refProd.FCJOB = book.FCJOB
			refProd.FCSECT = sec.FCSKID
			refProd.FCDEPT = sec.FCDEPT
			refProd.FCPROD = v.FCSKID
			refProd.FCPRODTYPE = v.FCPRTYPE
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
			tx.First(&stock, &models.Stock{FCPROD: v.FCSKID, FCWHOUSE: fcWhs})
			stock.FCDATASER = configs.FCDATASER
			stock.FCCORP = v.FCCORP
			stock.FCBRANCH = branch.FCSKID
			stock.FCWHOUSE = fcWhs
			stock.FCPROD = v.FCSKID
			stock.FDDATE = glref.FDDATE
			switch frm.FCSTEP {
			case "I":
				stock.FNQTY = stock.FNQTY + i.QTY
			default:
				if stock.FNQTY > 0 {
					stock.FNQTY = stock.FNQTY - i.QTY
				}
			}

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
	msg := fmt.Sprintf("\nบันทึก%s\nเลขที่: %s \nสินค้า: %d รายการ\nจำนวน: %d\nเรียบร้อยแล้ว\n%s", book.FCNAME, glref.FCREFNO, len(frm.REFPROD), int(fcamt), time.Now().Format("2006-01-02 15:04:05"))
	var line models.Linenotify
	if err := configs.Store.First(&line, &models.Linenotify{Jobs: book.FCREFTYPE}).Error; err == nil {
		if line.Token != "" {
			go services.LineNotify(line.Token, msg)
		}
	}
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

func GlrefHeaderTransferController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.GlRefForm
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	user_id, er := services.ValidateToken(token)
	if er != nil {
		r.Message = "Token is Expired"
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}

	db := WHSDb(c)
	tx := db.Begin()
	var glref models.Glref
	var isDuplicate int64
	if err := tx.Select("FCSKID").Find(&glref, &models.Glref{FCREFNO: frm.FCREFNO}).Count(&isDuplicate).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if isDuplicate > 0 {
		tx.Rollback()
		r.Message = fmt.Sprintf("%s is duplicate!", strings.ToUpper(frm.FCREFNO))
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if err := tx.First(&glref, &models.Glref{FCSKID: c.Params("id")}).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// insert to glref logger
	var glrefHistory models.GlrefHistory
	glrefHistory.FCSKID = glref.FCSKID
	glrefHistory.OLDREFNO = glref.FCREFNO
	glrefHistory.IsComplete = true
	glrefHistory.UpdateByID = fmt.Sprintf("%s", user_id)
	if err := configs.Store.FirstOrCreate(&glrefHistory, &models.GlrefHistory{FCSKID: glref.FCSKID}).Error; err != nil {
		tx.Rollback()
		r.Message = er.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// update glref new invoice
	glref.FCREFNO = frm.FCREFNO
	glref.FTLASTUPD = time.Now()
	glref.FTLASTEDIT = time.Now()
	if err := tx.Save(&glref).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	tx.Commit()
	return c.Status(fiber.StatusOK).JSON(&r)
}
