package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Whs struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Name        string    `gorm:"not null;index;unique;size:50" json:"name"`
	Description string    `json:"description"`
	IpAddress   string    `gorm:"size:15;" json:"ip_address"`
	OnPort      string    `gorm:"size:4;" json:"on_port"`
	User        string    `gorm:"size:50;" json:"user"`
	Passord     string    `gorm:"size:50;" json:"passord"`
	DBName      string    `gorm:"size:50;" json:"db_name"`
	IsActive    bool      `gorm:"null" json:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Whs) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Company struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Name        string    `gorm:"not null;index;unique;size:50" json:"name"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Company) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Position struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Position) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Department struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Company     string    `gorm:"size:8" json:"company" form:"company"`
	Code        string    `gorm:"size:8" json:"code" form:"code"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Department) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Section struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Section) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Permission struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsAdmin     bool      `gorm:"null" json:"admin" default:"false"`
	IsRead      bool      `gorm:"null" json:"read" default:"true"`
	IsWrite     bool      `gorm:"null" json:"write"  default:"false"`
	IsEdit      bool      `gorm:"null" json:"update" default:"false"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ErpUser struct {
	ID           string      `gorm:"primaryKey;size:21;" json:"id"`
	UserName     string      `json:"username"`
	Password     string      `json:"password"`
	DepartmentID string      `json:"department_id"`
	IsActive     bool        `gorm:"null" json:"is_active" default:"false"`
	CreatedAt    time.Time   `json:"created_at" default:"now"`
	UpdatedAt    time.Time   `json:"updated_at" default:"now"`
	Department   *Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"department"`
}

func (obj *ErpUser) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
