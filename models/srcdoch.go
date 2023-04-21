package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Srcdoch struct {
	FCACSTEP   string    `gorm:"column:FCACSTEP;" json:"fcacstep"  form:"fcacstep" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBOMSEQ   string    `gorm:"column:FCBOMSEQ;" json:"fcbomseq"  form:"fcbomseq" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCANCELBY string    `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	FCCHKDT    string    `gorm:"column:FCCHKDT;" json:"fcchkdt"  form:"fcchkdt" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCDOCFLOWI string    `gorm:"column:FCDOCFLOWI;" json:"fcdocflowi"  form:"fcdocflowi" `
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
	FCEMPL     string    `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCEMPLISSU string    `gorm:"column:FCEMPLISSU;" json:"fcemplissu"  form:"fcemplissu" `
	FCEMPLRETU string    `gorm:"column:FCEMPLRETU;" json:"fcemplretu"  form:"fcemplretu" `
	FCFRSTORE  string    `gorm:"column:FCFRSTORE;" json:"fcfrstore"  form:"fcfrstore" `
	FCFRWHOUSE string    `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCGCODE    string    `gorm:"column:FCGCODE;" json:"fcgcode"  form:"fcgcode" `
	FCGENBY    string    `gorm:"column:FCGENBY;" json:"fcgenby"  form:"fcgenby" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCGLHEADCA string    `gorm:"column:FCGLHEADCA;" json:"fcglheadca"  form:"fcglheadca" `
	FCGLREF    string    `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FCGMACHINE string    `gorm:"column:FCGMACHINE;" json:"fcgmachine"  form:"fcgmachine" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMACHINE  string    `gorm:"column:FCMACHINE;" json:"fcmachine"  form:"fcmachine" `
	FCMBOOK    string    `gorm:"column:FCMBOOK;" json:"fcmbook"  form:"fcmbook" `
	FCMLOTH    string    `gorm:"column:FCMLOTH;" json:"fcmloth"  form:"fcmloth" `
	FCMOOP     string    `gorm:"column:FCMOOP;" json:"fcmoop"  form:"fcmoop" `
	FCMOP      string    `gorm:"column:FCMOP;" json:"fcmop"  form:"fcmop" `
	FCMORDERH  string    `gorm:"column:FCMORDERH;" json:"fcmorderh"  form:"fcmorderh" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSETCODE  string    `gorm:"column:FCSETCODE;" json:"fcsetcode"  form:"fcsetcode" `
	FCSHIFT    string    `gorm:"column:FCSHIFT;" json:"fcshift"  form:"fcshift" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCSUBTXID  string    `gorm:"column:FCSUBTXID;" json:"fcsubtxid"  form:"fcsubtxid" `
	FCTIMESTAM string    `gorm:"column:FCTIMESTAM;" json:"fctimestam"  form:"fctimestam" `
	FCTOSTORE  string    `gorm:"column:FCTOSTORE;" json:"fctostore"  form:"fctostore" `
	FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCTRANWHY  string    `gorm:"column:FCTRANWHY;" json:"fctranwhy"  form:"fctranwhy" `
	FCTXID     string    `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
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
	FCWKCTRH   string    `gorm:"column:FCWKCTRH;" json:"fcwkctrh"  form:"fcwkctrh" `
	FCWORDERH  string    `gorm:"column:FCWORDERH;" json:"fcworderh"  form:"fcworderh" `
	FCXFERR2H  string    `gorm:"column:FCXFERR2H;" json:"fcxferr2h"  form:"fcxferr2h" `
	FCXFERSTEP string    `gorm:"column:FCXFERSTEP;" json:"fcxferstep"  form:"fcxferstep" `
	FDAPPROVE  string    `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDGDATE    string    `gorm:"column:FDGDATE;" json:"fdgdate"  form:"fdgdate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMDOCFLOW  string    `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FNAMT      string    `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNFQTY     string    `gorm:"column:FNFQTY;" json:"fnfqty"  form:"fnfqty" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FTXFER     string    `gorm:"column:FTXFER;" json:"ftxfer"  form:"ftxfer" `
}

func (Srcdoch) TableName() string {
	return "SRCDOCH"
}

func (obj *Srcdoch) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
