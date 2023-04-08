package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Bankacct struct{
     FCACCBOOK string `gorm:"column:FCACCBOOK;" json:"fcaccbook"  form:"fcaccbook" `
     FCACCBOOK2 string `gorm:"column:FCACCBOOK2;" json:"fcaccbook2"  form:"fcaccbook2" `
     FCACCBOOK3 string `gorm:"column:FCACCBOOK3;" json:"fcaccbook3"  form:"fcaccbook3" `
     FCACCBOOK4 string `gorm:"column:FCACCBOOK4;" json:"fcaccbook4"  form:"fcaccbook4" `
     FCACCBOOK5 string `gorm:"column:FCACCBOOK5;" json:"fcaccbook5"  form:"fcaccbook5" `
     FCACCHART string `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
     FCACC_CD string `gorm:"column:FCACC_CD;" json:"fcacc_cd"  form:"fcacc_cd" `
     FCACC_CW string `gorm:"column:FCACC_CW;" json:"fcacc_cw"  form:"fcacc_cw" `
     FCACC_HC string `gorm:"column:FCACC_HC;" json:"fcacc_hc"  form:"fcacc_hc" `
     FCADDAPVBY string `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
     FCADDR1 string `gorm:"column:FCADDR1;" json:"fcaddr1"  form:"fcaddr1" `
     FCADDR12 string `gorm:"column:FCADDR12;" json:"fcaddr12"  form:"fcaddr12" `
     FCADDR2 string `gorm:"column:FCADDR2;" json:"fcaddr2"  form:"fcaddr2" `
     FCADDR22 string `gorm:"column:FCADDR22;" json:"fcaddr22"  form:"fcaddr22" `
     FCADDR3 string `gorm:"column:FCADDR3;" json:"fcaddr3"  form:"fcaddr3" `
     FCADDR32 string `gorm:"column:FCADDR32;" json:"fcaddr32"  form:"fcaddr32" `
     FCADVICEBY string `gorm:"column:FCADVICEBY;" json:"fcadviceby"  form:"fcadviceby" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
     FCBANK string `gorm:"column:FCBANK;" json:"fcbank"  form:"fcbank" `
     FCBBRANCH string `gorm:"column:FCBBRANCH;" json:"fcbbranch"  form:"fcbbranch" `
     FCBICCODE string `gorm:"column:FCBICCODE;" json:"fcbiccode"  form:"fcbiccode" `
     FCBKBRCODE string `gorm:"column:FCBKBRCODE;" json:"fcbkbrcode"  form:"fcbkbrcode" `
     FCBKCORPID string `gorm:"column:FCBKCORPID;" json:"fcbkcorpid"  form:"fcbkcorpid" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCHGBEAR string `gorm:"column:FCCHGBEAR;" json:"fcchgbear"  form:"fcchgbear" `
     FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDCTID string `gorm:"column:FCDCTID;" json:"fcdctid"  form:"fcdctid" `
     FCDELAPVBY string `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
     FCDTYPE1 string `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
     FCDTYPE2 string `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
     FCDTYPE3 string `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
     FCDTYPE4 string `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
     FCDTYPE5 string `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
     FCDTYPE6 string `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
     FCDTYPE7 string `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
     FCDTYPE8 string `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
     FCDTYPE9 string `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCEDTAPVBY string `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
     FCFCHR string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
     FCFORMH string `gorm:"column:FCFORMH;" json:"fcformh"  form:"fcformh" `
     FCFORMH2 string `gorm:"column:FCFORMH2;" json:"fcformh2"  form:"fcformh2" `
     FCFORMH3 string `gorm:"column:FCFORMH3;" json:"fcformh3"  form:"fcformh3" `
     FCFORMH4 string `gorm:"column:FCFORMH4;" json:"fcformh4"  form:"fcformh4" `
     FCFORMHC2 string `gorm:"column:FCFORMHC2;" json:"fcformhc2"  form:"fcformhc2" `
     FCFORMHC3 string `gorm:"column:FCFORMHC3;" json:"fcformhc3"  form:"fcformhc3" `
     FCFORMHC4 string `gorm:"column:FCFORMHC4;" json:"fcformhc4"  form:"fcformhc4" `
     FCFORMHC5 string `gorm:"column:FCFORMHC5;" json:"fcformhc5"  form:"fcformhc5" `
     FCGID string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
     FCIBPID string `gorm:"column:FCIBPID;" json:"fcibpid"  form:"fcibpid" `
     FCISUSED string `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
     FCLID string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMANFLAG string `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
     FCNAME string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
     FCNAME2 string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
     FCNCHRGBR string `gorm:"column:FCNCHRGBR;" json:"fcnchrgbr"  form:"fcnchrgbr" `
     FCNO string `gorm:"column:FCNO;" json:"fcno"  form:"fcno" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCOVRACCHA string `gorm:"column:FCOVRACCHA;" json:"fcovraccha"  form:"fcovraccha" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSERVTYPE string `gorm:"column:FCSERVTYPE;" json:"fcservtype"  form:"fcservtype" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCTYPE string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
     FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
     FCU3STATUS string `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
     FCU4STATUS string `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
     FCU5STATUS string `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
     FCU6STATUS string `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
     FCU7STATUS string `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
     FCU8STATUS string `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
     FCU9STATUS string `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FDBEGDATE time.Time `gorm:"column:FDBEGDATE;" json:"fdbegdate"  form:"fdbegdate" default:"now"`
     FDSTARTDAT time.Time `gorm:"column:FDSTARTDAT;" json:"fdstartdat"  form:"fdstartdat" default:"now"`
     FIMILLISEC int64 `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMERRMSG string `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMMEMDATA string `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
     FNBANKGAMT float64 `gorm:"column:FNBANKGAMT;" json:"fnbankgamt"  form:"fnbankgamt" `
     FNBEGBALAN float64 `gorm:"column:FNBEGBALAN;" json:"fnbegbalan"  form:"fnbegbalan" `
     FNODAMT float64 `gorm:"column:FNODAMT;" json:"fnodamt"  form:"fnodamt" `
     FNU1CNT float64 `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
     FNU2CNT float64 `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
     FNU3CNT float64 `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
     FNU4CNT float64 `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
     FNU5CNT float64 `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
     FNU6CNT float64 `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
     FNU7CNT float64 `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
     FNU8CNT float64 `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
     FNU9CNT float64 `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
     FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
     FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
     FTLASTUPD time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
     FTSRCUPD time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
}
func (Bankacct) TableName() string{
return "BANKACCT"
}

func (obj *Bankacct) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
