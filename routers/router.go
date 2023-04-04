package routers

import (
	c "github.com/abe27/vcst/api.v1/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	r.Get("/", c.HelloController)
}
