package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Ptrust struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPOST     string    `gorm:"column:FCPOST;" json:"fcpost"  form:"fcpost" `
	FCPROMPT1  string    `gorm:"column:FCPROMPT1;" json:"fcprompt1"  form:"fcprompt1" `
	FCPROMPT2  string    `gorm:"column:FCPROMPT2;" json:"fcprompt2"  form:"fcprompt2" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSYSOBJ   string    `gorm:"column:FCSYSOBJ;" json:"fcsysobj"  form:"fcsysobj" `
	FCSYSTASK  string    `gorm:"column:FCSYSTASK;" json:"fcsystask"  form:"fcsystask" `
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
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNCANCELRI float64   `gorm:"column:FNCANCELRI;" json:"fncancelri"  form:"fncancelri" `
	FNDELERIG  float64   `gorm:"column:FNDELERIG;" json:"fndelerig"  form:"fndelerig" `
	FNEDITRIG  float64   `gorm:"column:FNEDITRIG;" json:"fneditrig"  form:"fneditrig" `
	FNENTERRIG float64   `gorm:"column:FNENTERRIG;" json:"fnenterrig"  form:"fnenterrig" `
	FNEXITRIG  float64   `gorm:"column:FNEXITRIG;" json:"fnexitrig"  form:"fnexitrig" `
	FNFILERIG  float64   `gorm:"column:FNFILERIG;" json:"fnfilerig"  form:"fnfilerig" `
	FNINSERTRI float64   `gorm:"column:FNINSERTRI;" json:"fninsertri"  form:"fninsertri" `
	FNPRINTRIG float64   `gorm:"column:FNPRINTRIG;" json:"fnprintrig"  form:"fnprintrig" `
	FNPWREQRIG float64   `gorm:"column:FNPWREQRIG;" json:"fnpwreqrig"  form:"fnpwreqrig" `
	FNRIGSETRI float64   `gorm:"column:FNRIGSETRI;" json:"fnrigsetri"  form:"fnrigsetri" `
	FNSHOWRIG  float64   `gorm:"column:FNSHOWRIG;" json:"fnshowrig"  form:"fnshowrig" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNVIEWRIG  float64   `gorm:"column:FNVIEWRIG;" json:"fnviewrig"  form:"fnviewrig" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
}

func (Ptrust) TableName() string {
	return "PTRUST"
}

func (obj *Ptrust) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
