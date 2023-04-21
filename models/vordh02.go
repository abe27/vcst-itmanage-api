package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vordh02 struct {
	FCA1_11    string    `gorm:"column:FCA1_11;" json:"fca1_11"  form:"fca1_11" `
	FCA1_12    string    `gorm:"column:FCA1_12;" json:"fca1_12"  form:"fca1_12" `
	FCA1_13    string    `gorm:"column:FCA1_13;" json:"fca1_13"  form:"fca1_13" `
	FCA1_14    string    `gorm:"column:FCA1_14;" json:"fca1_14"  form:"fca1_14" `
	FCA1_15    string    `gorm:"column:FCA1_15;" json:"fca1_15"  form:"fca1_15" `
	FCA1_16    string    `gorm:"column:FCA1_16;" json:"fca1_16"  form:"fca1_16" `
	FCA1_17    string    `gorm:"column:FCA1_17;" json:"fca1_17"  form:"fca1_17" `
	FCA1_18    string    `gorm:"column:FCA1_18;" json:"fca1_18"  form:"fca1_18" `
	FCA1_19    string    `gorm:"column:FCA1_19;" json:"fca1_19"  form:"fca1_19" `
	FCA1_20    string    `gorm:"column:FCA1_20;" json:"fca1_20"  form:"fca1_20" `
	FCA1_21    string    `gorm:"column:FCA1_21;" json:"fca1_21"  form:"fca1_21" `
	FCA1_22    string    `gorm:"column:FCA1_22;" json:"fca1_22"  form:"fca1_22" `
	FCA1_23    string    `gorm:"column:FCA1_23;" json:"fca1_23"  form:"fca1_23" `
	FCA1_24    string    `gorm:"column:FCA1_24;" json:"fca1_24"  form:"fca1_24" `
	FCA1_25    string    `gorm:"column:FCA1_25;" json:"fca1_25"  form:"fca1_25" `
	FCA1_26    string    `gorm:"column:FCA1_26;" json:"fca1_26"  form:"fca1_26" `
	FCA1_27    string    `gorm:"column:FCA1_27;" json:"fca1_27"  form:"fca1_27" `
	FCA1_28    string    `gorm:"column:FCA1_28;" json:"fca1_28"  form:"fca1_28" `
	FCA1_29    string    `gorm:"column:FCA1_29;" json:"fca1_29"  form:"fca1_29" `
	FCA1_30    string    `gorm:"column:FCA1_30;" json:"fca1_30"  form:"fca1_30" `
	FCA1_31    string    `gorm:"column:FCA1_31;" json:"fca1_31"  form:"fca1_31" `
	FCA1_32    string    `gorm:"column:FCA1_32;" json:"fca1_32"  form:"fca1_32" `
	FCA1_33    string    `gorm:"column:FCA1_33;" json:"fca1_33"  form:"fca1_33" `
	FCA1_34    string    `gorm:"column:FCA1_34;" json:"fca1_34"  form:"fca1_34" `
	FCA1_35    string    `gorm:"column:FCA1_35;" json:"fca1_35"  form:"fca1_35" `
	FCA1_36    string    `gorm:"column:FCA1_36;" json:"fca1_36"  form:"fca1_36" `
	FCA1_37    string    `gorm:"column:FCA1_37;" json:"fca1_37"  form:"fca1_37" `
	FCA1_38    string    `gorm:"column:FCA1_38;" json:"fca1_38"  form:"fca1_38" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOOK     string    `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
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
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORDERH   string    `gorm:"column:FCORDERH;" json:"fcorderh"  form:"fcorderh" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCRCVEDI   string    `gorm:"column:FCRCVEDI;" json:"fcrcvedi"  form:"fcrcvedi" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRENEWAL  string    `gorm:"column:FCRENEWAL;" json:"fcrenewal"  form:"fcrenewal" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCTXID     string    `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
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
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTCREATED  string    `gorm:"column:FTCREATED;" json:"ftcreated"  form:"ftcreated" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vordh02) TableName() string {
	return "VORDH02"
}

func (obj *Vordh02) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
