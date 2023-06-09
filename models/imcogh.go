package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Imcogh struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCCURR01   string    `gorm:"column:FCCURR01;" json:"fccurr01"  form:"fccurr01" `
	FCCURR02   string    `gorm:"column:FCCURR02;" json:"fccurr02"  form:"fccurr02" `
	FCCURR03   string    `gorm:"column:FCCURR03;" json:"fccurr03"  form:"fccurr03" `
	FCCURR04   string    `gorm:"column:FCCURR04;" json:"fccurr04"  form:"fccurr04" `
	FCCURR05   string    `gorm:"column:FCCURR05;" json:"fccurr05"  form:"fccurr05" `
	FCCURR06   string    `gorm:"column:FCCURR06;" json:"fccurr06"  form:"fccurr06" `
	FCCURR07   string    `gorm:"column:FCCURR07;" json:"fccurr07"  form:"fccurr07" `
	FCCURR08   string    `gorm:"column:FCCURR08;" json:"fccurr08"  form:"fccurr08" `
	FCCURR09   string    `gorm:"column:FCCURR09;" json:"fccurr09"  form:"fccurr09" `
	FCCURR10   string    `gorm:"column:FCCURR10;" json:"fccurr10"  form:"fccurr10" `
	FCCURR11   string    `gorm:"column:FCCURR11;" json:"fccurr11"  form:"fccurr11" `
	FCCURR12   string    `gorm:"column:FCCURR12;" json:"fccurr12"  form:"fccurr12" `
	FCCURR13   string    `gorm:"column:FCCURR13;" json:"fccurr13"  form:"fccurr13" `
	FCCURR14   string    `gorm:"column:FCCURR14;" json:"fccurr14"  form:"fccurr14" `
	FCCURR15   string    `gorm:"column:FCCURR15;" json:"fccurr15"  form:"fccurr15" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNACOST01  string    `gorm:"column:FNACOST01;" json:"fnacost01"  form:"fnacost01" `
	FNACOST02  string    `gorm:"column:FNACOST02;" json:"fnacost02"  form:"fnacost02" `
	FNACOST03  string    `gorm:"column:FNACOST03;" json:"fnacost03"  form:"fnacost03" `
	FNACOST04  string    `gorm:"column:FNACOST04;" json:"fnacost04"  form:"fnacost04" `
	FNACOST05  string    `gorm:"column:FNACOST05;" json:"fnacost05"  form:"fnacost05" `
	FNACOST06  string    `gorm:"column:FNACOST06;" json:"fnacost06"  form:"fnacost06" `
	FNACOST07  string    `gorm:"column:FNACOST07;" json:"fnacost07"  form:"fnacost07" `
	FNACOST08  string    `gorm:"column:FNACOST08;" json:"fnacost08"  form:"fnacost08" `
	FNACOST09  string    `gorm:"column:FNACOST09;" json:"fnacost09"  form:"fnacost09" `
	FNACOST10  string    `gorm:"column:FNACOST10;" json:"fnacost10"  form:"fnacost10" `
	FNACOST11  string    `gorm:"column:FNACOST11;" json:"fnacost11"  form:"fnacost11" `
	FNACOST12  string    `gorm:"column:FNACOST12;" json:"fnacost12"  form:"fnacost12" `
	FNACOST13  string    `gorm:"column:FNACOST13;" json:"fnacost13"  form:"fnacost13" `
	FNACOST14  string    `gorm:"column:FNACOST14;" json:"fnacost14"  form:"fnacost14" `
	FNACOST15  string    `gorm:"column:FNACOST15;" json:"fnacost15"  form:"fnacost15" `
	FNACOSTK01 string    `gorm:"column:FNACOSTK01;" json:"fnacostk01"  form:"fnacostk01" `
	FNACOSTK02 string    `gorm:"column:FNACOSTK02;" json:"fnacostk02"  form:"fnacostk02" `
	FNACOSTK03 string    `gorm:"column:FNACOSTK03;" json:"fnacostk03"  form:"fnacostk03" `
	FNACOSTK04 string    `gorm:"column:FNACOSTK04;" json:"fnacostk04"  form:"fnacostk04" `
	FNACOSTK05 string    `gorm:"column:FNACOSTK05;" json:"fnacostk05"  form:"fnacostk05" `
	FNACOSTK06 string    `gorm:"column:FNACOSTK06;" json:"fnacostk06"  form:"fnacostk06" `
	FNACOSTK07 string    `gorm:"column:FNACOSTK07;" json:"fnacostk07"  form:"fnacostk07" `
	FNACOSTK08 string    `gorm:"column:FNACOSTK08;" json:"fnacostk08"  form:"fnacostk08" `
	FNACOSTK09 string    `gorm:"column:FNACOSTK09;" json:"fnacostk09"  form:"fnacostk09" `
	FNACOSTK10 string    `gorm:"column:FNACOSTK10;" json:"fnacostk10"  form:"fnacostk10" `
	FNACOSTK11 string    `gorm:"column:FNACOSTK11;" json:"fnacostk11"  form:"fnacostk11" `
	FNACOSTK12 string    `gorm:"column:FNACOSTK12;" json:"fnacostk12"  form:"fnacostk12" `
	FNACOSTK13 string    `gorm:"column:FNACOSTK13;" json:"fnacostk13"  form:"fnacostk13" `
	FNACOSTK14 string    `gorm:"column:FNACOSTK14;" json:"fnacostk14"  form:"fnacostk14" `
	FNACOSTK15 string    `gorm:"column:FNACOSTK15;" json:"fnacostk15"  form:"fnacostk15" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNXRATE01  string    `gorm:"column:FNXRATE01;" json:"fnxrate01"  form:"fnxrate01" `
	FNXRATE02  string    `gorm:"column:FNXRATE02;" json:"fnxrate02"  form:"fnxrate02" `
	FNXRATE03  string    `gorm:"column:FNXRATE03;" json:"fnxrate03"  form:"fnxrate03" `
	FNXRATE04  string    `gorm:"column:FNXRATE04;" json:"fnxrate04"  form:"fnxrate04" `
	FNXRATE05  string    `gorm:"column:FNXRATE05;" json:"fnxrate05"  form:"fnxrate05" `
	FNXRATE06  string    `gorm:"column:FNXRATE06;" json:"fnxrate06"  form:"fnxrate06" `
	FNXRATE07  string    `gorm:"column:FNXRATE07;" json:"fnxrate07"  form:"fnxrate07" `
	FNXRATE08  string    `gorm:"column:FNXRATE08;" json:"fnxrate08"  form:"fnxrate08" `
	FNXRATE09  string    `gorm:"column:FNXRATE09;" json:"fnxrate09"  form:"fnxrate09" `
	FNXRATE10  string    `gorm:"column:FNXRATE10;" json:"fnxrate10"  form:"fnxrate10" `
	FNXRATE11  string    `gorm:"column:FNXRATE11;" json:"fnxrate11"  form:"fnxrate11" `
	FNXRATE12  string    `gorm:"column:FNXRATE12;" json:"fnxrate12"  form:"fnxrate12" `
	FNXRATE13  string    `gorm:"column:FNXRATE13;" json:"fnxrate13"  form:"fnxrate13" `
	FNXRATE14  string    `gorm:"column:FNXRATE14;" json:"fnxrate14"  form:"fnxrate14" `
	FNXRATE15  string    `gorm:"column:FNXRATE15;" json:"fnxrate15"  form:"fnxrate15" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Imcogh) TableName() string {
	return "IMCOGH"
}

func (obj *Imcogh) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
