package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe27/vcst/api.v1/configs"
	"github.com/abe27/vcst/api.v1/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// initial database
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPORT"), os.Getenv("SSLMODE"), os.Getenv("TZNAME"))
	if len(os.Getenv("DBPASSWORD")) > 0 {
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=%s TimeZone=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPORT"), os.Getenv("DBPASSWORD"), os.Getenv("SSLMODE"), os.Getenv("TZNAME"))
	}

	configs.Store, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: false,
		SkipDefaultTransaction:                   true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	if err != nil {
		panic("Failed to connect to database")
	}
	// configs.InitDB()
	// configs.Seed()

	configsGorm := gorm.Config{
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: false,
		SkipDefaultTransaction:                   true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		}}

	/// Setup database Master
	dnsAAA := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("MSSQL_AAA_USERNAME"), os.Getenv("MSSQL_AAA_PASSWORD"), os.Getenv("MSSQL_AAA_HOST"), os.Getenv("MSSQL_AAA_PORT"), os.Getenv("MSSQL_AAA_DBNAME"))
	dnsFormulaAAA := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("MSSQL_AAA_USERNAME"), os.Getenv("MSSQL_AAA_PASSWORD"), os.Getenv("MSSQL_AAA_HOST"), os.Getenv("MSSQL_AAA_PORT"), os.Getenv("MSSQL_AAA_FORMULA"))
	dnsBVS := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("MSSQL_BVS_USERNAME"), os.Getenv("MSSQL_BVS_PASSWORD"), os.Getenv("MSSQL_BVS_HOST"), os.Getenv("MSSQL_BVS_PORT"), os.Getenv("MSSQL_BVS_DBNAME"))
	dnsFormulaBVS := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("MSSQL_BVS_USERNAME"), os.Getenv("MSSQL_BVS_PASSWORD"), os.Getenv("MSSQL_BVS_HOST"), os.Getenv("MSSQL_BVS_PORT"), os.Getenv("MSSQL_BVS_FORMULA"))
	dnsVCS := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("MSSQL_VCS_USERNAME"), os.Getenv("MSSQL_VCS_PASSWORD"), os.Getenv("MSSQL_VCS_HOST"), os.Getenv("MSSQL_VCS_PORT"), os.Getenv("MSSQL_VCS_DBNAME"))
	dnsFormulaVCS := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("MSSQL_VCS_USERNAME"), os.Getenv("MSSQL_VCS_PASSWORD"), os.Getenv("MSSQL_VCS_HOST"), os.Getenv("MSSQL_VCS_PORT"), os.Getenv("MSSQL_VCS_FORMULA"))
	dnsVCST := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("MSSQL_VCST_USERNAME"), os.Getenv("MSSQL_VCST_PASSWORD"), os.Getenv("MSSQL_VCST_HOST"), os.Getenv("MSSQL_VCST_PORT"), os.Getenv("MSSQL_VCST_DBNAME"))
	dnsFormulaVCST := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&connection+timeout=30", os.Getenv("MSSQL_VCST_USERNAME"), os.Getenv("MSSQL_VCST_PASSWORD"), os.Getenv("MSSQL_VCST_HOST"), os.Getenv("MSSQL_VCST_PORT"), os.Getenv("MSSQL_VCST_FORMULA"))

	configs.StoreAAA, err = gorm.Open(sqlserver.Open(dnsAAA), &configsGorm)
	if err != nil {
		panic("Failed to connect to StoreAAA database")
	}
	configs.StoreFormulaAAA, err = gorm.Open(sqlserver.Open(dnsFormulaAAA), &configsGorm)
	if err != nil {
		panic("Failed to connect to StoreFormulaAAA database")
	}

	configs.StoreBVS, err = gorm.Open(sqlserver.Open(dnsBVS), &configsGorm)
	if err != nil {
		panic("Failed to connect to StoreBVS database")
	}
	configs.StoreFormulaBVS, err = gorm.Open(sqlserver.Open(dnsFormulaBVS), &configsGorm)
	if err != nil {
		panic("Failed to connect to StoreFormulaBVS database")
	}

	configs.StoreVCS, err = gorm.Open(sqlserver.Open(dnsVCS), &configsGorm)
	if err != nil {
		panic("Failed to connect to StoreVCS database")
	}
	configs.StoreFormulaVCS, err = gorm.Open(sqlserver.Open(dnsFormulaVCS), &configsGorm)
	if err != nil {
		panic("Failed to connect to StoreFormulaVCS database")
	}

	configs.StoreVCST, err = gorm.Open(sqlserver.Open(dnsVCST), &configsGorm)
	if err != nil {
		panic("Failed to connect to StoreVCST database")
	}
	configs.StoreFormulaVCST, err = gorm.Open(sqlserver.Open(dnsFormulaVCST), &configsGorm)
	if err != nil {
		panic("Failed to connect to StoreFormulaVCST database")
	}

	// aside variables
	configs.FCDATASER = "$$$+"
	configs.LINE_NOTIFY_TOKEN = os.Getenv("LINE_NOTIFY_TOKEN")
}

func main() {
	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "VCST Server API Service", // add custom server header
		AppName:       "API Version 1.0",
		BodyLimit:     10 * 1024 * 1024, // this is the default limit of 10MB
	}

	app := fiber.New(config)
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Static("/", "./public")
	routers.SetupRouter(app)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("ON_PORT")))
}
