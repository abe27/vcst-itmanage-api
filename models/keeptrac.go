package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Keeptrac struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPSKID  string    `gorm:"column:FCAPPSKID;" json:"fcappskid"  form:"fcappskid" `
	FCBEGTIME  string    `gorm:"column:FCBEGTIME;" json:"fcbegtime"  form:"fcbegtime" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCOMNAME  string    `gorm:"column:FCCOMNAME;" json:"fccomname"  form:"fccomname" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCENDTIME  string    `gorm:"column:FCENDTIME;" json:"fcendtime"  form:"fcendtime" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSYSTASK  string    `gorm:"column:FCSYSTASK;" json:"fcsystask"  form:"fcsystask" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUSER     string    `gorm:"column:FCUSER;" json:"fcuser"  form:"fcuser" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMDESC     string    `gorm:"column:FMDESC;" json:"fmdesc"  form:"fmdesc" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMENUNAM2 string    `gorm:"column:FMMENUNAM2;" json:"fmmenunam2"  form:"fmmenunam2" `
	FMMENUNAME string    `gorm:"column:FMMENUNAME;" json:"fmmenuname"  form:"fmmenuname" `
	FNCANCELNO float64   `gorm:"column:FNCANCELNO;" json:"fncancelno"  form:"fncancelno" `
	FNDELNO    float64   `gorm:"column:FNDELNO;" json:"fndelno"  form:"fndelno" `
	FNEDITNO   float64   `gorm:"column:FNEDITNO;" json:"fneditno"  form:"fneditno" `
	FNINSNO    float64   `gorm:"column:FNINSNO;" json:"fninsno"  form:"fninsno" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Keeptrac) TableName() string {
	return "KEEPTRAC"
}

func (obj *Keeptrac) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
