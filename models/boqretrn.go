package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Boqretrn struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOOK     string `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCANCELBY string `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	FCCODE     string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCGLHEAD   string `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCREFNO    string `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVATISOUT string `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FDDATE     string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDDUEDATE  string `gorm:"column:FDDUEDATE;" json:"fdduedate"  form:"fdduedate" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FMMEMDATA2 string `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
	FMMEMDATA3 string `gorm:"column:FMMEMDATA3;" json:"fmmemdata3"  form:"fmmemdata3" `
	FMMEMDATA4 string `gorm:"column:FMMEMDATA4;" json:"fmmemdata4"  form:"fmmemdata4" `
	FMMEMDATA5 string `gorm:"column:FMMEMDATA5;" json:"fmmemdata5"  form:"fmmemdata5" `
	FNAMT      string `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNISUEAMT  string `gorm:"column:FNISUEAMT;" json:"fnisueamt"  form:"fnisueamt" `
	FNPAYAMT   string `gorm:"column:FNPAYAMT;" json:"fnpayamt"  form:"fnpayamt" `
	FNRATE     string `gorm:"column:FNRATE;" json:"fnrate"  form:"fnrate" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Boqretrn) TableName() string {
	return "BOQRETRN"
}

func (obj *Boqretrn) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
