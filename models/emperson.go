package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Emperson struct {
	FCADD11    string `gorm:"column:FCADD11;" json:"fcadd11"  form:"fcadd11" `
	FCADD12    string `gorm:"column:FCADD12;" json:"fcadd12"  form:"fcadd12" `
	FCADD21    string `gorm:"column:FCADD21;" json:"fcadd21"  form:"fcadd21" `
	FCADD22    string `gorm:"column:FCADD22;" json:"fcadd22"  form:"fcadd22" `
	FCADD31    string `gorm:"column:FCADD31;" json:"fcadd31"  form:"fcadd31" `
	FCADD32    string `gorm:"column:FCADD32;" json:"fcadd32"  form:"fcadd32" `
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBEGYEAR  string `gorm:"column:FCBEGYEAR;" json:"fcbegyear"  form:"fcbegyear" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
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
	FCEDULEV   string `gorm:"column:FCEDULEV;" json:"fcedulev"  form:"fcedulev" `
	FCEDUMAJOR string `gorm:"column:FCEDUMAJOR;" json:"fcedumajor"  form:"fcedumajor" `
	FCEMAIL    string `gorm:"column:FCEMAIL;" json:"fcemail"  form:"fcemail" `
	FCEMPLX1   string `gorm:"column:FCEMPLX1;" json:"fcemplx1"  form:"fcemplx1" `
	FCFCHR     string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCID       string `gorm:"column:FCID;" json:"fcid"  form:"fcid" `
	FCIDDIST   string `gorm:"column:FCIDDIST;" json:"fciddist"  form:"fciddist" `
	FCIDPROV   string `gorm:"column:FCIDPROV;" json:"fcidprov"  form:"fcidprov" `
	FCISEMADDR string `gorm:"column:FCISEMADDR;" json:"fcisemaddr"  form:"fcisemaddr" `
	FCISLIVE   string `gorm:"column:FCISLIVE;" json:"fcislive"  form:"fcislive" `
	FCISSTUDY  string `gorm:"column:FCISSTUDY;" json:"fcisstudy"  form:"fcisstudy" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMOBILE   string `gorm:"column:FCMOBILE;" json:"fcmobile"  form:"fcmobile" `
	FCNAME     string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNATION   string `gorm:"column:FCNATION;" json:"fcnation"  form:"fcnation" `
	FCOCCUP    string `gorm:"column:FCOCCUP;" json:"fcoccup"  form:"fcoccup" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPOST     string `gorm:"column:FCPOST;" json:"fcpost"  form:"fcpost" `
	FCPRENAME  string `gorm:"column:FCPRENAME;" json:"fcprename"  form:"fcprename" `
	FCRACE     string `gorm:"column:FCRACE;" json:"fcrace"  form:"fcrace" `
	FCRELATION string `gorm:"column:FCRELATION;" json:"fcrelation"  form:"fcrelation" `
	FCSCHOOL   string `gorm:"column:FCSCHOOL;" json:"fcschool"  form:"fcschool" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEX      string `gorm:"column:FCSEX;" json:"fcsex"  form:"fcsex" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCTEL22    string `gorm:"column:FCTEL22;" json:"fctel22"  form:"fctel22" `
	FCTEL32    string `gorm:"column:FCTEL32;" json:"fctel32"  form:"fctel32" `
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
	FDBIRTH    string `gorm:"column:FDBIRTH;" json:"fdbirth"  form:"fdbirth" `
	FDIDDATE   string `gorm:"column:FDIDDATE;" json:"fdiddate"  form:"fdiddate" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMPICADDCR string `gorm:"column:FMPICADDCR;" json:"fmpicaddcr"  form:"fmpicaddcr" `
	FMPICADDWK string `gorm:"column:FMPICADDWK;" json:"fmpicaddwk"  form:"fmpicaddwk" `
	FNAGE      string `gorm:"column:FNAGE;" json:"fnage"  form:"fnage" `
	FNSALARY   string `gorm:"column:FNSALARY;" json:"fnsalary"  form:"fnsalary" `
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

func (Emperson) TableName() string {
	return "EMPERSON"
}

func (obj *Emperson) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
