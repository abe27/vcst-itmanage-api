package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Sema struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPSYS   string    `gorm:"column:FCAPPSYS;" json:"fcappsys"  form:"fcappsys" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCTRLSTR  string    `gorm:"column:FCCTRLSTR;" json:"fcctrlstr"  form:"fcctrlstr" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCCUSTENAM string    `gorm:"column:FCCUSTENAM;" json:"fccustenam"  form:"fccustenam" `
	FCCUSTTNAM string    `gorm:"column:FCCUSTTNAM;" json:"fccusttnam"  form:"fccusttnam" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCFINGLE   string    `gorm:"column:FCFINGLE;" json:"fcfingle"  form:"fcfingle" `
	FCMASTER   string    `gorm:"column:FCMASTER;" json:"fcmaster"  form:"fcmaster" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCRESET    string    `gorm:"column:FCRESET;" json:"fcreset"  form:"fcreset" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSERIALNO string    `gorm:"column:FCSERIALNO;" json:"fcserialno"  form:"fcserialno" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSYSENAME string    `gorm:"column:FCSYSENAME;" json:"fcsysename"  form:"fcsysename" `
	FCSYSTNAME string    `gorm:"column:FCSYSTNAME;" json:"fcsystname"  form:"fcsystname" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUSER     string    `gorm:"column:FCUSER;" json:"fcuser"  form:"fcuser" `
	FCUSET     string    `gorm:"column:FCUSET;" json:"fcuset"  form:"fcuset" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNRETRACE  string    `gorm:"column:FNRETRACE;" json:"fnretrace"  form:"fnretrace" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Sema) TableName() string {
	return "SEMA"
}

func (obj *Sema) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
