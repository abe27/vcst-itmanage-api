package configs

import (
	"encoding/json"

	"github.com/abe27/vcst/api.v1/models"
	"github.com/abe27/vcst/api.v1/services"
	"gorm.io/gorm"
)

var (
	Store            *gorm.DB
	StoreAAA         *gorm.DB
	StoreFormulaAAA  *gorm.DB
	StoreBVS         *gorm.DB
	StoreFormulaBVS  *gorm.DB
	StoreVCS         *gorm.DB
	StoreFormulaVCS  *gorm.DB
	StoreVCST        *gorm.DB
	StoreFormulaVCST *gorm.DB
)

func InitDB() {
	// fmt.Println("Initializing database!")
	if !Store.Migrator().HasTable(&models.Whs{}) {
		Store.AutoMigrate(&models.Whs{})
	}

	if !Store.Migrator().HasTable(&models.Company{}) {
		Store.AutoMigrate(&models.Company{})
	}

	if !Store.Migrator().HasTable(&models.Department{}) {
		Store.AutoMigrate(&models.Department{})
	}

	if !Store.Migrator().HasTable(&models.Position{}) {
		Store.AutoMigrate(&models.Position{})
	}

	if !Store.Migrator().HasTable(&models.Section{}) {
		Store.AutoMigrate(&models.Section{})
	}

	if !Store.Migrator().HasTable(&models.Permission{}) {
		Store.AutoMigrate(&models.Permission{})
	}

	if !Store.Migrator().HasTable(&models.ErpUser{}) {
		Store.AutoMigrate(&models.ErpUser{})
	}

	if !Store.Migrator().HasTable(&models.Ticket{}) {
		Store.AutoMigrate(&models.Ticket{})
	}

	if !Store.Migrator().HasTable(&models.Member{}) {
		Store.AutoMigrate(&models.Member{})
	}

	if !Store.Migrator().HasTable(&models.User{}) {
		Store.AutoMigrate(&models.User{})
	}

	if !Store.Migrator().HasTable(&models.ActivityLogging{}) {
		Store.AutoMigrate(&models.ActivityLogging{})
	}

	if !Store.Migrator().HasTable(&models.BillingDocument{}) {
		Store.AutoMigrate(&models.BillingDocument{})
	}

	if !Store.Migrator().HasTable(&models.BillingStatus{}) {
		Store.AutoMigrate(&models.BillingStatus{})
	}
}

func Seed() {
	data, _ := services.ReadJson("public/mock/permission.json")
	var perms []models.Permission
	json.Unmarshal(data, &perms)

	for _, p := range perms {
		Store.FirstOrCreate(&p, &models.Permission{Title: p.Title})
	}

	data, _ = services.ReadJson("public/mock/company.json")
	var company []models.Company
	json.Unmarshal(data, &company)

	for _, p := range company {
		Store.FirstOrCreate(&p, &models.Company{Name: p.Name})
	}

	data, _ = services.ReadJson("public/mock/whs.json")
	var whs []models.Whs
	json.Unmarshal(data, &whs)

	for _, p := range whs {
		Store.FirstOrCreate(&p, &models.Whs{Name: p.Name})
	}

	data, _ = services.ReadJson("public/mock/position.json")
	var position []models.Position
	json.Unmarshal(data, &position)

	for _, p := range position {
		Store.FirstOrCreate(&p, &models.Position{Title: p.Title})
	}

	data, _ = services.ReadJson("public/mock/depart.json")
	var department []models.Department
	json.Unmarshal(data, &department)

	for _, p := range department {
		Store.FirstOrCreate(&p, &models.Department{Title: p.Title})
	}

	data, _ = services.ReadJson("public/mock/section.json")
	var section []models.Section
	json.Unmarshal(data, &section)

	for _, p := range section {
		Store.FirstOrCreate(&p, &models.Section{Title: p.Title})
	}

	data, _ = services.ReadJson("public/mock/billing_document.json")
	var document []models.BillingDocument
	json.Unmarshal(data, &document)

	for _, p := range document {
		Store.FirstOrCreate(&p, &models.BillingDocument{Title: p.Title})
	}

	data, _ = services.ReadJson("public/mock/billing_status.json")
	var billingStatus []models.BillingStatus
	json.Unmarshal(data, &billingStatus)

	for _, p := range billingStatus {
		Store.FirstOrCreate(&p, &models.BillingStatus{Title: p.Title})
	}
}
