package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Prodx8 struct {
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCCATEGORY string    `gorm:"column:FCCATEGORY;" json:"fccategory"  form:"fccategory" `
	FCCHKS     string    `gorm:"column:FCCHKS;" json:"fcchks"  form:"fcchks" `
	FCCLASS    string    `gorm:"column:FCCLASS;" json:"fcclass"  form:"fcclass" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOUNTRY  string    `gorm:"column:FCCOUNTRY;" json:"fccountry"  form:"fccountry" `
	FCCOUPON   string    `gorm:"column:FCCOUPON;" json:"fccoupon"  form:"fccoupon" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCDESC     string    `gorm:"column:FCDESC;" json:"fcdesc"  form:"fcdesc" `
	FCDESC2    string    `gorm:"column:FCDESC2;" json:"fcdesc2"  form:"fcdesc2" `
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
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGRADE    string    `gorm:"column:FCGRADE;" json:"fcgrade"  form:"fcgrade" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCMEDIA    string    `gorm:"column:FCMEDIA;" json:"fcmedia"  form:"fcmedia" `
	FCNBK      string    `gorm:"column:FCNBK;" json:"fcnbk"  form:"fcnbk" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCQCMV     string    `gorm:"column:FCQCMV;" json:"fcqcmv"  form:"fcqcmv" `
	FCQCSUPP   string    `gorm:"column:FCQCSUPP;" json:"fcqcsupp"  form:"fcqcsupp" `
	FCQCTITLE  string    `gorm:"column:FCQCTITLE;" json:"fcqctitle"  form:"fcqctitle" `
	FCQNMV     string    `gorm:"column:FCQNMV;" json:"fcqnmv"  form:"fcqnmv" `
	FCSDESC    string    `gorm:"column:FCSDESC;" json:"fcsdesc"  form:"fcsdesc" `
	FCSDESC2   string    `gorm:"column:FCSDESC2;" json:"fcsdesc2"  form:"fcsdesc2" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
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
	FCVERSION  string    `gorm:"column:FCVERSION;" json:"fcversion"  form:"fcversion" `
	FDBEGRIGHT string    `gorm:"column:FDBEGRIGHT;" json:"fdbegright"  form:"fdbegright" `
	FDDATEIN   string    `gorm:"column:FDDATEIN;" json:"fddatein"  form:"fddatein" `
	FDENDRIGHT string    `gorm:"column:FDENDRIGHT;" json:"fdendright"  form:"fdendright" `
	FDENDRT1   string    `gorm:"column:FDENDRT1;" json:"fdendrt1"  form:"fdendrt1" `
	FDENDRT2   string    `gorm:"column:FDENDRT2;" json:"fdendrt2"  form:"fdendrt2" `
	FDENDRT3   string    `gorm:"column:FDENDRT3;" json:"fdendrt3"  form:"fdendrt3" `
	FDUPDATE   string    `gorm:"column:FDUPDATE;" json:"fdupdate"  form:"fdupdate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNCVPRICE  string    `gorm:"column:FNCVPRICE;" json:"fncvprice"  form:"fncvprice" `
	FNMAXQTY   string    `gorm:"column:FNMAXQTY;" json:"fnmaxqty"  form:"fnmaxqty" `
	FNMBPRICE  string    `gorm:"column:FNMBPRICE;" json:"fnmbprice"  form:"fnmbprice" `
	FNMINQTY   string    `gorm:"column:FNMINQTY;" json:"fnminqty"  form:"fnminqty" `
	FNPOPRICE  string    `gorm:"column:FNPOPRICE;" json:"fnpoprice"  form:"fnpoprice" `
	FNRTDAY1   string    `gorm:"column:FNRTDAY1;" json:"fnrtday1"  form:"fnrtday1" `
	FNRTDAY2   string    `gorm:"column:FNRTDAY2;" json:"fnrtday2"  form:"fnrtday2" `
	FNRTDAY3   string    `gorm:"column:FNRTDAY3;" json:"fnrtday3"  form:"fnrtday3" `
	FNRTDAY4   string    `gorm:"column:FNRTDAY4;" json:"fnrtday4"  form:"fnrtday4" `
	FNRTPRICE1 string    `gorm:"column:FNRTPRICE1;" json:"fnrtprice1"  form:"fnrtprice1" `
	FNRTPRICE2 string    `gorm:"column:FNRTPRICE2;" json:"fnrtprice2"  form:"fnrtprice2" `
	FNRTPRICE3 string    `gorm:"column:FNRTPRICE3;" json:"fnrtprice3"  form:"fnrtprice3" `
	FNRTPRICE4 string    `gorm:"column:FNRTPRICE4;" json:"fnrtprice4"  form:"fnrtprice4" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Prodx8) TableName() string {
	return "PRODX8"
}

func (obj *Prodx8) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
