package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Prcvi struct {
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
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLAYH     string    `gorm:"column:FCLAYH;" json:"fclayh"  form:"fclayh" `
	FCLAYH2    string    `gorm:"column:FCLAYH2;" json:"fclayh2"  form:"fclayh2" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPRCVH    string    `gorm:"column:FCPRCVH;" json:"fcprcvh"  form:"fcprcvh" `
	FCREFPAY   string    `gorm:"column:FCREFPAY;" json:"fcrefpay"  form:"fcrefpay" `
	FCREFPYTYP string    `gorm:"column:FCREFPYTYP;" json:"fcrefpytyp"  form:"fcrefpytyp" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNBILLAMT  string    `gorm:"column:FNBILLAMT;" json:"fnbillamt"  form:"fnbillamt" `
	FNBILLVAMT string    `gorm:"column:FNBILLVAMT;" json:"fnbillvamt"  form:"fnbillvamt" `
	FNPAYAMT   string    `gorm:"column:FNPAYAMT;" json:"fnpayamt"  form:"fnpayamt" `
	FNPAYAMTKE string    `gorm:"column:FNPAYAMTKE;" json:"fnpayamtke"  form:"fnpayamtke" `
	FNRETENAMT string    `gorm:"column:FNRETENAMT;" json:"fnretenamt"  form:"fnretenamt" `
	FNVATAMT   string    `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATAMTKE string    `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
	FNXRATE    string    `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Prcvi) TableName() string {
	return "PRCVI"
}

func (obj *Prcvi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
