package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Welfare struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
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
	FCEMPSTAT  string    `gorm:"column:FCEMPSTAT;" json:"fcempstat"  form:"fcempstat" `
	FCEMPTYPE  string    `gorm:"column:FCEMPTYPE;" json:"fcemptype"  form:"fcemptype" `
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCHASDOC   string    `gorm:"column:FCHASDOC;" json:"fchasdoc"  form:"fchasdoc" `
	FCLONGTIME string    `gorm:"column:FCLONGTIME;" json:"fclongtime"  form:"fclongtime" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string    `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPAYTYPE  string    `gorm:"column:FCPAYTYPE;" json:"fcpaytype"  form:"fcpaytype" `
	FCRECTYPE  string    `gorm:"column:FCRECTYPE;" json:"fcrectype"  form:"fcrectype" `
	FCSALADJA  string    `gorm:"column:FCSALADJA;" json:"fcsaladja"  form:"fcsaladja" `
	FCSALADJD  string    `gorm:"column:FCSALADJD;" json:"fcsaladjd"  form:"fcsaladjd" `
	FCSALADJI  string    `gorm:"column:FCSALADJI;" json:"fcsaladji"  form:"fcsaladji" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
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
	FCWELPAY   string    `gorm:"column:FCWELPAY;" json:"fcwelpay"  form:"fcwelpay" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMREFDOC   string    `gorm:"column:FMREFDOC;" json:"fmrefdoc"  form:"fmrefdoc" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FNINTERRES string    `gorm:"column:FNINTERRES;" json:"fninterres"  form:"fninterres" `
	FNMAXDOUBL string    `gorm:"column:FNMAXDOUBL;" json:"fnmaxdoubl"  form:"fnmaxdoubl" `
	FNMAXNUM   string    `gorm:"column:FNMAXNUM;" json:"fnmaxnum"  form:"fnmaxnum" `
	FNMAXPAYAD string    `gorm:"column:FNMAXPAYAD;" json:"fnmaxpayad"  form:"fnmaxpayad" `
	FNMAXPAYM  string    `gorm:"column:FNMAXPAYM;" json:"fnmaxpaym"  form:"fnmaxpaym" `
	FNMAXPRNUM string    `gorm:"column:FNMAXPRNUM;" json:"fnmaxprnum"  form:"fnmaxprnum" `
	FNMAXWAR   string    `gorm:"column:FNMAXWAR;" json:"fnmaxwar"  form:"fnmaxwar" `
	FNMINPAYB  string    `gorm:"column:FNMINPAYB;" json:"fnminpayb"  form:"fnminpayb" `
	FNMINPAYPC string    `gorm:"column:FNMINPAYPC;" json:"fnminpaypc"  form:"fnminpaypc" `
	FNMINWAR   string    `gorm:"column:FNMINWAR;" json:"fnminwar"  form:"fnminwar" `
	FNMINWKYR  string    `gorm:"column:FNMINWKYR;" json:"fnminwkyr"  form:"fnminwkyr" `
	FNPYPERIOD string    `gorm:"column:FNPYPERIOD;" json:"fnpyperiod"  form:"fnpyperiod" `
	FNRECPAY   string    `gorm:"column:FNRECPAY;" json:"fnrecpay"  form:"fnrecpay" `
	FNRPERIOD  string    `gorm:"column:FNRPERIOD;" json:"fnrperiod"  form:"fnrperiod" `
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
}

func (Welfare) TableName() string {
	return "WELFARE"
}

func (obj *Welfare) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
