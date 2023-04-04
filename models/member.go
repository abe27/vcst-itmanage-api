package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Ticket struct {
	ID        string    `gorm:"size:21;primaryKey;not null;" json:"id"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	ExpireAt  time.Time `json:"expire_at"`
	IsActive  bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
}

func (obj *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Member struct {
	ID        string    `gorm:"size:21;primaryKey;not null;" json:"id"`
	Name      string    `gorm:"size:20;not null;" json:"name"`
	MobileNo  string    `gorm:"size:10;not null;unique;" json:"mobile_no"`
	IdCard    string    `gorm:"size:13;not null;unique;" json:"id_card"`
	TicketID  *string   `json:"ticket_id"`
	IsActive  bool      `gorm:"null" json:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
	Ticket    *Ticket   `gorm:"foreignKey:TicketID;references:ID;" json:"ticket"`
}

func (obj *Member) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
