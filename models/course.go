package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Course struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCERTIFIC string    `gorm:"column:FCCERTIFIC;" json:"fccertific"  form:"fccertific" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOURSEPN string    `gorm:"column:FCCOURSEPN;" json:"fccoursepn"  form:"fccoursepn" `
	FCCOURSETY string    `gorm:"column:FCCOURSETY;" json:"fccoursety"  form:"fccoursety" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEMPLFIX  string    `gorm:"column:FCEMPLFIX;" json:"fcemplfix"  form:"fcemplfix" `
	FCEMPLTYPE string    `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCINSTITUT string    `gorm:"column:FCINSTITUT;" json:"fcinstitut"  form:"fcinstitut" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string    `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPLACE    string    `gorm:"column:FCPLACE;" json:"fcplace"  form:"fcplace" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEMINABF string    `gorm:"column:FCSEMINABF;" json:"fcseminabf"  form:"fcseminabf" `
	FCSEX      string    `gorm:"column:FCSEX;" json:"fcsex"  form:"fcsex" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCSTUDYBF  string    `gorm:"column:FCSTUDYBF;" json:"fcstudybf"  form:"fcstudybf" `
	FCTRAINBF  string    `gorm:"column:FCTRAINBF;" json:"fctrainbf"  form:"fctrainbf" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
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
	FDINACTIVE string    `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMCOND     string    `gorm:"column:FMCOND;" json:"fmcond"  form:"fmcond" `
	FMOBJECTIV string    `gorm:"column:FMOBJECTIV;" json:"fmobjectiv"  form:"fmobjectiv" `
	FMPROPERTY string    `gorm:"column:FMPROPERTY;" json:"fmproperty"  form:"fmproperty" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FNBEGAGE   string    `gorm:"column:FNBEGAGE;" json:"fnbegage"  form:"fnbegage" `
	FNDAYS     string    `gorm:"column:FNDAYS;" json:"fndays"  form:"fndays" `
	FNENDAGE   string    `gorm:"column:FNENDAGE;" json:"fnendage"  form:"fnendage" `
	FNFOODEXP  string    `gorm:"column:FNFOODEXP;" json:"fnfoodexp"  form:"fnfoodexp" `
	FNHOTELEXP string    `gorm:"column:FNHOTELEXP;" json:"fnhotelexp"  form:"fnhotelexp" `
	FNOTHEREXP string    `gorm:"column:FNOTHEREXP;" json:"fnotherexp"  form:"fnotherexp" `
	FNPERSON   string    `gorm:"column:FNPERSON;" json:"fnperson"  form:"fnperson" `
	FNPOINT    string    `gorm:"column:FNPOINT;" json:"fnpoint"  form:"fnpoint" `
	FNTEACHEXP string    `gorm:"column:FNTEACHEXP;" json:"fnteachexp"  form:"fnteachexp" `
	FNTHEATEXP string    `gorm:"column:FNTHEATEXP;" json:"fntheatexp"  form:"fntheatexp" `
	FNTOOLEXP  string    `gorm:"column:FNTOOLEXP;" json:"fntoolexp"  form:"fntoolexp" `
	FNTOTALEXP string    `gorm:"column:FNTOTALEXP;" json:"fntotalexp"  form:"fntotalexp" `
	FNTRAINEXP string    `gorm:"column:FNTRAINEXP;" json:"fntrainexp"  form:"fntrainexp" `
	FNTRAVEXP  string    `gorm:"column:FNTRAVEXP;" json:"fntravexp"  form:"fntravexp" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNWORKAGE  string    `gorm:"column:FNWORKAGE;" json:"fnworkage"  form:"fnworkage" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Course) TableName() string {
	return "COURSE"
}

func (obj *Course) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
