package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Agl struct {
	FCACCHART  string    `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNAMT0     float64   `gorm:"column:FNAMT0;" json:"fnamt0"  form:"fnamt0" `
	FNAMT1     float64   `gorm:"column:FNAMT1;" json:"fnamt1"  form:"fnamt1" `
	FNAMT2     float64   `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
	FNAMT3     float64   `gorm:"column:FNAMT3;" json:"fnamt3"  form:"fnamt3" `
	FNAMT4     float64   `gorm:"column:FNAMT4;" json:"fnamt4"  form:"fnamt4" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
}

func (Agl) TableName() string {
	return "AGL"
}

func (obj *Agl) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
