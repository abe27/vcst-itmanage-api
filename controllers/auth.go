package controllers

import (
	"fmt"
	"strings"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterController(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmRegister
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	db := configs.Store

	var company models.Company
	if err := db.Select("id").First(&company, &models.Company{Name: frm.CompanyID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var position models.Position
	if err := db.Select("id").First(&position, &models.Position{Title: frm.PositionID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var section models.Section
	if err := db.Select("id").First(&section, &models.Section{Title: frm.SectionID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var department models.Department
	if err := db.Select("id").First(&department, &models.Department{Title: frm.DepartmentID, Company: frm.WhsID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var permission models.Permission
	if err := db.Select("id").First(&permission, &models.Permission{Title: frm.PermissionID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var whs models.Whs
	if err := db.Select("id").First(&whs, &models.Whs{Name: frm.WhsID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	password := services.HashingPassword(frm.Password)
	isMatch := services.CheckPasswordHashing(frm.Password, password)

	if !isMatch {
		r.Message = "เกิดข้อผิดพลาดหร่างการเข้ารหัส"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var user models.User
	user.FirstName = frm.FirstName
	user.LastName = frm.LastName
	user.Email = frm.Email
	user.UserName = frm.UserName
	user.Password = password
	user.CompanyID = &company.ID
	user.PositionID = &position.ID
	user.SectionID = &section.ID
	user.DepartmentID = &department.ID
	user.PermissionID = &permission.ID
	user.WhsID = &whs.ID
	user.IsActive = frm.IsActive

	file, err := c.FormFile("avatar_url")
	if err == nil {
		fName := fmt.Sprintf("./public/user/%s", file.Filename)
		if err := c.SaveFile(file, fName); err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		user.AvatarUrl = fmt.Sprintf("/user/%s", file.Filename)
	}

	if err := db.Create(&user).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "ok"
	r.Data = &user
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func VerifyToken(c *fiber.Ctx) error {
	var r models.Response
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	if _, err := services.ValidateToken(token); err != nil {
		r.Message = err.Error()
		return c.Status(401).JSON(&r)
	}

	r.Message = "ok"
	return c.Status(fiber.StatusOK).JSON(&r)
}

func LoginController(c *fiber.Ctx) error {
	var r models.Response
	var user models.UserLoginForm
	if err := c.BodyParser(&user); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	// Check AuthorizationRequired
	db := configs.Store
	var userData models.User
	if err := db.
		Preload("Company").
		Preload("Position").
		Preload("Section").
		Preload("Department").
		Preload("Permission").
		Preload("Whs").
		First(&userData, &models.User{UserName: user.UserName}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	isMatched := services.CheckPasswordHashing(user.Password, userData.Password)
	if !isMatched {
		r.Message = "Password not match!"
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Create Token
	auth := services.CreateToken(&userData)
	r.Message = "Auth success!"
	r.Data = &auth
	return c.Status(fiber.StatusOK).JSON(&r)
}
