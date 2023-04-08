package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Planmoop struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOICARD  string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCRITICAL string `gorm:"column:FCCRITICAL;" json:"fccritical"  form:"fccritical" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCISCOMMEN string `gorm:"column:FCISCOMMEN;" json:"fciscommen"  form:"fciscommen" `
	FCJOB      string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMOP      string `gorm:"column:FCMOP;" json:"fcmop"  form:"fcmop" `
	FCOPSEQ    string `gorm:"column:FCOPSEQ;" json:"fcopseq"  form:"fcopseq" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPLANMOH  string `gorm:"column:FCPLANMOH;" json:"fcplanmoh"  form:"fcplanmoh" `
	FCPLANT    string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPROJ     string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSHIFT    string `gorm:"column:FCSHIFT;" json:"fcshift"  form:"fcshift" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSUBBR    string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCSUBOP    string `gorm:"column:FCSUBOP;" json:"fcsubop"  form:"fcsubop" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCWKCTRH   string `gorm:"column:FCWKCTRH;" json:"fcwkctrh"  form:"fcwkctrh" `
	FDDATE     string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMREMARK   string `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FNLT_MOVE  string `gorm:"column:FNLT_MOVE;" json:"fnlt_move"  form:"fnlt_move" `
	FNLT_PROC  string `gorm:"column:FNLT_PROC;" json:"fnlt_proc"  form:"fnlt_proc" `
	FNLT_QUE   string `gorm:"column:FNLT_QUE;" json:"fnlt_que"  form:"fnlt_que" `
	FNLT_SETUP string `gorm:"column:FNLT_SETUP;" json:"fnlt_setup"  form:"fnlt_setup" `
	FNLT_TEAR  string `gorm:"column:FNLT_TEAR;" json:"fnlt_tear"  form:"fnlt_tear" `
	FNLT_WAIT  string `gorm:"column:FNLT_WAIT;" json:"fnlt_wait"  form:"fnlt_wait" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTFINISH   string `gorm:"column:FTFINISH;" json:"ftfinish"  form:"ftfinish" `
	FTGFINISH  string `gorm:"column:FTGFINISH;" json:"ftgfinish"  form:"ftgfinish" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Planmoop) TableName() string {
	return "PLANMOOP"
}

func (obj *Planmoop) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
