package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Glhead struct {
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCGLHEADRE string    `gorm:"column:FCGLHEADRE;" json:"fcglheadre"  form:"fcglheadre" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCLOSED   string    `gorm:"column:FCCLOSED;" json:"fcclosed"  form:"fcclosed" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTYLE    string    `gorm:"column:FCSTYLE;" json:"fcstyle"  form:"fcstyle" `
	FCLINK     string    `gorm:"column:FCLINK;" json:"fclink"  form:"fclink" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FCACCBOOK  string    `gorm:"column:FCACCBOOK;" json:"fcaccbook"  form:"fcaccbook" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FMREMARK2  string    `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCPOSTBY   string    `gorm:"column:FCPOSTBY;" json:"fcpostby"  form:"fcpostby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCPOSTCORR string    `gorm:"column:FCPOSTCORR;" json:"fcpostcorr"  form:"fcpostcorr" `
	FCFROMPGL  string    `gorm:"column:FCFROMPGL;" json:"fcfrompgl"  form:"fcfrompgl" `
	FCCANLINK  string    `gorm:"column:FCCANLINK;" json:"fccanlink"  form:"fccanlink" `
	FCCPATADJ  string    `gorm:"column:FCCPATADJ;" json:"fccpatadj"  form:"fccpatadj" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCATSTEP   string    `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCAPPROVER string    `gorm:"column:FCAPPROVER;" json:"fcapprover"  form:"fcapprover" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FDENTRYDAT string    `gorm:"column:FDENTRYDAT;" json:"fdentrydat"  form:"fdentrydat" `
	FCENTRYTIM string    `gorm:"column:FCENTRYTIM;" json:"fcentrytim"  form:"fcentrytim" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FNLOCKED   float64   `gorm:"column:FNLOCKED;" json:"fnlocked"  form:"fnlocked" `
	FCFIX      string    `gorm:"column:FCFIX;" json:"fcfix"  form:"fcfix" `
	FCMANDATE  string    `gorm:"column:FCMANDATE;" json:"fcmandate"  form:"fcmandate" `
	FDAUDIT    string    `gorm:"column:FDAUDIT;" json:"fdaudit"  form:"fdaudit" `
	FNISAUDIT  float64   `gorm:"column:FNISAUDIT;" json:"fnisaudit"  form:"fnisaudit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FDMANDATE  string    `gorm:"column:FDMANDATE;" json:"fdmandate"  form:"fdmandate" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCCHKDT    string    `gorm:"column:FCCHKDT;" json:"fcchkdt"  form:"fcchkdt" `
	FCRECALBY  string    `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FCLOCKBY   string    `gorm:"column:FCLOCKBY;" json:"fclockby"  form:"fclockby" `
	FTLOCK     string    `gorm:"column:FTLOCK;" json:"ftlock"  form:"ftlock" `
	FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FCDOCFLOWI string    `gorm:"column:FCDOCFLOWI;" json:"fcdocflowi"  form:"fcdocflowi" `
	FMDOCFLOW  string    `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
	FCISBUDGET string    `gorm:"column:FCISBUDGET;" json:"fcisbudget"  form:"fcisbudget" `
	FMREMARK3  string    `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FCCMD1     string    `gorm:"column:FCCMD1;" json:"fccmd1"  form:"fccmd1" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Glhead) TableName() string {
	return "GLHEAD"
}
func (obj *Glhead) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
