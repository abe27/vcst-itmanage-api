package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Pdsth struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBTYPE    string    `gorm:"column:FCBTYPE;" json:"fcbtype"  form:"fcbtype" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCCUM      string    `gorm:"column:FCCUM;" json:"fccum"  form:"fccum" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDRAWPIC  string    `gorm:"column:FCDRAWPIC;" json:"fcdrawpic"  form:"fcdrawpic" `
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
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCGAG      string    `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGMACHINE string    `gorm:"column:FCGMACHINE;" json:"fcgmachine"  form:"fcgmachine" `
	FCISDEFA   string    `gorm:"column:FCISDEFA;" json:"fcisdefa"  form:"fcisdefa" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMTYPE    string    `gorm:"column:FCMTYPE;" json:"fcmtype"  form:"fcmtype" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string    `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCOUTMFUM  string    `gorm:"column:FCOUTMFUM;" json:"fcoutmfum"  form:"fcoutmfum" `
	FCOUTSTUM  string    `gorm:"column:FCOUTSTUM;" json:"fcoutstum"  form:"fcoutstum" `
	FCPACKSTYL string    `gorm:"column:FCPACKSTYL;" json:"fcpackstyl"  form:"fcpackstyl" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPPROD    string    `gorm:"column:FCPPROD;" json:"fcpprod"  form:"fcpprod" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string    `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCPTYPE    string    `gorm:"column:FCPTYPE;" json:"fcptype"  form:"fcptype" `
	FCPUM      string    `gorm:"column:FCPUM;" json:"fcpum"  form:"fcpum" `
	FCREVISION string    `gorm:"column:FCREVISION;" json:"fcrevision"  form:"fcrevision" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCSUBCONTR string    `gorm:"column:FCSUBCONTR;" json:"fcsubcontr"  form:"fcsubcontr" `
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
	FCUM       string    `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVREPSECT string    `gorm:"column:FCVREPSECT;" json:"fcvrepsect"  form:"fcvrepsect" `
	FDINACTIVE string    `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	FDSTART    string    `gorm:"column:FDSTART;" json:"fdstart"  form:"fdstart" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FNBATCHINT string    `gorm:"column:FNBATCHINT;" json:"fnbatchint"  form:"fnbatchint" `
	FNBATCHMAX string    `gorm:"column:FNBATCHMAX;" json:"fnbatchmax"  form:"fnbatchmax" `
	FNBATCHMIN string    `gorm:"column:FNBATCHMIN;" json:"fnbatchmin"  form:"fnbatchmin" `
	FNBATCHSIZ string    `gorm:"column:FNBATCHSIZ;" json:"fnbatchsiz"  form:"fnbatchsiz" `
	FNCAPACITY string    `gorm:"column:FNCAPACITY;" json:"fncapacity"  form:"fncapacity" `
	FNCAPCONS2 string    `gorm:"column:FNCAPCONS2;" json:"fncapcons2"  form:"fncapcons2" `
	FNCAPCONSU string    `gorm:"column:FNCAPCONSU;" json:"fncapconsu"  form:"fncapconsu" `
	FNCONTENT  string    `gorm:"column:FNCONTENT;" json:"fncontent"  form:"fncontent" `
	FNCOSTADP1 string    `gorm:"column:FNCOSTADP1;" json:"fncostadp1"  form:"fncostadp1" `
	FNCOSTADP2 string    `gorm:"column:FNCOSTADP2;" json:"fncostadp2"  form:"fncostadp2" `
	FNCOSTADP3 string    `gorm:"column:FNCOSTADP3;" json:"fncostadp3"  form:"fncostadp3" `
	FNCOSTADP4 string    `gorm:"column:FNCOSTADP4;" json:"fncostadp4"  form:"fncostadp4" `
	FNCOSTADP5 string    `gorm:"column:FNCOSTADP5;" json:"fncostadp5"  form:"fncostadp5" `
	FNCUMQTY   string    `gorm:"column:FNCUMQTY;" json:"fncumqty"  form:"fncumqty" `
	FNLTDEMPLO string    `gorm:"column:FNLTDEMPLO;" json:"fnltdemplo"  form:"fnltdemplo" `
	FNLTDMFG   string    `gorm:"column:FNLTDMFG;" json:"fnltdmfg"  form:"fnltdmfg" `
	FNLTDQC    string    `gorm:"column:FNLTDQC;" json:"fnltdqc"  form:"fnltdqc" `
	FNLT_MOVE  string    `gorm:"column:FNLT_MOVE;" json:"fnlt_move"  form:"fnlt_move" `
	FNLT_PROC  string    `gorm:"column:FNLT_PROC;" json:"fnlt_proc"  form:"fnlt_proc" `
	FNLT_QUE   string    `gorm:"column:FNLT_QUE;" json:"fnlt_que"  form:"fnlt_que" `
	FNLT_SETUP string    `gorm:"column:FNLT_SETUP;" json:"fnlt_setup"  form:"fnlt_setup" `
	FNLT_TEAR  string    `gorm:"column:FNLT_TEAR;" json:"fnlt_tear"  form:"fnlt_tear" `
	FNLT_WAIT  string    `gorm:"column:FNLT_WAIT;" json:"fnlt_wait"  form:"fnlt_wait" `
	FNMUMQTY   string    `gorm:"column:FNMUMQTY;" json:"fnmumqty"  form:"fnmumqty" `
	FNOUTMFQTY string    `gorm:"column:FNOUTMFQTY;" json:"fnoutmfqty"  form:"fnoutmfqty" `
	FNOUTQTY   string    `gorm:"column:FNOUTQTY;" json:"fnoutqty"  form:"fnoutqty" `
	FNOUTSTQTY string    `gorm:"column:FNOUTSTQTY;" json:"fnoutstqty"  form:"fnoutstqty" `
	FNPQTY     string    `gorm:"column:FNPQTY;" json:"fnpqty"  form:"fnpqty" `
	FNSUBCOST  string    `gorm:"column:FNSUBCOST;" json:"fnsubcost"  form:"fnsubcost" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNWEIGHT   string    `gorm:"column:FNWEIGHT;" json:"fnweight"  form:"fnweight" `
	FNWIPCOST1 string    `gorm:"column:FNWIPCOST1;" json:"fnwipcost1"  form:"fnwipcost1" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Pdsth) TableName() string {
	return "PDSTH"
}

func (obj *Pdsth) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
