package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Empl struct {
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCADDR1    string    `gorm:"column:FCADDR1;" json:"fcaddr1"  form:"fcaddr1" `
	FCADDR12   string    `gorm:"column:FCADDR12;" json:"fcaddr12"  form:"fcaddr12" `
	FCADDR2    string    `gorm:"column:FCADDR2;" json:"fcaddr2"  form:"fcaddr2" `
	FCADDR22   string    `gorm:"column:FCADDR22;" json:"fcaddr22"  form:"fcaddr22" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCDRIVENO  string    `gorm:"column:FCDRIVENO;" json:"fcdriveno"  form:"fcdriveno" `
	FCDRVADDR1 string    `gorm:"column:FCDRVADDR1;" json:"fcdrvaddr1"  form:"fcdrvaddr1" `
	FCDRVADDR2 string    `gorm:"column:FCDRVADDR2;" json:"fcdrvaddr2"  form:"fcdrvaddr2" `
	FCDTYPE    string    `gorm:"column:FCDTYPE;" json:"fcdtype"  form:"fcdtype" `
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
	FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCEMAIL    string    `gorm:"column:FCEMAIL;" json:"fcemail"  form:"fcemail" `
	FCEMPL     string    `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCEMPLTYPE string    `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
	FCEMZONE   string    `gorm:"column:FCEMZONE;" json:"fcemzone"  form:"fcemzone" `
	FCENCOTYPE string    `gorm:"column:FCENCOTYPE;" json:"fcencotype"  form:"fcencotype" `
	FCFAX      string    `gorm:"column:FCFAX;" json:"fcfax"  form:"fcfax" `
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCFCHRS    string    `gorm:"column:FCFCHRS;" json:"fcfchrs"  form:"fcfchrs" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCIDCARD   string    `gorm:"column:FCIDCARD;" json:"fcidcard"  form:"fcidcard" `
	FCISFMUSER string    `gorm:"column:FCISFMUSER;" json:"fcisfmuser"  form:"fcisfmuser" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCMOBILE   string    `gorm:"column:FCMOBILE;" json:"fcmobile"  form:"fcmobile" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string    `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPERMITBY string    `gorm:"column:FCPERMITBY;" json:"fcpermitby"  form:"fcpermitby" `
	FCPICTURE  string    `gorm:"column:FCPICTURE;" json:"fcpicture"  form:"fcpicture" `
	FCPWZIPSAL string    `gorm:"column:FCPWZIPSAL;" json:"fcpwzipsal"  form:"fcpwzipsal" `
	FCRCODE    string    `gorm:"column:FCRCODE;" json:"fcrcode"  form:"fcrcode" `
	FCREMARK   string    `gorm:"column:FCREMARK;" json:"fcremark"  form:"fcremark" `
	FCREMARK2  string    `gorm:"column:FCREMARK2;" json:"fcremark2"  form:"fcremark2" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSNAME    string    `gorm:"column:FCSNAME;" json:"fcsname"  form:"fcsname" `
	FCSNAME2   string    `gorm:"column:FCSNAME2;" json:"fcsname2"  form:"fcsname2" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCTEL      string    `gorm:"column:FCTEL;" json:"fctel"  form:"fctel" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
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
	FCZIP      string    `gorm:"column:FCZIP;" json:"fczip"  form:"fczip" `
	FDBEGDATE  time.Time `gorm:"column:FDBEGDATE;" json:"fdbegdate"  form:"fdbegdate" default:"now"`
	FDBIRTH    string    `gorm:"column:FDBIRTH;" json:"fdbirth"  form:"fdbirth" `
	FDEXPIRE   string    `gorm:"column:FDEXPIRE;" json:"fdexpire"  form:"fdexpire" `
	FDEXPIREDA string    `gorm:"column:FDEXPIREDA;" json:"fdexpireda"  form:"fdexpireda" `
	FDINACTIVE string    `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	FDPERMIT   string    `gorm:"column:FDPERMIT;" json:"fdpermit"  form:"fdpermit" `
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMHISTORYD string    `gorm:"column:FMHISTORYD;" json:"fmhistoryd"  form:"fmhistoryd" `
	FMSIGNATUR string    `gorm:"column:FMSIGNATUR;" json:"fmsignatur"  form:"fmsignatur" `
	FNCOMMISSI float64   `gorm:"column:FNCOMMISSI;" json:"fncommissi"  form:"fncommissi" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
}

func (Empl) TableName() string {
	return "EMPL"
}

func (obj *Empl) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
