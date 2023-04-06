package controllers

import (
	"fmt"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HelloController(c *fiber.Ctx) error {
	var r models.Response
	r.Message = fmt.Sprintf("Hello World: %d", fiber.StatusOK)
	return c.Status(fiber.StatusOK).JSON(&r)
}

func WHSDb(c *fiber.Ctx) *gorm.DB {
	whs := "VCST"
	if c.Query("whs") != "" {
		whs = c.Query("whs")
	}

	var db *gorm.DB
	switch whs {
	case "BVS":
		db = configs.StoreFormulaBVS
	case "VCS":
		db = configs.StoreFormulaVCS
	case "AAA":
		db = configs.StoreFormulaAAA
	default:
		db = configs.StoreFormulaVCST
	}

	fmt.Println(whs)
	return db
}
