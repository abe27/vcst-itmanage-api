package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Stdhour struct {
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBEGTIME1 string    `gorm:"column:FCBEGTIME1;" json:"fcbegtime1"  form:"fcbegtime1" `
	FCBEGTIME2 string    `gorm:"column:FCBEGTIME2;" json:"fcbegtime2"  form:"fcbegtime2" `
	FCBEGTIME3 string    `gorm:"column:FCBEGTIME3;" json:"fcbegtime3"  form:"fcbegtime3" `
	FCBEGTIME4 string    `gorm:"column:FCBEGTIME4;" json:"fcbegtime4"  form:"fcbegtime4" `
	FCBEGTIME5 string    `gorm:"column:FCBEGTIME5;" json:"fcbegtime5"  form:"fcbegtime5" `
	FCBEGTIME6 string    `gorm:"column:FCBEGTIME6;" json:"fcbegtime6"  form:"fcbegtime6" `
	FCBEGTIME7 string    `gorm:"column:FCBEGTIME7;" json:"fcbegtime7"  form:"fcbegtime7" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
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
	FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCENDTIME1 string    `gorm:"column:FCENDTIME1;" json:"fcendtime1"  form:"fcendtime1" `
	FCENDTIME2 string    `gorm:"column:FCENDTIME2;" json:"fcendtime2"  form:"fcendtime2" `
	FCENDTIME3 string    `gorm:"column:FCENDTIME3;" json:"fcendtime3"  form:"fcendtime3" `
	FCENDTIME4 string    `gorm:"column:FCENDTIME4;" json:"fcendtime4"  form:"fcendtime4" `
	FCENDTIME5 string    `gorm:"column:FCENDTIME5;" json:"fcendtime5"  form:"fcendtime5" `
	FCENDTIME6 string    `gorm:"column:FCENDTIME6;" json:"fcendtime6"  form:"fcendtime6" `
	FCENDTIME7 string    `gorm:"column:FCENDTIME7;" json:"fcendtime7"  form:"fcendtime7" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGMACHINE string    `gorm:"column:FCGMACHINE;" json:"fcgmachine"  form:"fcgmachine" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMACHINE  string    `gorm:"column:FCMACHINE;" json:"fcmachine"  form:"fcmachine" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSHIFT    string    `gorm:"column:FCSHIFT;" json:"fcshift"  form:"fcshift" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
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
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCWKCTRH   string    `gorm:"column:FCWKCTRH;" json:"fcwkctrh"  form:"fcwkctrh" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNOTHOUR1  string    `gorm:"column:FNOTHOUR1;" json:"fnothour1"  form:"fnothour1" `
	FNOTHOUR2  string    `gorm:"column:FNOTHOUR2;" json:"fnothour2"  form:"fnothour2" `
	FNOTHOUR3  string    `gorm:"column:FNOTHOUR3;" json:"fnothour3"  form:"fnothour3" `
	FNOTHOUR4  string    `gorm:"column:FNOTHOUR4;" json:"fnothour4"  form:"fnothour4" `
	FNOTHOUR5  string    `gorm:"column:FNOTHOUR5;" json:"fnothour5"  form:"fnothour5" `
	FNOTHOUR6  string    `gorm:"column:FNOTHOUR6;" json:"fnothour6"  form:"fnothour6" `
	FNOTHOUR7  string    `gorm:"column:FNOTHOUR7;" json:"fnothour7"  form:"fnothour7" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNWKHOUR1  string    `gorm:"column:FNWKHOUR1;" json:"fnwkhour1"  form:"fnwkhour1" `
	FNWKHOUR2  string    `gorm:"column:FNWKHOUR2;" json:"fnwkhour2"  form:"fnwkhour2" `
	FNWKHOUR3  string    `gorm:"column:FNWKHOUR3;" json:"fnwkhour3"  form:"fnwkhour3" `
	FNWKHOUR4  string    `gorm:"column:FNWKHOUR4;" json:"fnwkhour4"  form:"fnwkhour4" `
	FNWKHOUR5  string    `gorm:"column:FNWKHOUR5;" json:"fnwkhour5"  form:"fnwkhour5" `
	FNWKHOUR6  string    `gorm:"column:FNWKHOUR6;" json:"fnwkhour6"  form:"fnwkhour6" `
	FNWKHOUR7  string    `gorm:"column:FNWKHOUR7;" json:"fnwkhour7"  form:"fnwkhour7" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Stdhour) TableName() string {
	return "STDHOUR"
}

func (obj *Stdhour) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
