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

	prodStock := auth.Group("/stock")
	prodStock.Get("", c.StockGetController)
	prodStock.Post("", c.StockPostController)
	prodStock.Put("/:id", c.StockPutController)
	prodStock.Delete("/:id", c.StockDeleteController)
}
