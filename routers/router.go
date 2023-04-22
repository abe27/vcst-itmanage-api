package routers

import (
	c "github.com/abe27/vcst/api.v1/controllers"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	r.Get("/", c.HelloController)

	api := r.Group("/api/v1")
	api.Get("", c.HelloController)

	// Register
	api.Post("/register", c.RegisterController)
	api.Post("/login", c.LoginController)
	auth := api.Use(services.AuthorizationRequired)
	auth.Get("/verify", c.VerifyToken)

	whs := auth.Group("/whs")
	whs.Get("", c.WhsGetController)
	whs.Post("", c.WhsPostController)
	whs.Put("/:id", c.WhsPutController)
	whs.Delete("/:id", c.WhsDeleteController)

	company := auth.Group("/company")
	company.Get("", c.CompanyGetController)
	company.Post("", c.CompanyPostController)
	company.Put("/:id", c.CompanyPutController)
	company.Delete("/:id", c.CompanyDeleteController)

	position := auth.Group("/position")
	position.Get("", c.PositionGetController)
	position.Post("", c.PositionPostController)
	position.Put("/:id", c.PositionPutController)
	position.Delete("/:id", c.PositionDeleteController)

	department := auth.Group("/department")
	department.Get("", c.DepartmentGetController)
	department.Post("", c.DepartmentPostController)
	department.Put("/:id", c.DepartmentPutController)
	department.Delete("/:id", c.DepartmentDeleteController)

	section := auth.Group("/section")
	section.Get("", c.SectionGetController)
	section.Post("", c.SectionPostController)
	section.Put("/:id", c.SectionPutController)
	section.Delete("/:id", c.SectionDeleteController)

	erpUser := auth.Group("/erpUser")
	erpUser.Get("", c.ErpUserGetController)
	erpUser.Post("", c.ErpUserPostController)
	erpUser.Put("/:id", c.ErpUserPutController)
	erpUser.Delete("/:id", c.ErpUserDeleteController)

	billing := auth.Group("/billing")
	document := billing.Group("/document")
	document.Get("", c.DocumentGetController)
	document.Post("", c.DocumentPostController)
	document.Put("/:id", c.DocumentPutController)
	document.Delete("/:id", c.DocumentDeleteController)

	// Route product
	prod := auth.Group("/product")
	prod.Get("", c.ProductGetController)
	prod.Post("", c.ProductPostController)
	prod.Put("/:id", c.ProductPutController)
	prod.Delete("/:id", c.ProductDeleteController)

	prodType := auth.Group("/productType")
	prodType.Get("", c.ProductTypeGetController)
	prodType.Post("", c.ProductTypePostController)
	prodType.Put("/:id", c.ProductTypePutController)
	prodType.Delete("/:id", c.ProductTypeDeleteController)

	prodUnit := auth.Group("/unit")
	prodUnit.Get("", c.UnitGetController)
	prodUnit.Post("", c.UnitPostController)
	prodUnit.Put("/:id", c.UnitPutController)
	prodUnit.Delete("/:id", c.UnitDeleteController)

	stockWhs := auth.Group("/whouse")
	stockWhs.Get("", c.WHouseGetController)
	stockWhs.Post("", c.WHousePostController)
	stockWhs.Put("/:id", c.WHousePutController)
	stockWhs.Delete("/:id", c.WHouseDeleteController)

	glRef := auth.Group("/gl")
	glRefHeader := glRef.Group("/ref")
	glRefHeader.Get("", c.GlrefHeaderGetController)
	glRefHeader.Post("", c.GlrefHeaderPostController)
	glRefHeader.Put("/:id", c.GlrefHeaderPutController)
	glRefHeader.Delete("/:id", c.GlrefHeaderDeleteController)

	glRefProd := glRef.Group("/prod")
	glRefProd.Get("", c.RefProdGetController)

	// orderDetail := order.Group("/detail")
	// orderDetail.Get("", c.OrderDetailGetController)
	// orderDetail.Post("", c.OrderDetailPostController)
	// orderDetail.Put("/:id", c.OrderDetailPutController)
	// orderDetail.Delete("/:id", c.OrderDetailDeleteController)

	prodStock := auth.Group("/stock")
	prodStock.Get("", c.StockGetController)
	prodStock.Post("", c.StockPostController)
	prodStock.Put("/:id", c.StockPutController)
	prodStock.Delete("/:id", c.StockDeleteController)

	book := auth.Group("/book")
	book.Get("", c.BookGetController)
	book.Post("", c.BookPostController)
	book.Put("/:id", c.BookPutController)
	book.Delete("/:id", c.BookDeleteController)

	order := auth.Group("/order")
	orderHead := order.Group("/head")
	orderHead.Get("", c.OrderHeadGetController)
	orderHead.Post("", c.OrderHeadPostController)
	orderHead.Put("/:id", c.OrderHeadPutController)
	orderHead.Delete("/:id", c.OrderHeadDeleteController)

	orderDetail := order.Group("/detail")
	orderDetail.Get("", c.OrderDetailGetController)
	orderDetail.Post("", c.OrderDetailPostController)
	orderDetail.Put("/:id", c.OrderDetailPutController)
	orderDetail.Delete("/:id", c.OrderDetailDeleteController)

}
