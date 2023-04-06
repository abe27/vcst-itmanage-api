package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type BillingDocument struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *BillingDocument) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type BillingStatus struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Value       int64     `json:"value" form:"value"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *BillingStatus) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
