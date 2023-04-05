package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type WHouse struct {
	ID         string    `gorm:"primaryKey;column:FCSKID;size:8;" json:"fcskid"  form:"fcskid"` //FCSKID char(8) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCCODE     string    `gorm:"column:FCCODE;unique;size:35;" json:"fccode"  form:"fccode"`    //FCSNAME2 char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME     string    `gorm:"column:FCNAME" json:"fcname"  form:"fcname"`                    //FCNAME char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME2    string    `gorm:"column:FCNAME2" json:"fcname2"  form:"fcname2"`
	FTDATETIME time.Time `gorm:"column:FTDATETIME" json:"ftdatetime"  form:"ftdatetime" default:"now"` //FTDATETIME datetime DEFAULT getdate() NOT NULL,
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD" json:"ftlastupd"  form:"ftlastupd" default:"now"`    //FTLASTUPD datetime DEFAULT getdate() NULL,
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT" json:"ftlastedit"  form:"ftlastedit" default:"now"` //FTLASTEDIT datetime NULL,
}

func (WHouse) TableName() string {
	return "WHOUSE"
}

func (obj *WHouse) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.ID = id
	return
}

type Stock struct {
	ID         string    `gorm:"primaryKey;column:FCSKID;size:8;" json:"fcskid"  form:"fcskid"` //FCSKID char(8) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCCODE     string    `gorm:"column:FCCODE;unique;size:35;" json:"fccode"  form:"fccode"`    //FCSNAME2 char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME     string    `gorm:"column:FCNAME" json:"fcname"  form:"fcname"`                    //FCNAME char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME2    string    `gorm:"column:FCNAME2" json:"fcname2"  form:"fcname2"`
	FTDATETIME time.Time `gorm:"column:FTDATETIME" json:"ftdatetime"  form:"ftdatetime" default:"now"` //FTDATETIME datetime DEFAULT getdate() NOT NULL,
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD" json:"ftlastupd"  form:"ftlastupd" default:"now"`    //FTLASTUPD datetime DEFAULT getdate() NULL,
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT" json:"ftlastedit"  form:"ftlastedit" default:"now"` //FTLASTEDIT datetime NULL,
}

func (Stock) TableName() string {
	return "WHOUSE"
}

func (obj *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.ID = id
	return
}
