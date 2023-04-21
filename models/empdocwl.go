package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Empdocwl struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCANCELBY string    `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
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
	FCEMPLTYPE string    `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
	FCEMPLX1   string    `gorm:"column:FCEMPLX1;" json:"fcemplx1"  form:"fcemplx1" `
	FCEMPPAID  string    `gorm:"column:FCEMPPAID;" json:"fcemppaid"  form:"fcemppaid" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCHRTYPE   string    `gorm:"column:FCHRTYPE;" json:"fchrtype"  form:"fchrtype" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPISCASH  string    `gorm:"column:FCPISCASH;" json:"fcpiscash"  form:"fcpiscash" `
	FCPMTHB    string    `gorm:"column:FCPMTHB;" json:"fcpmthb"  form:"fcpmthb" `
	FCPPERIOD  string    `gorm:"column:FCPPERIOD;" json:"fcpperiod"  form:"fcpperiod" `
	FCPPERIODB string    `gorm:"column:FCPPERIODB;" json:"fcpperiodb"  form:"fcpperiodb" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCPSALADJ  string    `gorm:"column:FCPSALADJ;" json:"fcpsaladj"  form:"fcpsaladj" `
	FCPYRB     string    `gorm:"column:FCPYRB;" json:"fcpyrb"  form:"fcpyrb" `
	FCREFHRTYP string    `gorm:"column:FCREFHRTYP;" json:"fcrefhrtyp"  form:"fcrefhrtyp" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFSKID  string    `gorm:"column:FCREFSKID;" json:"fcrefskid"  form:"fcrefskid" `
	FCRISCASH  string    `gorm:"column:FCRISCASH;" json:"fcriscash"  form:"fcriscash" `
	FCRMTHB    string    `gorm:"column:FCRMTHB;" json:"fcrmthb"  form:"fcrmthb" `
	FCRNGDATE  string    `gorm:"column:FCRNGDATE;" json:"fcrngdate"  form:"fcrngdate" `
	FCRPERIOD  string    `gorm:"column:FCRPERIOD;" json:"fcrperiod"  form:"fcrperiod" `
	FCRPERIODB string    `gorm:"column:FCRPERIODB;" json:"fcrperiodb"  form:"fcrperiodb" `
	FCRSALADJ  string    `gorm:"column:FCRSALADJ;" json:"fcrsaladj"  form:"fcrsaladj" `
	FCRSALADJI string    `gorm:"column:FCRSALADJI;" json:"fcrsaladji"  form:"fcrsaladji" `
	FCRYRB     string    `gorm:"column:FCRYRB;" json:"fcryrb"  form:"fcryrb" `
	FCSALADJTS string    `gorm:"column:FCSALADJTS;" json:"fcsaladjts"  form:"fcsaladjts" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
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
	FDPAID     string    `gorm:"column:FDPAID;" json:"fdpaid"  form:"fdpaid" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FMREMARKA  string    `gorm:"column:FMREMARKA;" json:"fmremarka"  form:"fmremarka" `
	FNAMT      string    `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNDEBTAMT  string    `gorm:"column:FNDEBTAMT;" json:"fndebtamt"  form:"fndebtamt" `
	FNINTERAMT string    `gorm:"column:FNINTERAMT;" json:"fninteramt"  form:"fninteramt" `
	FNOTHDEDU  string    `gorm:"column:FNOTHDEDU;" json:"fnothdedu"  form:"fnothdedu" `
	FNOTHINCO  string    `gorm:"column:FNOTHINCO;" json:"fnothinco"  form:"fnothinco" `
	FNPAIDAMT  string    `gorm:"column:FNPAIDAMT;" json:"fnpaidamt"  form:"fnpaidamt" `
	FNPPERIODA string    `gorm:"column:FNPPERIODA;" json:"fnpperioda"  form:"fnpperioda" `
	FNRETAMT   string    `gorm:"column:FNRETAMT;" json:"fnretamt"  form:"fnretamt" `
	FNRPERIODA string    `gorm:"column:FNRPERIODA;" json:"fnrperioda"  form:"fnrperioda" `
	FNRPERIODI string    `gorm:"column:FNRPERIODI;" json:"fnrperiodi"  form:"fnrperiodi" `
	FNSALARY   string    `gorm:"column:FNSALARY;" json:"fnsalary"  form:"fnsalary" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTAPPROVE  string    `gorm:"column:FTAPPROVE;" json:"ftapprove"  form:"ftapprove" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Empdocwl) TableName() string {
	return "EMPDOCWL"
}

func (obj *Empdocwl) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
