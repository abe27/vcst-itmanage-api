package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Tablist struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCONNID   string    `gorm:"column:FCCONNID;" json:"fcconnid"  form:"fcconnid" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCIDTYPE   string    `gorm:"column:FCIDTYPE;" json:"fcidtype"  form:"fcidtype" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCREMARK   string    `gorm:"column:FCREMARK;" json:"fcremark"  form:"fcremark" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FIRECCOUNT int64     `gorm:"column:FIRECCOUNT;" json:"fireccount"  form:"fireccount" `
	FISKIDLEN  int64     `gorm:"column:FISKIDLEN;" json:"fiskidlen"  form:"fiskidlen" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNUSERSETI float64   `gorm:"column:FNUSERSETI;" json:"fnuserseti"  form:"fnuserseti" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
}

func (Tablist) TableName() string {
	return "TABLIST"
}

func (obj *Tablist) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
