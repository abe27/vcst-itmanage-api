package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Mstock struct {
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
	FCMTH      string `gorm:"column:FCMTH;" json:"fcmth"  form:"fcmth" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPOSTBY   string `gorm:"column:FCPOSTBY;" json:"fcpostby"  form:"fcpostby" `
	FCPOSTCORR string `gorm:"column:FCPOSTCORR;" json:"fcpostcorr"  form:"fcpostcorr" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCTYPE     string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCYR       string `gorm:"column:FCYR;" json:"fcyr"  form:"fcyr" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNAMT0     string `gorm:"column:FNAMT0;" json:"fnamt0"  form:"fnamt0" `
	FNAMT1     string `gorm:"column:FNAMT1;" json:"fnamt1"  form:"fnamt1" `
	FNAMT2     string `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
	FNAMT3     string `gorm:"column:FNAMT3;" json:"fnamt3"  form:"fnamt3" `
	FNAMT4     string `gorm:"column:FNAMT4;" json:"fnamt4"  form:"fnamt4" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Mstock) TableName() string {
	return "MSTOCK"
}

func (obj *Mstock) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
