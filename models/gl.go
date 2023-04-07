package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Gl struct {
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCGLHEAD   string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FCACCHART  string    `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FNAMT      float64   `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FCPAYMENT  string    `gorm:"column:FCPAYMENT;" json:"fcpayment"  form:"fcpayment" `
	FCCLOSED   string    `gorm:"column:FCCLOSED;" json:"fcclosed"  form:"fcclosed" `
	FCISINBROW string    `gorm:"column:FCISINBROW;" json:"fcisinbrow"  form:"fcisinbrow" `
	FCISSUMOFB string    `gorm:"column:FCISSUMOFB;" json:"fcissumofb"  form:"fcissumofb" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCGRP      string    `gorm:"column:FCGRP;" json:"fcgrp"  form:"fcgrp" `
	FMDETAIL   string    `gorm:"column:FMDETAIL;" json:"fmdetail"  form:"fmdetail" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCACLINK   string    `gorm:"column:FCACLINK;" json:"fcaclink"  form:"fcaclink" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCATSTEP   string    `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FNCALLINE  float64   `gorm:"column:FNCALLINE;" json:"fncalline"  form:"fncalline" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCRECALBY  string    `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCISBUDGET string    `gorm:"column:FCISBUDGET;" json:"fcisbudget"  form:"fcisbudget" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Gl) TableName() string {
	return "GL"
}

func (obj *Gl) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
