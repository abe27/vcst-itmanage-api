package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vmoop1i struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEMPL2RCV string    `gorm:"column:FCEMPL2RCV;" json:"fcempl2rcv"  form:"fcempl2rcv" `
	FCEMPL2SND string    `gorm:"column:FCEMPL2SND;" json:"fcempl2snd"  form:"fcempl2snd" `
	FCFRMOP    string    `gorm:"column:FCFRMOP;" json:"fcfrmop"  form:"fcfrmop" `
	FCFRWKCTRH string    `gorm:"column:FCFRWKCTRH;" json:"fcfrwkctrh"  form:"fcfrwkctrh" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMORDERH  string    `gorm:"column:FCMORDERH;" json:"fcmorderh"  form:"fcmorderh" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSHIFT    string    `gorm:"column:FCSHIFT;" json:"fcshift"  form:"fcshift" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCTIMESTAM string    `gorm:"column:FCTIMESTAM;" json:"fctimestam"  form:"fctimestam" `
	FCTOMOP    string    `gorm:"column:FCTOMOP;" json:"fctomop"  form:"fctomop" `
	FCTOWKCTRH string    `gorm:"column:FCTOWKCTRH;" json:"fctowkctrh"  form:"fctowkctrh" `
	FCTRAN2WHY string    `gorm:"column:FCTRAN2WHY;" json:"fctran2why"  form:"fctran2why" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string    `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string    `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVMOOP1H  string    `gorm:"column:FCVMOOP1H;" json:"fcvmoop1h"  form:"fcvmoop1h" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FNGOODQTY  string    `gorm:"column:FNGOODQTY;" json:"fngoodqty"  form:"fngoodqty" `
	FNHOLDQTY  string    `gorm:"column:FNHOLDQTY;" json:"fnholdqty"  form:"fnholdqty" `
	FNUMQTY    string    `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNWASTEQTY string    `gorm:"column:FNWASTEQTY;" json:"fnwasteqty"  form:"fnwasteqty" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTFINISH   string    `gorm:"column:FTFINISH;" json:"ftfinish"  form:"ftfinish" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FTSTART    string    `gorm:"column:FTSTART;" json:"ftstart"  form:"ftstart" `
}

func (Vmoop1i) TableName() string {
	return "VMOOP1I"
}

func (obj *Vmoop1i) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
