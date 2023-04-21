package models

import (
    "fmt"
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Tax1rate struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDTYPE1   string `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNBEG1     string `gorm:"column:FNBEG1;" json:"fnbeg1"  form:"fnbeg1" `
	FNBEG2     string `gorm:"column:FNBEG2;" json:"fnbeg2"  form:"fnbeg2" `
	FNBEG3     string `gorm:"column:FNBEG3;" json:"fnbeg3"  form:"fnbeg3" `
	FNBEG4     string `gorm:"column:FNBEG4;" json:"fnbeg4"  form:"fnbeg4" `
	FNBEG5     string `gorm:"column:FNBEG5;" json:"fnbeg5"  form:"fnbeg5" `
	FNBEG6     string `gorm:"column:FNBEG6;" json:"fnbeg6"  form:"fnbeg6" `
	FNBEG7     string `gorm:"column:FNBEG7;" json:"fnbeg7"  form:"fnbeg7" `
	FNBEG8     string `gorm:"column:FNBEG8;" json:"fnbeg8"  form:"fnbeg8" `
	FNEND1     string `gorm:"column:FNEND1;" json:"fnend1"  form:"fnend1" `
	FNEND2     string `gorm:"column:FNEND2;" json:"fnend2"  form:"fnend2" `
	FNEND3     string `gorm:"column:FNEND3;" json:"fnend3"  form:"fnend3" `
	FNEND4     string `gorm:"column:FNEND4;" json:"fnend4"  form:"fnend4" `
	FNEND5     string `gorm:"column:FNEND5;" json:"fnend5"  form:"fnend5" `
	FNEND6     string `gorm:"column:FNEND6;" json:"fnend6"  form:"fnend6" `
	FNEND7     string `gorm:"column:FNEND7;" json:"fnend7"  form:"fnend7" `
	FNEND8     string `gorm:"column:FNEND8;" json:"fnend8"  form:"fnend8" `
	FNRATE1    string `gorm:"column:FNRATE1;" json:"fnrate1"  form:"fnrate1" `
	FNRATE2    string `gorm:"column:FNRATE2;" json:"fnrate2"  form:"fnrate2" `
	FNRATE3    string `gorm:"column:FNRATE3;" json:"fnrate3"  form:"fnrate3" `
	FNRATE4    string `gorm:"column:FNRATE4;" json:"fnrate4"  form:"fnrate4" `
	FNRATE5    string `gorm:"column:FNRATE5;" json:"fnrate5"  form:"fnrate5" `
	FNRATE6    string `gorm:"column:FNRATE6;" json:"fnrate6"  form:"fnrate6" `
	FNRATE7    string `gorm:"column:FNRATE7;" json:"fnrate7"  form:"fnrate7" `
	FNRATE8    string `gorm:"column:FNRATE8;" json:"fnrate8"  form:"fnrate8" `
	FNU1CNT    string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Tax1rate) TableName() string {
	return "TAX1RATE"
}

func (obj *Tax1rate) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
