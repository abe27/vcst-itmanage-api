package routers

import (
	c "github.com/abe27/vcst/api.v1/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	r.Get("/", c.HelloController)

	api := r.Group("/api/v1")

	// Route product
	prod := api.Group("/product")
	prod.Get("", c.ProductGetController)
	prod.Post("", c.ProductPostController)
	prod.Put("/:id", c.ProductPutController)
	prod.Delete("/:id", c.ProductDeleteController)

	// prodBVS := api.Group("/product/bvs")
	// prodBVS.Get("", c.ProductGetBVSController)
	// prodBVS.Post("", c.ProductPostBVSController)
	// prodBVS.Put("/:id", c.ProductPutBVSController)
	// prodBVS.Delete("/:id", c.ProductDeleteBVSController)

	// prodVCS := api.Group("/product/vcs")
	// prodVCS.Get("", c.ProductGetVCSController)
	// prodVCS.Post("", c.ProductPostVCSController)
	// prodVCS.Put("/:id", c.ProductPutVCSController)
	// prodVCS.Delete("/:id", c.ProductDeleteVCSController)

	// prodVCST := api.Group("/product/vcst")
	// prodVCST.Get("", c.ProductGetVCSTController)
	// prodVCST.Post("", c.ProductPostVCSTController)
	// prodVCST.Put("/:id", c.ProductPutVCSTController)
	// prodVCST.Delete("/:id", c.ProductDeleteVCSTController)
}
