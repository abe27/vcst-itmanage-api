package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Orderix4 struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCREFITEM  string `gorm:"column:FCREFITEM;" json:"fcrefitem"  form:"fcrefitem" `
	FCREFTYPE  string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNACOST01  string `gorm:"column:FNACOST01;" json:"fnacost01"  form:"fnacost01" `
	FNACOST02  string `gorm:"column:FNACOST02;" json:"fnacost02"  form:"fnacost02" `
	FNACOST03  string `gorm:"column:FNACOST03;" json:"fnacost03"  form:"fnacost03" `
	FNACOST04  string `gorm:"column:FNACOST04;" json:"fnacost04"  form:"fnacost04" `
	FNACOST05  string `gorm:"column:FNACOST05;" json:"fnacost05"  form:"fnacost05" `
	FNACOST06  string `gorm:"column:FNACOST06;" json:"fnacost06"  form:"fnacost06" `
	FNACOST07  string `gorm:"column:FNACOST07;" json:"fnacost07"  form:"fnacost07" `
	FNACOST08  string `gorm:"column:FNACOST08;" json:"fnacost08"  form:"fnacost08" `
	FNACOST09  string `gorm:"column:FNACOST09;" json:"fnacost09"  form:"fnacost09" `
	FNACOST10  string `gorm:"column:FNACOST10;" json:"fnacost10"  form:"fnacost10" `
	FNACOST11  string `gorm:"column:FNACOST11;" json:"fnacost11"  form:"fnacost11" `
	FNACOST12  string `gorm:"column:FNACOST12;" json:"fnacost12"  form:"fnacost12" `
	FNACOST13  string `gorm:"column:FNACOST13;" json:"fnacost13"  form:"fnacost13" `
	FNACOST14  string `gorm:"column:FNACOST14;" json:"fnacost14"  form:"fnacost14" `
	FNACOST15  string `gorm:"column:FNACOST15;" json:"fnacost15"  form:"fnacost15" `
	FNCSTOMPRC string `gorm:"column:FNCSTOMPRC;" json:"fncstomprc"  form:"fncstomprc" `
	FNPCNCOST1 string `gorm:"column:FNPCNCOST1;" json:"fnpcncost1"  form:"fnpcncost1" `
	FNPCNCOST2 string `gorm:"column:FNPCNCOST2;" json:"fnpcncost2"  form:"fnpcncost2" `
	FNPCNCOST3 string `gorm:"column:FNPCNCOST3;" json:"fnpcncost3"  form:"fnpcncost3" `
	FNPCNCOST4 string `gorm:"column:FNPCNCOST4;" json:"fnpcncost4"  form:"fnpcncost4" `
	FNPCNCOST5 string `gorm:"column:FNPCNCOST5;" json:"fnpcncost5"  form:"fnpcncost5" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Orderix4) TableName() string {
	return "ORDERIX4"
}

func (obj *Orderix4) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
