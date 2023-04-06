package routers

import (
	c "github.com/abe27/vcst/api.v1/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	r.Get("/", c.HelloController)

	api := r.Group("/api/v1")
	api.Get("", c.HelloController)

	// Register
	api.Post("/register", c.RegisterController)

	whs := api.Group("/whs")
	whs.Get("", c.WhsGetController)
	whs.Post("", c.WhsPostController)
	whs.Put("/:id", c.WhsPutController)
	whs.Delete("/:id", c.WhsDeleteController)

	// Route product
	prod := api.Group("/product")
	prod.Get("", c.ProductGetController)
	prod.Post("", c.ProductPostController)
	prod.Put("/:id", c.ProductPutController)
	prod.Delete("/:id", c.ProductDeleteController)

	prodType := api.Group("/productType")
	prodType.Get("", c.ProductTypeGetController)
	prodType.Post("", c.ProductTypePostController)
	prodType.Put("/:id", c.ProductTypePutController)
	prodType.Delete("/:id", c.ProductTypeDeleteController)

	prodUnit := api.Group("/unit")
	prodUnit.Get("", c.UnitGetController)
	prodUnit.Post("", c.UnitPostController)
	prodUnit.Put("/:id", c.UnitPutController)
	prodUnit.Delete("/:id", c.UnitDeleteController)

	stockWhs := api.Group("/whouse")
	stockWhs.Get("", c.WHouseGetController)
	stockWhs.Post("", c.WHousePostController)
	stockWhs.Put("/:id", c.WHousePutController)
	stockWhs.Delete("/:id", c.WHouseDeleteController)

	prodStock := api.Group("/stock")
	prodStock.Get("", c.StockGetController)
	prodStock.Post("", c.StockPostController)
	prodStock.Put("/:id", c.StockPutController)
	prodStock.Delete("/:id", c.StockDeleteController)
}
