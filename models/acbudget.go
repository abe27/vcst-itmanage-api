package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Acbudget struct {
	FCACCHART  string `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FCACGROUP  string `gorm:"column:FCACGROUP;" json:"fcacgroup"  form:"fcacgroup" `
	FCACTYPE   string `gorm:"column:FCACTYPE;" json:"fcactype"  form:"fcactype" `
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOICARD  string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCBUDGRP   string `gorm:"column:FCBUDGRP;" json:"fcbudgrp"  form:"fcbudgrp" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCISCENTER string `gorm:"column:FCISCENTER;" json:"fciscenter"  form:"fciscenter" `
	FCJOB      string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPROJ     string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSUBBR    string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCWARNING  string `gorm:"column:FCWARNING;" json:"fcwarning"  form:"fcwarning" `
	FCYEAR     string `gorm:"column:FCYEAR;" json:"fcyear"  form:"fcyear" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNAMT01    string `gorm:"column:FNAMT01;" json:"fnamt01"  form:"fnamt01" `
	FNAMT02    string `gorm:"column:FNAMT02;" json:"fnamt02"  form:"fnamt02" `
	FNAMT03    string `gorm:"column:FNAMT03;" json:"fnamt03"  form:"fnamt03" `
	FNAMT04    string `gorm:"column:FNAMT04;" json:"fnamt04"  form:"fnamt04" `
	FNAMT05    string `gorm:"column:FNAMT05;" json:"fnamt05"  form:"fnamt05" `
	FNAMT06    string `gorm:"column:FNAMT06;" json:"fnamt06"  form:"fnamt06" `
	FNAMT07    string `gorm:"column:FNAMT07;" json:"fnamt07"  form:"fnamt07" `
	FNAMT08    string `gorm:"column:FNAMT08;" json:"fnamt08"  form:"fnamt08" `
	FNAMT09    string `gorm:"column:FNAMT09;" json:"fnamt09"  form:"fnamt09" `
	FNAMT10    string `gorm:"column:FNAMT10;" json:"fnamt10"  form:"fnamt10" `
	FNAMT11    string `gorm:"column:FNAMT11;" json:"fnamt11"  form:"fnamt11" `
	FNAMT12    string `gorm:"column:FNAMT12;" json:"fnamt12"  form:"fnamt12" `
	FNPCT01    string `gorm:"column:FNPCT01;" json:"fnpct01"  form:"fnpct01" `
	FNPCT02    string `gorm:"column:FNPCT02;" json:"fnpct02"  form:"fnpct02" `
	FNPCT03    string `gorm:"column:FNPCT03;" json:"fnpct03"  form:"fnpct03" `
	FNPCT04    string `gorm:"column:FNPCT04;" json:"fnpct04"  form:"fnpct04" `
	FNPCT05    string `gorm:"column:FNPCT05;" json:"fnpct05"  form:"fnpct05" `
	FNPCT06    string `gorm:"column:FNPCT06;" json:"fnpct06"  form:"fnpct06" `
	FNPCT07    string `gorm:"column:FNPCT07;" json:"fnpct07"  form:"fnpct07" `
	FNPCT08    string `gorm:"column:FNPCT08;" json:"fnpct08"  form:"fnpct08" `
	FNPCT09    string `gorm:"column:FNPCT09;" json:"fnpct09"  form:"fnpct09" `
	FNPCT10    string `gorm:"column:FNPCT10;" json:"fnpct10"  form:"fnpct10" `
	FNPCT11    string `gorm:"column:FNPCT11;" json:"fnpct11"  form:"fnpct11" `
	FNPCT12    string `gorm:"column:FNPCT12;" json:"fnpct12"  form:"fnpct12" `
	FNRESERVE  string `gorm:"column:FNRESERVE;" json:"fnreserve"  form:"fnreserve" `
	FNTOTAL    string `gorm:"column:FNTOTAL;" json:"fntotal"  form:"fntotal" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Acbudget) TableName() string {
	return "ACBUDGET"
}

func (obj *Acbudget) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
