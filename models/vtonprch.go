package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vtonprch struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
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
	FCFCHR     string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCLID      string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCNAME     string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTATUS   string `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCTYPE     string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDBEGDATE  string `gorm:"column:FDBEGDATE;" json:"fdbegdate"  form:"fdbegdate" `
	FDENDDATE  string `gorm:"column:FDENDDATE;" json:"fdenddate"  form:"fdenddate" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNBEGAMT1  string `gorm:"column:FNBEGAMT1;" json:"fnbegamt1"  form:"fnbegamt1" `
	FNBEGAMT2  string `gorm:"column:FNBEGAMT2;" json:"fnbegamt2"  form:"fnbegamt2" `
	FNBEGAMT3  string `gorm:"column:FNBEGAMT3;" json:"fnbegamt3"  form:"fnbegamt3" `
	FNBEGAMT4  string `gorm:"column:FNBEGAMT4;" json:"fnbegamt4"  form:"fnbegamt4" `
	FNBEGAMT5  string `gorm:"column:FNBEGAMT5;" json:"fnbegamt5"  form:"fnbegamt5" `
	FNBEGAMT6  string `gorm:"column:FNBEGAMT6;" json:"fnbegamt6"  form:"fnbegamt6" `
	FNBEGAMT7  string `gorm:"column:FNBEGAMT7;" json:"fnbegamt7"  form:"fnbegamt7" `
	FNBEGAMT8  string `gorm:"column:FNBEGAMT8;" json:"fnbegamt8"  form:"fnbegamt8" `
	FNBEGAMT9  string `gorm:"column:FNBEGAMT9;" json:"fnbegamt9"  form:"fnbegamt9" `
	FNBEGAMTA  string `gorm:"column:FNBEGAMTA;" json:"fnbegamta"  form:"fnbegamta" `
	FNDISCPCN1 string `gorm:"column:FNDISCPCN1;" json:"fndiscpcn1"  form:"fndiscpcn1" `
	FNDISCPCN2 string `gorm:"column:FNDISCPCN2;" json:"fndiscpcn2"  form:"fndiscpcn2" `
	FNDISCPCN3 string `gorm:"column:FNDISCPCN3;" json:"fndiscpcn3"  form:"fndiscpcn3" `
	FNDISCPCN4 string `gorm:"column:FNDISCPCN4;" json:"fndiscpcn4"  form:"fndiscpcn4" `
	FNDISCPCN5 string `gorm:"column:FNDISCPCN5;" json:"fndiscpcn5"  form:"fndiscpcn5" `
	FNDISCPCN6 string `gorm:"column:FNDISCPCN6;" json:"fndiscpcn6"  form:"fndiscpcn6" `
	FNDISCPCN7 string `gorm:"column:FNDISCPCN7;" json:"fndiscpcn7"  form:"fndiscpcn7" `
	FNDISCPCN8 string `gorm:"column:FNDISCPCN8;" json:"fndiscpcn8"  form:"fndiscpcn8" `
	FNDISCPCN9 string `gorm:"column:FNDISCPCN9;" json:"fndiscpcn9"  form:"fndiscpcn9" `
	FNDISCPCNA string `gorm:"column:FNDISCPCNA;" json:"fndiscpcna"  form:"fndiscpcna" `
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
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vtonprch) TableName() string {
	return "VTONPRCH"
}

func (obj *Vtonprch) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
