package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Docflowi struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCASENUM  string `gorm:"column:FCCASENUM;" json:"fccasenum"  form:"fccasenum" `
	FCCLSSNAME string `gorm:"column:FCCLSSNAME;" json:"fcclssname"  form:"fcclssname" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDESC1    string `gorm:"column:FCDESC1;" json:"fcdesc1"  form:"fcdesc1" `
	FCDESC2    string `gorm:"column:FCDESC2;" json:"fcdesc2"  form:"fcdesc2" `
	FCDESC3    string `gorm:"column:FCDESC3;" json:"fcdesc3"  form:"fcdesc3" `
	FCDOCFLOWH string `gorm:"column:FCDOCFLOWH;" json:"fcdocflowh"  form:"fcdocflowh" `
	FCDOCTYPE  string `gorm:"column:FCDOCTYPE;" json:"fcdoctype"  form:"fcdoctype" `
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
	FCLOWDOCH  string `gorm:"column:FCLOWDOCH;" json:"fclowdoch"  form:"fclowdoch" `
	FCLOWLEVEL string `gorm:"column:FCLOWLEVEL;" json:"fclowlevel"  form:"fclowlevel" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMDESC1   string `gorm:"column:FCMDESC1;" json:"fcmdesc1"  form:"fcmdesc1" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPRCCODE  string `gorm:"column:FCPRCCODE;" json:"fcprccode"  form:"fcprccode" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTEP     string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSUBSYS   string `gorm:"column:FCSUBSYS;" json:"fcsubsys"  form:"fcsubsys" `
	FCSYSTEM   string `gorm:"column:FCSYSTEM;" json:"fcsystem"  form:"fcsystem" `
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
	FMMAILTO   string `gorm:"column:FMMAILTO;" json:"fmmailto"  form:"fmmailto" `
	FMNALERTML string `gorm:"column:FMNALERTML;" json:"fmnalertml"  form:"fmnalertml" `
	FMNALERTSM string `gorm:"column:FMNALERTSM;" json:"fmnalertsm"  form:"fmnalertsm" `
	FMSMSTO    string `gorm:"column:FMSMSTO;" json:"fmsmsto"  form:"fmsmsto" `
	FNALERTLOG string `gorm:"column:FNALERTLOG;" json:"fnalertlog"  form:"fnalertlog" `
	FNALERTOCC string `gorm:"column:FNALERTOCC;" json:"fnalertocc"  form:"fnalertocc" `
	FNHEIGHT   string `gorm:"column:FNHEIGHT;" json:"fnheight"  form:"fnheight" `
	FNIMG2SHOW string `gorm:"column:FNIMG2SHOW;" json:"fnimg2show"  form:"fnimg2show" `
	FNLEFT     string `gorm:"column:FNLEFT;" json:"fnleft"  form:"fnleft" `
	FNNALERT   string `gorm:"column:FNNALERT;" json:"fnnalert"  form:"fnnalert" `
	FNNALERTCM string `gorm:"column:FNNALERTCM;" json:"fnnalertcm"  form:"fnnalertcm" `
	FNTOP      string `gorm:"column:FNTOP;" json:"fntop"  form:"fntop" `
	FNU1CNT    string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNWIDTH    string `gorm:"column:FNWIDTH;" json:"fnwidth"  form:"fnwidth" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTNALERT   string `gorm:"column:FTNALERT;" json:"ftnalert"  form:"ftnalert" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Docflowi) TableName() string {
	return "DOCFLOWI"
}

func (obj *Docflowi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
