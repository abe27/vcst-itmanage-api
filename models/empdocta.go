package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Empdocta struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBEGTIME  string    `gorm:"column:FCBEGTIME;" json:"fcbegtime"  form:"fcbegtime" `
	FCBEHAVE   string    `gorm:"column:FCBEHAVE;" json:"fcbehave"  form:"fcbehave" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCBRANCHF  string    `gorm:"column:FCBRANCHF;" json:"fcbranchf"  form:"fcbranchf" `
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
	FCDEPTF    string    `gorm:"column:FCDEPTF;" json:"fcdeptf"  form:"fcdeptf" `
	FCDIFFTIME string    `gorm:"column:FCDIFFTIME;" json:"fcdifftime"  form:"fcdifftime" `
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
	FCEMPLGRP  string    `gorm:"column:FCEMPLGRP;" json:"fcemplgrp"  form:"fcemplgrp" `
	FCEMPLTYPE string    `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
	FCEMPLTYPF string    `gorm:"column:FCEMPLTYPF;" json:"fcempltypf"  form:"fcempltypf" `
	FCEMPLX1   string    `gorm:"column:FCEMPLX1;" json:"fcemplx1"  form:"fcemplx1" `
	FCEMPLX1F  string    `gorm:"column:FCEMPLX1F;" json:"fcemplx1f"  form:"fcemplx1f" `
	FCENDTIME  string    `gorm:"column:FCENDTIME;" json:"fcendtime"  form:"fcendtime" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCHRTYPE   string    `gorm:"column:FCHRTYPE;" json:"fchrtype"  form:"fchrtype" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCJOBF     string    `gorm:"column:FCJOBF;" json:"fcjobf"  form:"fcjobf" `
	FCLEAVE    string    `gorm:"column:FCLEAVE;" json:"fcleave"  form:"fcleave" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPERIODTY string    `gorm:"column:FCPERIODTY;" json:"fcperiodty"  form:"fcperiodty" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCPROJF    string    `gorm:"column:FCPROJF;" json:"fcprojf"  form:"fcprojf" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREGTIME  string    `gorm:"column:FCREGTIME;" json:"fcregtime"  form:"fcregtime" `
	FCRNGDATE  string    `gorm:"column:FCRNGDATE;" json:"fcrngdate"  form:"fcrngdate" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSECTF    string    `gorm:"column:FCSECTF;" json:"fcsectf"  form:"fcsectf" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSHIFT    string    `gorm:"column:FCSHIFT;" json:"fcshift"  form:"fcshift" `
	FCSHIFTNEW string    `gorm:"column:FCSHIFTNEW;" json:"fcshiftnew"  form:"fcshiftnew" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCTYPEOT   string    `gorm:"column:FCTYPEOT;" json:"fctypeot"  form:"fctypeot" `
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
	FCWORKTAB  string    `gorm:"column:FCWORKTAB;" json:"fcworktab"  form:"fcworktab" `
	FCWORKTIME string    `gorm:"column:FCWORKTIME;" json:"fcworktime"  form:"fcworktime" `
	FDBEGDATE  string    `gorm:"column:FDBEGDATE;" json:"fdbegdate"  form:"fdbegdate" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDENDDATE  string    `gorm:"column:FDENDDATE;" json:"fdenddate"  form:"fdenddate" `
	FDREGDATE  string    `gorm:"column:FDREGDATE;" json:"fdregdate"  form:"fdregdate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FMREMARK2  string    `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FMREMARKA  string    `gorm:"column:FMREMARKA;" json:"fmremarka"  form:"fmremarka" `
	FNPOINT    string    `gorm:"column:FNPOINT;" json:"fnpoint"  form:"fnpoint" `
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
	FTBEG2     string    `gorm:"column:FTBEG2;" json:"ftbeg2"  form:"ftbeg2" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTEND2     string    `gorm:"column:FTEND2;" json:"ftend2"  form:"ftend2" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Empdocta) TableName() string {
	return "EMPDOCTA"
}

func (obj *Empdocta) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
