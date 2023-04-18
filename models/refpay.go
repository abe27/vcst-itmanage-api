package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Refpay struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCHILDGLR string    `gorm:"column:FCCHILDGLR;" json:"fcchildglr"  form:"fcchildglr" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPREF   string    `gorm:"column:FCDEPREF;" json:"fcdepref"  form:"fcdepref" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCGLHEAD   string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCGLREF    string    `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FCLAYH     string    `gorm:"column:FCLAYH;" json:"fclayh"  form:"fclayh" `
	FCLAYH2    string    `gorm:"column:FCLAYH2;" json:"fclayh2"  form:"fclayh2" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMNMGLH   string    `gorm:"column:FCMNMGLH;" json:"fcmnmglh"  form:"fcmnmglh" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPRCVI    string    `gorm:"column:FCPRCVI;" json:"fcprcvi"  form:"fcprcvi" `
	FCREFPYTYP string    `gorm:"column:FCREFPYTYP;" json:"fcrefpytyp"  form:"fcrefpytyp" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCWTNOTE   string    `gorm:"column:FCWTNOTE;" json:"fcwtnote"  form:"fcwtnote" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNPAYAMT   float64   `gorm:"column:FNPAYAMT;" json:"fnpayamt"  form:"fnpayamt" `
	FNPAYAMTKE float64   `gorm:"column:FNPAYAMTKE;" json:"fnpayamtke"  form:"fnpayamtke" `
	FNRETENAMT float64   `gorm:"column:FNRETENAMT;" json:"fnretenamt"  form:"fnretenamt" `
	FNVATAMT   float64   `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATAMTKE float64   `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
	FNXRATE    float64   `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Refpay) TableName() string {
	return "REFPAY"
}

func (obj *Refpay) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
