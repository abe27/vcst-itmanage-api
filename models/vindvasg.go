package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vindvasg struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDRVEMPL  string    `gorm:"column:FCDRVEMPL;" json:"fcdrvempl"  form:"fcdrvempl" `
	FCDRVRMK   string    `gorm:"column:FCDRVRMK;" json:"fcdrvrmk"  form:"fcdrvrmk" `
	FCDRVSTEP  string    `gorm:"column:FCDRVSTEP;" json:"fcdrvstep"  form:"fcdrvstep" `
	FCDRVTYPE  string    `gorm:"column:FCDRVTYPE;" json:"fcdrvtype"  form:"fcdrvtype" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEXDRIVER string    `gorm:"column:FCEXDRIVER;" json:"fcexdriver"  form:"fcexdriver" `
	FCLEFTRMK  string    `gorm:"column:FCLEFTRMK;" json:"fcleftrmk"  form:"fcleftrmk" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCRECRMK   string    `gorm:"column:FCRECRMK;" json:"fcrecrmk"  form:"fcrecrmk" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCTABNAME  string    `gorm:"column:FCTABNAME;" json:"fctabname"  form:"fctabname" `
	FCTABSKID  string    `gorm:"column:FCTABSKID;" json:"fctabskid"  form:"fctabskid" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVINNO    string    `gorm:"column:FCVINNO;" json:"fcvinno"  form:"fcvinno" `
	FDDRVASGN  string    `gorm:"column:FDDRVASGN;" json:"fddrvasgn"  form:"fddrvasgn" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNMILEQTY  string    `gorm:"column:FNMILEQTY;" json:"fnmileqty"  form:"fnmileqty" `
	FTCARLEFT  string    `gorm:"column:FTCARLEFT;" json:"ftcarleft"  form:"ftcarleft" `
	FTCARREC   string    `gorm:"column:FTCARREC;" json:"ftcarrec"  form:"ftcarrec" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vindvasg) TableName() string {
	return "VINDVASG"
}

func (obj *Vindvasg) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
