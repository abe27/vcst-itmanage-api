package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Prodxa struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBNDSTL   string `gorm:"column:FCBNDSTL;" json:"fcbndstl"  form:"fcbndstl" `
	FCBRAND    string `gorm:"column:FCBRAND;" json:"fcbrand"  form:"fcbrand" `
	FCCOLOR    string `gorm:"column:FCCOLOR;" json:"fccolor"  form:"fccolor" `
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
	FCGENDER   string `gorm:"column:FCGENDER;" json:"fcgender"  form:"fcgender" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMCCAT    string `gorm:"column:FCMCCAT;" json:"fcmccat"  form:"fcmccat" `
	FCMCDIV    string `gorm:"column:FCMCDIV;" json:"fcmcdiv"  form:"fcmcdiv" `
	FCMCGRP    string `gorm:"column:FCMCGRP;" json:"fcmcgrp"  form:"fcmcgrp" `
	FCMCSRC    string `gorm:"column:FCMCSRC;" json:"fcmcsrc"  form:"fcmcsrc" `
	FCMODEL    string `gorm:"column:FCMODEL;" json:"fcmodel"  form:"fcmodel" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPDBRAND  string `gorm:"column:FCPDBRAND;" json:"fcpdbrand"  form:"fcpdbrand" `
	FCPDCAT    string `gorm:"column:FCPDCAT;" json:"fcpdcat"  form:"fcpdcat" `
	FCPDCOLOR  string `gorm:"column:FCPDCOLOR;" json:"fcpdcolor"  form:"fcpdcolor" `
	FCPDDESIGN string `gorm:"column:FCPDDESIGN;" json:"fcpddesign"  form:"fcpddesign" `
	FCPDGRADE  string `gorm:"column:FCPDGRADE;" json:"fcpdgrade"  form:"fcpdgrade" `
	FCPDMODEL  string `gorm:"column:FCPDMODEL;" json:"fcpdmodel"  form:"fcpdmodel" `
	FCPDPATERN string `gorm:"column:FCPDPATERN;" json:"fcpdpatern"  form:"fcpdpatern" `
	FCPDSIZE   string `gorm:"column:FCPDSIZE;" json:"fcpdsize"  form:"fcpdsize" `
	FCPROD     string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCSEASON   string `gorm:"column:FCSEASON;" json:"fcseason"  form:"fcseason" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSIZE     string `gorm:"column:FCSIZE;" json:"fcsize"  form:"fcsize" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Prodxa) TableName() string {
	return "PRODXA"
}

func (obj *Prodxa) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
