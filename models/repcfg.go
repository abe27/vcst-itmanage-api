package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Repcfg struct {
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCPI      string    `gorm:"column:FCCPI;" json:"fccpi"  form:"fccpi" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCDIR      string    `gorm:"column:FCDIR;" json:"fcdir"  form:"fcdir" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCFILENAME string    `gorm:"column:FCFILENAME;" json:"fcfilename"  form:"fcfilename" `
	FCFILETYPE string    `gorm:"column:FCFILETYPE;" json:"fcfiletype"  form:"fcfiletype" `
	FCFONTNAME string    `gorm:"column:FCFONTNAME;" json:"fcfontname"  form:"fcfontname" `
	FCFONTSTYL string    `gorm:"column:FCFONTSTYL;" json:"fcfontstyl"  form:"fcfontstyl" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPFONT    string    `gorm:"column:FCPFONT;" json:"fcpfont"  form:"fcpfont" `
	FCPRINNAME string    `gorm:"column:FCPRINNAME;" json:"fcprinname"  form:"fcprinname" `
	FCREPCODE  string    `gorm:"column:FCREPCODE;" json:"fcrepcode"  form:"fcrepcode" `
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
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNBEGPAGE  float64   `gorm:"column:FNBEGPAGE;" json:"fnbegpage"  form:"fnbegpage" `
	FNENDPAGE  float64   `gorm:"column:FNENDPAGE;" json:"fnendpage"  form:"fnendpage" `
	FNFHEIGH   float64   `gorm:"column:FNFHEIGH;" json:"fnfheigh"  form:"fnfheigh" `
	FNFILE     float64   `gorm:"column:FNFILE;" json:"fnfile"  form:"fnfile" `
	FNFONTSIZE float64   `gorm:"column:FNFONTSIZE;" json:"fnfontsize"  form:"fnfontsize" `
	FNFWIDTH   float64   `gorm:"column:FNFWIDTH;" json:"fnfwidth"  form:"fnfwidth" `
	FNGHEIGH   float64   `gorm:"column:FNGHEIGH;" json:"fngheigh"  form:"fngheigh" `
	FNGMODE    float64   `gorm:"column:FNGMODE;" json:"fngmode"  form:"fngmode" `
	FNGWIDTH   float64   `gorm:"column:FNGWIDTH;" json:"fngwidth"  form:"fngwidth" `
	FNSCREEN   float64   `gorm:"column:FNSCREEN;" json:"fnscreen"  form:"fnscreen" `
	FNSHEIGH   float64   `gorm:"column:FNSHEIGH;" json:"fnsheigh"  form:"fnsheigh" `
	FNSHOWPAGE float64   `gorm:"column:FNSHOWPAGE;" json:"fnshowpage"  form:"fnshowpage" `
	FNSOMEPAGE float64   `gorm:"column:FNSOMEPAGE;" json:"fnsomepage"  form:"fnsomepage" `
	FNSWIDTH   float64   `gorm:"column:FNSWIDTH;" json:"fnswidth"  form:"fnswidth" `
	FNTHEIGH   float64   `gorm:"column:FNTHEIGH;" json:"fntheigh"  form:"fntheigh" `
	FNTWIDTH   float64   `gorm:"column:FNTWIDTH;" json:"fntwidth"  form:"fntwidth" `
	FNTXTMODE  float64   `gorm:"column:FNTXTMODE;" json:"fntxtmode"  form:"fntxtmode" `
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
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Repcfg) TableName() string {
	return "REPCFG"
}

func (obj *Repcfg) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
