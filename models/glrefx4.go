package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Glrefx4 struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCCUREN01  string `gorm:"column:FCCUREN01;" json:"fccuren01"  form:"fccuren01" `
	FCCUREN02  string `gorm:"column:FCCUREN02;" json:"fccuren02"  form:"fccuren02" `
	FCCUREN03  string `gorm:"column:FCCUREN03;" json:"fccuren03"  form:"fccuren03" `
	FCCUREN04  string `gorm:"column:FCCUREN04;" json:"fccuren04"  form:"fccuren04" `
	FCCUREN05  string `gorm:"column:FCCUREN05;" json:"fccuren05"  form:"fccuren05" `
	FCCUREN06  string `gorm:"column:FCCUREN06;" json:"fccuren06"  form:"fccuren06" `
	FCCUREN07  string `gorm:"column:FCCUREN07;" json:"fccuren07"  form:"fccuren07" `
	FCCUREN08  string `gorm:"column:FCCUREN08;" json:"fccuren08"  form:"fccuren08" `
	FCCUREN09  string `gorm:"column:FCCUREN09;" json:"fccuren09"  form:"fccuren09" `
	FCCUREN10  string `gorm:"column:FCCUREN10;" json:"fccuren10"  form:"fccuren10" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
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
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLREF    string `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FCLID      string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
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
	FNACOST16  string `gorm:"column:FNACOST16;" json:"fnacost16"  form:"fnacost16" `
	FNACOST17  string `gorm:"column:FNACOST17;" json:"fnacost17"  form:"fnacost17" `
	FNACOST18  string `gorm:"column:FNACOST18;" json:"fnacost18"  form:"fnacost18" `
	FNACOST19  string `gorm:"column:FNACOST19;" json:"fnacost19"  form:"fnacost19" `
	FNACOST20  string `gorm:"column:FNACOST20;" json:"fnacost20"  form:"fnacost20" `
	FNCOSTKE01 string `gorm:"column:FNCOSTKE01;" json:"fncostke01"  form:"fncostke01" `
	FNCOSTKE02 string `gorm:"column:FNCOSTKE02;" json:"fncostke02"  form:"fncostke02" `
	FNCOSTKE03 string `gorm:"column:FNCOSTKE03;" json:"fncostke03"  form:"fncostke03" `
	FNCOSTKE04 string `gorm:"column:FNCOSTKE04;" json:"fncostke04"  form:"fncostke04" `
	FNCOSTKE05 string `gorm:"column:FNCOSTKE05;" json:"fncostke05"  form:"fncostke05" `
	FNCOSTKE06 string `gorm:"column:FNCOSTKE06;" json:"fncostke06"  form:"fncostke06" `
	FNCOSTKE07 string `gorm:"column:FNCOSTKE07;" json:"fncostke07"  form:"fncostke07" `
	FNCOSTKE08 string `gorm:"column:FNCOSTKE08;" json:"fncostke08"  form:"fncostke08" `
	FNCOSTKE09 string `gorm:"column:FNCOSTKE09;" json:"fncostke09"  form:"fncostke09" `
	FNCOSTKE10 string `gorm:"column:FNCOSTKE10;" json:"fncostke10"  form:"fncostke10" `
	FNU1CNT    string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNXRATE01  string `gorm:"column:FNXRATE01;" json:"fnxrate01"  form:"fnxrate01" `
	FNXRATE02  string `gorm:"column:FNXRATE02;" json:"fnxrate02"  form:"fnxrate02" `
	FNXRATE03  string `gorm:"column:FNXRATE03;" json:"fnxrate03"  form:"fnxrate03" `
	FNXRATE04  string `gorm:"column:FNXRATE04;" json:"fnxrate04"  form:"fnxrate04" `
	FNXRATE05  string `gorm:"column:FNXRATE05;" json:"fnxrate05"  form:"fnxrate05" `
	FNXRATE06  string `gorm:"column:FNXRATE06;" json:"fnxrate06"  form:"fnxrate06" `
	FNXRATE07  string `gorm:"column:FNXRATE07;" json:"fnxrate07"  form:"fnxrate07" `
	FNXRATE08  string `gorm:"column:FNXRATE08;" json:"fnxrate08"  form:"fnxrate08" `
	FNXRATE09  string `gorm:"column:FNXRATE09;" json:"fnxrate09"  form:"fnxrate09" `
	FNXRATE10  string `gorm:"column:FNXRATE10;" json:"fnxrate10"  form:"fnxrate10" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Glrefx4) TableName() string {
	return "GLREFX4"
}

func (obj *Glrefx4) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
