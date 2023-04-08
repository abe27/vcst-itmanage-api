package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Hsalary struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCALTMETH string `gorm:"column:FCCALTMETH;" json:"fccaltmeth"  form:"fccaltmeth" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
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
	FCEMPL     string `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCEMPLTYPE string `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
	FCEMPLX1   string `gorm:"column:FCEMPLX1;" json:"fcemplx1"  form:"fcemplx1" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCJOB      string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCMTH      string `gorm:"column:FCMTH;" json:"fcmth"  form:"fcmth" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPROJ     string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCPRPAYTYP string `gorm:"column:FCPRPAYTYP;" json:"fcprpaytyp"  form:"fcprpaytyp" `
	FCPYPERIOD string `gorm:"column:FCPYPERIOD;" json:"fcpyperiod"  form:"fcpyperiod" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCYR       string `gorm:"column:FCYR;" json:"fcyr"  form:"fcyr" `
	FDBEGDATE  string `gorm:"column:FDBEGDATE;" json:"fdbegdate"  form:"fdbegdate" `
	FDENDDATE  string `gorm:"column:FDENDDATE;" json:"fdenddate"  form:"fdenddate" `
	FDPAYDATE  string `gorm:"column:FDPAYDATE;" json:"fdpaydate"  form:"fdpaydate" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMREMARK   string `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FNDAYINMTH string `gorm:"column:FNDAYINMTH;" json:"fndayinmth"  form:"fndayinmth" `
	FNDAYINYR  string `gorm:"column:FNDAYINYR;" json:"fndayinyr"  form:"fndayinyr" `
	FNFUND     string `gorm:"column:FNFUND;" json:"fnfund"  form:"fnfund" `
	FNFUNDC    string `gorm:"column:FNFUNDC;" json:"fnfundc"  form:"fnfundc" `
	FNHRINDAY  string `gorm:"column:FNHRINDAY;" json:"fnhrinday"  form:"fnhrinday" `
	FNINCBFFUN string `gorm:"column:FNINCBFFUN;" json:"fnincbffun"  form:"fnincbffun" `
	FNINCBFSOC string `gorm:"column:FNINCBFSOC;" json:"fnincbfsoc"  form:"fnincbfsoc" `
	FNINCBFTAX string `gorm:"column:FNINCBFTAX;" json:"fnincbftax"  form:"fnincbftax" `
	FNOTAMT    string `gorm:"column:FNOTAMT;" json:"fnotamt"  form:"fnotamt" `
	FNOTHDEDU  string `gorm:"column:FNOTHDEDU;" json:"fnothdedu"  form:"fnothdedu" `
	FNOTHINCO  string `gorm:"column:FNOTHINCO;" json:"fnothinco"  form:"fnothinco" `
	FNSALARY   string `gorm:"column:FNSALARY;" json:"fnsalary"  form:"fnsalary" `
	FNSAPERDAY string `gorm:"column:FNSAPERDAY;" json:"fnsaperday"  form:"fnsaperday" `
	FNSOCIAL   string `gorm:"column:FNSOCIAL;" json:"fnsocial"  form:"fnsocial" `
	FNTAX1     string `gorm:"column:FNTAX1;" json:"fntax1"  form:"fntax1" `
	FNTAX1_2   string `gorm:"column:FNTAX1_2;" json:"fntax1_2"  form:"fntax1_2" `
	FNU1CNT    string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNWORKDAY  string `gorm:"column:FNWORKDAY;" json:"fnworkday"  form:"fnworkday" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Hsalary) TableName() string {
	return "HSALARY"
}

func (obj *Hsalary) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
