package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)
    
type Packiv01 struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCOMMODIT string `gorm:"column:FCCOMMODIT;" json:"fccommodit"  form:"fccommodit" `
     FCCONTAINN string `gorm:"column:FCCONTAINN;" json:"fccontainn"  form:"fccontainn" `
     FCCONTAINO string `gorm:"column:FCCONTAINO;" json:"fccontaino"  form:"fccontaino" `
     FCCONTAINS string `gorm:"column:FCCONTAINS;" json:"fccontains"  form:"fccontains" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCEUORDNO string `gorm:"column:FCEUORDNO;" json:"fceuordno"  form:"fceuordno" `
     FCHSCODE string `gorm:"column:FCHSCODE;" json:"fchscode"  form:"fchscode" `
     FCINVNO string `gorm:"column:FCINVNO;" json:"fcinvno"  form:"fcinvno" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORDNO string `gorm:"column:FCORDNO;" json:"fcordno"  form:"fcordno" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPACKH string `gorm:"column:FCPACKH;" json:"fcpackh"  form:"fcpackh" `
     FCPACKI string `gorm:"column:FCPACKI;" json:"fcpacki"  form:"fcpacki" `
     FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
     FCSEAL01 string `gorm:"column:FCSEAL01;" json:"fcseal01"  form:"fcseal01" `
     FCSEAL02 string `gorm:"column:FCSEAL02;" json:"fcseal02"  form:"fcseal02" `
     FCSEAL03 string `gorm:"column:FCSEAL03;" json:"fcseal03"  form:"fcseal03" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Packiv01) TableName() string{
return "PACKIV01"
}

func (obj *Packiv01) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
