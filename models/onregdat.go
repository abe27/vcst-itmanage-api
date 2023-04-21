package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Onregdat struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCENCOTYPE string    `gorm:"column:FCENCOTYPE;" json:"fcencotype"  form:"fcencotype" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCREGSTR01 string    `gorm:"column:FCREGSTR01;" json:"fcregstr01"  form:"fcregstr01" `
	FCREGSTR02 string    `gorm:"column:FCREGSTR02;" json:"fcregstr02"  form:"fcregstr02" `
	FCREGSTR03 string    `gorm:"column:FCREGSTR03;" json:"fcregstr03"  form:"fcregstr03" `
	FCREGSTR04 string    `gorm:"column:FCREGSTR04;" json:"fcregstr04"  form:"fcregstr04" `
	FCREGSTR11 string    `gorm:"column:FCREGSTR11;" json:"fcregstr11"  form:"fcregstr11" `
	FCREGSTR12 string    `gorm:"column:FCREGSTR12;" json:"fcregstr12"  form:"fcregstr12" `
	FCREGSTR13 string    `gorm:"column:FCREGSTR13;" json:"fcregstr13"  form:"fcregstr13" `
	FCREGSTR14 string    `gorm:"column:FCREGSTR14;" json:"fcregstr14"  form:"fcregstr14" `
	FCREGSTR15 string    `gorm:"column:FCREGSTR15;" json:"fcregstr15"  form:"fcregstr15" `
	FCREGSTR21 string    `gorm:"column:FCREGSTR21;" json:"fcregstr21"  form:"fcregstr21" `
	FCREGSTR22 string    `gorm:"column:FCREGSTR22;" json:"fcregstr22"  form:"fcregstr22" `
	FCREGSTR23 string    `gorm:"column:FCREGSTR23;" json:"fcregstr23"  form:"fcregstr23" `
	FCREGSTR24 string    `gorm:"column:FCREGSTR24;" json:"fcregstr24"  form:"fcregstr24" `
	FCREGSTR25 string    `gorm:"column:FCREGSTR25;" json:"fcregstr25"  form:"fcregstr25" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMDATA01   string    `gorm:"column:FMDATA01;" json:"fmdata01"  form:"fmdata01" `
	FMDATA02   string    `gorm:"column:FMDATA02;" json:"fmdata02"  form:"fmdata02" `
	FMDATA03   string    `gorm:"column:FMDATA03;" json:"fmdata03"  form:"fmdata03" `
	FMDATA04   string    `gorm:"column:FMDATA04;" json:"fmdata04"  form:"fmdata04" `
	FMDATA05   string    `gorm:"column:FMDATA05;" json:"fmdata05"  form:"fmdata05" `
	FMDATA11   string    `gorm:"column:FMDATA11;" json:"fmdata11"  form:"fmdata11" `
	FMDATA12   string    `gorm:"column:FMDATA12;" json:"fmdata12"  form:"fmdata12" `
	FMDATA13   string    `gorm:"column:FMDATA13;" json:"fmdata13"  form:"fmdata13" `
	FMDATA14   string    `gorm:"column:FMDATA14;" json:"fmdata14"  form:"fmdata14" `
	FMDATA15   string    `gorm:"column:FMDATA15;" json:"fmdata15"  form:"fmdata15" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Onregdat) TableName() string {
	return "ONREGDAT"
}

func (obj *Onregdat) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
