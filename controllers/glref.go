package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
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

		// var isComplete int64
		// if err := configs.Store.Select("id").First(&models.GlrefHistory{}, &models.GlrefHistory{GLREF: gl.FCSKID}).Count(&isComplete).Error; err != nil {
		// 	gl.FCSTATUS = false
		// }
		// gl.FCSTATUS = isComplete > 0

		var glHistory models.GlrefHistory
		if err := configs.Store.Select("is_complete").First(&glHistory, &models.GlrefHistory{GLREF: gl.FCSKID}).Error; err != nil {
			gl.FCSTATUS = false
		}

		txt := fmt.Sprintf("%s ==> %v\n", gl.FCSKID, glHistory.IsComplete)
		fmt.Println(txt)
		gl.FCSTATUS = glHistory.IsComplete
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
			var glHistory models.GlrefHistory
			if err := configs.Store.Select("is_complete").First(&glHistory, &models.GlrefHistory{GLREF: x.FCSKID}).Error; err != nil {
				x.FCSTATUS = false
			}
			x.FCSTATUS = glHistory.IsComplete
			glChecked = append(glChecked, x)
		}
		r.Data = &glChecked
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("fcreftype") != "" {
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
			Find(&gl, &models.Glref{FCREFTYPE: c.Query("fcreftype")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		for _, x := range gl {
			var glHistory models.GlrefHistory
			if err := configs.Store.Select("is_complete").First(&glHistory, &models.GlrefHistory{GLREF: x.FCSKID}).Error; err != nil {
				x.FCSTATUS = false
			}
			x.FCSTATUS = glHistory.IsComplete
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
			var glHistory models.GlrefHistory
			if err := configs.Store.Select("is_complete").First(&glHistory, &models.GlrefHistory{GLREF: x.FCSKID}).Error; err != nil {
				x.FCSTATUS = false
			}
			x.FCSTATUS = glHistory.IsComplete
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
		var glHistory models.GlrefHistory
		if err := configs.Store.Select("is_complete").First(&glHistory, &models.GlrefHistory{GLREF: x.FCSKID}).Error; err != nil {
			x.FCSTATUS = false
		}
		x.FCSTATUS = glHistory.IsComplete
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
		r.Message = fmt.Sprintf("%s %s", user_id, err.Error())
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

	var branch models.Branch
	if err := tx.Select("FCSKID").First(&branch, &models.Branch{FCCODE: frm.FCBRANCH}).Error; err != nil {
		r.Message = fmt.Sprintf("%s %s", frm.FCBRANCH, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var book models.Book
	if err := tx.Preload("RefType").First(&book, &models.Book{FCCODE: frm.FCBOOK, FCREFTYPE: frm.FCREFTYPE}).Error; err != nil {
		r.Message = fmt.Sprintf("%s %s", frm.FCBOOK, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var rnn int64
	if err := tx.Select("FCCODE").Where("FCCODE LIKE ?", (time.Now().Format("20060102"))[2:6]+"%").Model(&models.Glref{}).Count(&rnn).Error; err != nil {
		panic(err)
	}

	frm.FCCODE = fmt.Sprintf("%s%04d", (time.Now().Format("20060102"))[2:6], (rnn + 1))
	glRefPrefix := "ADJ"
	if book.FCPREFIX != "" {
		glRefPrefix = book.FCPREFIX
	}
	frm.FCREFNO = strings.ReplaceAll(fmt.Sprintf("%s%s", glRefPrefix, frm.FCCODE), " ", "")

	var fcamt float64 = 0
	for _, i := range frm.REFPROD {
		fcamt += i.QTY
	}

	var glref models.Glref
	glref.FCDATASER = configs.FCDATASER
	glref.FCCODE = frm.FCCODE
	glref.FCREFNO = frm.FCREFNO
	glref.FCREFTYPE = book.FCREFTYPE
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
	if err := tx.Select("FCSKID").First(&glWhs, &models.WHouse{FCCODE: frm.FROMWHOUSE}).Error; err != nil {
		tx.Rollback()
		r.Message = fmt.Sprintf("Not Found %s := %s", frm.FROMWHOUSE, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	glref.FCFRWHOUSE = glWhs.FCSKID

	if strings.ReplaceAll(book.FCWHOUSE, " ", "") != "" {
		glref.FCTOWHOUSE = book.FCWHOUSE
	} else {
		var glToWhs models.WHouse
		if err := tx.Select("FCSKID").First(&glToWhs, &models.WHouse{FCCODE: frm.TOWHOUSE}).Error; err != nil {
			tx.Rollback()
			r.Message = fmt.Sprintf("Not Found %s := %s", frm.TOWHOUSE, err.Error())
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		glref.FCTOWHOUSE = glToWhs.FCSKID
	}

	if frm.FCREMARK != "" {
		// glref.FMMEMDATA = fmt.Sprintf("Rem%sRem", frm.FCREMARK)
		glref.FMMEMDATA = frm.FCREMARK
	}

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
			refProd.FCGLREF = glref.FCSKID
			refProd.FDDATE = glref.FDDATE
			refProd.FCIOTYPE = glref.FCSTEP
			refProd.FCRFTYPE = book.RefType.FCRFTYPE
			refProd.FCREFTYPE = book.FCREFTYPE
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

			// insert to glref logger
			var glrefHistory models.GlrefHistory
			glrefHistory.GLREF = glref.FCSKID
			glrefHistory.REFPROD = refProd.FCSKID
			glrefHistory.REFNO = glref.FCREFNO
			glrefHistory.PRODID = v.FCSKID
			glrefHistory.PRODUCTCODE = strings.ReplaceAll(v.FCCODE, " ", "")
			glrefHistory.PRODUCTNAME = strings.ReplaceAll(v.FCNAME, " ", "")
			glrefHistory.QTY = i.QTY
			glrefHistory.Remark = glref.FMMEMDATA
			glrefHistory.IsComplete = false
			glrefHistory.UpdateByID = fmt.Sprintf("%s", user_id)
			if err := configs.Store.Create(&glrefHistory).Error; err != nil {
				tx.Rollback()
				r.Message = er.Error()
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
	glrefHistory.GLREF = glref.FCSKID
	glrefHistory.REFNO = glref.FCREFNO
	glrefHistory.IsComplete = true
	glrefHistory.UpdateByID = fmt.Sprintf("%s", user_id)
	if err := configs.Store.FirstOrCreate(&glrefHistory, &models.GlrefHistory{GLREF: glref.FCSKID}).Error; err != nil {
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

func GlrefHeaderConfirmInvoiceController(c *fiber.Ctx) error {
	db := WHSDb(c)
	var r models.Response
	r.Message = "Patch"

	var frm models.GlrefPatchForm
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	tx := db.Begin()
	var glRef models.Glref
	if err := db.First(&glRef, &models.Glref{FCSKID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Check Part from orderi
	var listOrderI []models.OrderiView
	for _, p := range frm.REFPROD {
		var prodOrderI models.OrderiView
		if err := tx.Where("FNBACKQTY > ?", 0).First(&prodOrderI, &models.OrderiView{FCORDERH: frm.ORDERH, FCPROD: p.FCPROD}).Error; err != nil {
			r.Message = "พบสินค้าไม่ตรงกับเอกสาร"
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		if p.FNQTY > prodOrderI.FNQTY {
			r.Message = "ระบุจำนวนเกินกับเอกสาร"
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		prodOrderI.FNRECEIVEQTY = p.FNQTY
		listOrderI = append(listOrderI, prodOrderI)
	}

	// CREATE GLHEAD
	var glhead models.Glhead
	glhead.FCCORP = glRef.FCCORP
	glhead.FCBRANCH = glRef.FCBRANCH
	glhead.FDDATE = glRef.FDDATE
	var accBook models.Accbook
	if err := tx.First(&accBook, &models.Accbook{FCCODE: "PD"}).Error; err != nil {
		r.Message = "ไม่พบข้อมูล ACCBOOK!"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	glhead.FCACCBOOK = accBook.FCSKID
	var rnn int64
	if err := tx.Select("FCCODE").Where("FCCODE LIKE ?", (time.Now().Format("20060102"))[2:6]+"%").Model(&models.Glhead{}).Count(&rnn).Error; err != nil {
		panic(err)
	}

	glhead.FCCODE = fmt.Sprintf("%s%04d", (time.Now().Format("20060102"))[2:6], (rnn + 1))
	// glhead.FMREMARK
	// glhead.FMREMARK2
	// glhead.FMREMARK3
	// glhead.FCCREATEBY
	// glhead.FCCORRECTB
	glhead.FTDATETIME = time.Now()
	glhead.FTLASTEDIT = time.Now()
	glhead.FTLASTUPD = time.Now()
	if err := tx.Create(&glhead).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// fmt.Println("GLHEAD: %s", glhead.FCSKID)
	// UPDATE GLREF
	glRef.FCGLHEAD = glhead.FCSKID
	glRef.FCRFTYPE = "B"
	glRef.FCREFTYPE = "BI"
	glRef.FCSTEP = "P"
	glRef.FTLASTEDIT = time.Now()
	glRef.FTLASTUPD = time.Now()
	if err := tx.Save(&glRef).Error; err != nil {
		tx.Rollback()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	store := configs.Store.Begin()
	// var sumRefProd float64 = 0
	var sumCtn float64 = 0
	for _, i := range listOrderI {
		// fmt.Println("GLREF: %s PROD: %s ORDERH: %s ORDERI: %s", glRef.FCSKID, i.FCPROD, i.FCORDERH, i.FCSKID)
		var refProd models.Refprod
		if err := tx.Select("FCSKID,FNQTY,FCSEQ").First(&refProd, &models.Refprod{FCGLREF: glRef.FCSKID, FCPROD: i.FCPROD}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		// UPDATE REFPROD
		// sumRefProd += refProd.FNQTY
		if err := tx.Model(&refProd).Updates(&models.Refprod{
			FCGLHEAD:   glhead.FCSKID,
			FCRFTYPE:   "B",
			FCREFTYPE:  "BI",
			FTLASTEDIT: time.Now(),
			FTLASTUPD:  time.Now(),
		}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		isCompleted := true
		orderIQty := (i.FNBACKQTY - i.FNRECEIVEQTY)
		orderIStatus := "P"
		if orderIQty > 0 {
			orderIStatus = "1"
			isCompleted = false
		}

		// UPDATE ORDERI
		if err := tx.Model(&models.Orderi{FCSKID: i.FCSKID}).Updates(&models.Orderi{
			FCSTEP:     orderIStatus,
			FNBACKQTY:  orderIQty,
			FTLASTEDIT: time.Now(),
			FTLASTUPD:  time.Now(),
		}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		fmt.Println("orderIStatus: ", orderIStatus)
		sumCtn += orderIQty

		// UPDATE GLREF HISTORY
		var orderH models.Orderh
		if err := tx.Select("FCREFNO").First(&orderH, &models.Orderh{FCSKID: listOrderI[0].FCORDERH}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		// CREATE GL 1031
		var gl models.Gl
		var accChart models.Acchart
		if err := tx.First(&accChart, &models.Acchart{FCCODE: "1031"}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		// var seqGl int64
		// if err := tx.Find(&models.Gl{}, &models.Gl{FCGLHEAD: glRef.FCGLHEAD}).Count(&seqGl).Error; err != nil {
		// 	r.Message = err.Error()
		// 	return c.Status(fiber.StatusInternalServerError).JSON(&r)
		// }

		gl.FCACCHART = accChart.FCSKID
		gl.FCBRANCH = glRef.FCBRANCH
		gl.FCCORP = glRef.FCCORP
		gl.FCSECT = glRef.FCSECT
		gl.FCDEPT = glRef.FCDEPT
		gl.FCGLHEAD = glRef.FCGLHEAD
		gl.FCSEQ = refProd.FCSEQ
		gl.FDDATE = glRef.FDDATE
		gl.FNAMT = i.FNRECEIVEQTY
		if err := tx.Create(&gl).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		// CREATE NOTCUT
		var noteCut models.Notecut
		noteCut.FCBRANCH = glRef.FCBRANCH
		noteCut.FCCHILDH = listOrderI[0].FCORDERH
		noteCut.FCCHILDI = i.FCSKID
		noteCut.FCCORP = glRef.FCCORP
		noteCut.FCMASTERH = glRef.FCSKID
		noteCut.FCMASTERI = refProd.FCSKID
		noteCut.FNQTY = i.FNRECEIVEQTY
		if err := tx.Create(&noteCut).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		// Update History
		var glHistory models.GlrefHistory
		if err := store.First(&glHistory, &models.GlrefHistory{REFPROD: refProd.FCSKID}).Error; err != nil {
			tx.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		glHistory.ORDERH = listOrderI[0].FCORDERH
		glHistory.ORDERI = i.FCSKID
		glHistory.PONO = strings.ReplaceAll(orderH.FCREFNO, " ", "")
		glHistory.RECEIVEQTY = i.FNRECEIVEQTY
		glHistory.IsComplete = isCompleted

		if err := store.Save(&glHistory).Error; err != nil {
			tx.Rollback()
			store.Rollback()
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
	}

	fmt.Println("sumCtn: ", sumCtn)
	// UPDATE ORDERH
	orderStatus := "P"
	if sumCtn > 0 {
		orderStatus = "1"
	}

	if err := tx.Model(&models.Orderh{FCSKID: listOrderI[0].FCORDERH}).Updates(&models.Orderh{
		FCSTEP:     orderStatus,
		FTLASTEDIT: time.Now(),
		FTLASTUPD:  time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		store.Commit()
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	tx.Commit()
	store.Commit()
	// tx.Rollback()
	// store.Rollback()
	r.Data = nil
	return c.Status(fiber.StatusOK).JSON(&r)
}
