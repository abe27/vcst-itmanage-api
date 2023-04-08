package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Repfile struct {
	FCACC1     string `gorm:"column:FCACC1;" json:"fcacc1"  form:"fcacc1" `
	FCACC2     string `gorm:"column:FCACC2;" json:"fcacc2"  form:"fcacc2" `
	FCACC3     string `gorm:"column:FCACC3;" json:"fcacc3"  form:"fcacc3" `
	FCACC4     string `gorm:"column:FCACC4;" json:"fcacc4"  form:"fcacc4" `
	FCADDAPVBY string `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCADJDESC  string `gorm:"column:FCADJDESC;" json:"fcadjdesc"  form:"fcadjdesc" `
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELAPVBY string `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCDESC1    string `gorm:"column:FCDESC1;" json:"fcdesc1"  form:"fcdesc1" `
	FCDESC2    string `gorm:"column:FCDESC2;" json:"fcdesc2"  form:"fcdesc2" `
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
	FCEDTAPVBY string `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCEND1     string `gorm:"column:FCEND1;" json:"fcend1"  form:"fcend1" `
	FCEND2     string `gorm:"column:FCEND2;" json:"fcend2"  form:"fcend2" `
	FCEND3     string `gorm:"column:FCEND3;" json:"fcend3"  form:"fcend3" `
	FCEND4     string `gorm:"column:FCEND4;" json:"fcend4"  form:"fcend4" `
	FCFONTSTYL string `gorm:"column:FCFONTSTYL;" json:"fcfontstyl"  form:"fcfontstyl" `
	FCFORM1    string `gorm:"column:FCFORM1;" json:"fcform1"  form:"fcform1" `
	FCFORM2    string `gorm:"column:FCFORM2;" json:"fcform2"  form:"fcform2" `
	FCFORM3    string `gorm:"column:FCFORM3;" json:"fcform3"  form:"fcform3" `
	FCFORM4    string `gorm:"column:FCFORM4;" json:"fcform4"  form:"fcform4" `
	FCFTYPE1   string `gorm:"column:FCFTYPE1;" json:"fcftype1"  form:"fcftype1" `
	FCFTYPE2   string `gorm:"column:FCFTYPE2;" json:"fcftype2"  form:"fcftype2" `
	FCFTYPE3   string `gorm:"column:FCFTYPE3;" json:"fcftype3"  form:"fcftype3" `
	FCFTYPE4   string `gorm:"column:FCFTYPE4;" json:"fcftype4"  form:"fcftype4" `
	FCGETVAR1  string `gorm:"column:FCGETVAR1;" json:"fcgetvar1"  form:"fcgetvar1" `
	FCGETVAR2  string `gorm:"column:FCGETVAR2;" json:"fcgetvar2"  form:"fcgetvar2" `
	FCGETVAR3  string `gorm:"column:FCGETVAR3;" json:"fcgetvar3"  form:"fcgetvar3" `
	FCGETVAR4  string `gorm:"column:FCGETVAR4;" json:"fcgetvar4"  form:"fcgetvar4" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCISDEFIV1 string `gorm:"column:FCISDEFIV1;" json:"fcisdefiv1"  form:"fcisdefiv1" `
	FCISDEFIV2 string `gorm:"column:FCISDEFIV2;" json:"fcisdefiv2"  form:"fcisdefiv2" `
	FCISDEFIV3 string `gorm:"column:FCISDEFIV3;" json:"fcisdefiv3"  form:"fcisdefiv3" `
	FCISDEFIV4 string `gorm:"column:FCISDEFIV4;" json:"fcisdefiv4"  form:"fcisdefiv4" `
	FCISNOPRN1 string `gorm:"column:FCISNOPRN1;" json:"fcisnoprn1"  form:"fcisnoprn1" `
	FCISNOPRN2 string `gorm:"column:FCISNOPRN2;" json:"fcisnoprn2"  form:"fcisnoprn2" `
	FCISNOPRN3 string `gorm:"column:FCISNOPRN3;" json:"fcisnoprn3"  form:"fcisnoprn3" `
	FCISNOPRN4 string `gorm:"column:FCISNOPRN4;" json:"fcisnoprn4"  form:"fcisnoprn4" `
	FCISQTY1   string `gorm:"column:FCISQTY1;" json:"fcisqty1"  form:"fcisqty1" `
	FCISQTY2   string `gorm:"column:FCISQTY2;" json:"fcisqty2"  form:"fcisqty2" `
	FCISQTY3   string `gorm:"column:FCISQTY3;" json:"fcisqty3"  form:"fcisqty3" `
	FCISQTY4   string `gorm:"column:FCISQTY4;" json:"fcisqty4"  form:"fcisqty4" `
	FCISUSED   string `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCLID      string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMANFLAG  string `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCMFORM    string `gorm:"column:FCMFORM;" json:"fcmform"  form:"fcmform" `
	FCOP1      string `gorm:"column:FCOP1;" json:"fcop1"  form:"fcop1" `
	FCOP2      string `gorm:"column:FCOP2;" json:"fcop2"  form:"fcop2" `
	FCOP3      string `gorm:"column:FCOP3;" json:"fcop3"  form:"fcop3" `
	FCOP4      string `gorm:"column:FCOP4;" json:"fcop4"  form:"fcop4" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPAGE     string `gorm:"column:FCPAGE;" json:"fcpage"  form:"fcpage" `
	FCQFORM1   string `gorm:"column:FCQFORM1;" json:"fcqform1"  form:"fcqform1" `
	FCQFORM2   string `gorm:"column:FCQFORM2;" json:"fcqform2"  form:"fcqform2" `
	FCQFORM3   string `gorm:"column:FCQFORM3;" json:"fcqform3"  form:"fcqform3" `
	FCQFORM4   string `gorm:"column:FCQFORM4;" json:"fcqform4"  form:"fcqform4" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCTIME1    string `gorm:"column:FCTIME1;" json:"fctime1"  form:"fctime1" `
	FCTIME2    string `gorm:"column:FCTIME2;" json:"fctime2"  form:"fctime2" `
	FCTIME3    string `gorm:"column:FCTIME3;" json:"fctime3"  form:"fctime3" `
	FCTIME4    string `gorm:"column:FCTIME4;" json:"fctime4"  form:"fctime4" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
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
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNLINE     string `gorm:"column:FNLINE;" json:"fnline"  form:"fnline" `
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

func (Repfile) TableName() string {
	return "REPFILE"
}

func (obj *Repfile) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
