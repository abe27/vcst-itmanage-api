package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)
    
type Refpayi struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCREFPAY string `gorm:"column:FCREFPAY;" json:"fcrefpay"  form:"fcrefpay" `
     FCREFPROD string `gorm:"column:FCREFPROD;" json:"fcrefprod"  form:"fcrefprod" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNPAYAMT string `gorm:"column:FNPAYAMT;" json:"fnpayamt"  form:"fnpayamt" `
     FNPAYAMTKE string `gorm:"column:FNPAYAMTKE;" json:"fnpayamtke"  form:"fnpayamtke" `
     FNVATAMT string `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
     FNVATAMTKE string `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Refpayi) TableName() string{
return "REFPAYI"
}

func (obj *Refpayi) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
