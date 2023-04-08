package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Taxop struct {
	FCADDAPVBY string `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
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
	FCDELAPVBY string `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCDTYPE1   string `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEDTAPVBY string `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCISUSED   string `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCLID      string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMANFLAG  string `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCRDCODE   string `gorm:"column:FCRDCODE;" json:"fcrdcode"  form:"fcrdcode" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCTAXTYPE  string `gorm:"column:FCTAXTYPE;" json:"fctaxtype"  form:"fctaxtype" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCYEAR     string `gorm:"column:FCYEAR;" json:"fcyear"  form:"fcyear" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNAMT1     string `gorm:"column:FNAMT1;" json:"fnamt1"  form:"fnamt1" `
	FNAMT2     string `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
	FNAMT3     string `gorm:"column:FNAMT3;" json:"fnamt3"  form:"fnamt3" `
	FNU1CNT    string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Taxop) TableName() string {
	return "TAXOP"
}

func (obj *Taxop) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
