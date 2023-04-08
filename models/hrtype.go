package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Hrtype struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBACCBOOK string `gorm:"column:FCBACCBOOK;" json:"fcbaccbook"  form:"fcbaccbook" `
	FCCANCPROC string `gorm:"column:FCCANCPROC;" json:"fccancproc"  form:"fccancproc" `
	FCCODE     string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOORTYPE string `gorm:"column:FCCOORTYPE;" json:"fccoortype"  form:"fccoortype" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASERF string `gorm:"column:FCDATASERF;" json:"fcdataserf"  form:"fcdataserf" `
	FCDELPROC  string `gorm:"column:FCDELPROC;" json:"fcdelproc"  form:"fcdelproc" `
	FCDTYPE1   string `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEDITPROC string `gorm:"column:FCEDITPROC;" json:"fceditproc"  form:"fceditproc" `
	FCFCHR     string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCFM4TYPE  string `gorm:"column:FCFM4TYPE;" json:"fcfm4type"  form:"fcfm4type" `
	FCFORNGL   string `gorm:"column:FCFORNGL;" json:"fcforngl"  form:"fcforngl" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCHASRET   string `gorm:"column:FCHASRET;" json:"fchasret"  form:"fchasret" `
	FCISCASH   string `gorm:"column:FCISCASH;" json:"fciscash"  form:"fciscash" `
	FCISSERVIC string `gorm:"column:FCISSERVIC;" json:"fcisservic"  form:"fcisservic" `
	FCLINKGRP  string `gorm:"column:FCLINKGRP;" json:"fclinkgrp"  form:"fclinkgrp" `
	FCNAME     string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCNGLNAME  string `gorm:"column:FCNGLNAME;" json:"fcnglname"  form:"fcnglname" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCREPNAME  string `gorm:"column:FCREPNAME;" json:"fcrepname"  form:"fcrepname" `
	FCREPNAME2 string `gorm:"column:FCREPNAME2;" json:"fcrepname2"  form:"fcrepname2" `
	FCRFTYPE   string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSYSMOD   string `gorm:"column:FCSYSMOD;" json:"fcsysmod"  form:"fcsysmod" `
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
	FCVATDUE   string `gorm:"column:FCVATDUE;" json:"fcvatdue"  form:"fcvatdue" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNDEBT     string `gorm:"column:FNDEBT;" json:"fndebt"  form:"fndebt" `
	FNFACTOR   string `gorm:"column:FNFACTOR;" json:"fnfactor"  form:"fnfactor" `
	FNIC       string `gorm:"column:FNIC;" json:"fnic"  form:"fnic" `
	FNLEDGER   string `gorm:"column:FNLEDGER;" json:"fnledger"  form:"fnledger" `
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

func (Hrtype) TableName() string {
	return "HRTYPE"
}

func (obj *Hrtype) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
