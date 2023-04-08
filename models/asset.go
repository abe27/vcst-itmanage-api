package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Asset struct{
     FCACCDEPRA string `gorm:"column:FCACCDEPRA;" json:"fcaccdepra"  form:"fcaccdepra" `
     FCACCHART string `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
     FCACGROUP string `gorm:"column:FCACGROUP;" json:"fcacgroup"  form:"fcacgroup" `
     FCACTYPE string `gorm:"column:FCACTYPE;" json:"fcactype"  form:"fcactype" `
     FCADDAPVBY string `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCASCHAIN string `gorm:"column:FCASCHAIN;" json:"fcaschain"  form:"fcaschain" `
     FCASGRP string `gorm:"column:FCASGRP;" json:"fcasgrp"  form:"fcasgrp" `
     FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
     FCBARCODE string `gorm:"column:FCBARCODE;" json:"fcbarcode"  form:"fcbarcode" `
     FCBOICARD string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCBRAND string `gorm:"column:FCBRAND;" json:"fcbrand"  form:"fcbrand" `
     FCBUYFROM string `gorm:"column:FCBUYFROM;" json:"fcbuyfrom"  form:"fcbuyfrom" `
     FCCARID string `gorm:"column:FCCARID;" json:"fccarid"  form:"fccarid" `
     FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
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
     FCEXIST string `gorm:"column:FCEXIST;" json:"fcexist"  form:"fcexist" `
     FCFCHR string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
     FCGID string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
     FCINSTRUCT string `gorm:"column:FCINSTRUCT;" json:"fcinstruct"  form:"fcinstruct" `
     FCINSURECO string `gorm:"column:FCINSURECO;" json:"fcinsureco"  form:"fcinsureco" `
     FCINSURENO string `gorm:"column:FCINSURENO;" json:"fcinsureno"  form:"fcinsureno" `
     FCISCAL string `gorm:"column:FCISCAL;" json:"fciscal"  form:"fciscal" `
     FCISUSED string `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLID string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
     FCLOCA1 string `gorm:"column:FCLOCA1;" json:"fcloca1"  form:"fcloca1" `
     FCLOCATION string `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMAINASS string `gorm:"column:FCMAINASS;" json:"fcmainass"  form:"fcmainass" `
     FCMANFLAG string `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
     FCMETHOD string `gorm:"column:FCMETHOD;" json:"fcmethod"  form:"fcmethod" `
     FCMODEL string `gorm:"column:FCMODEL;" json:"fcmodel"  form:"fcmodel" `
     FCMTHDEPRA string `gorm:"column:FCMTHDEPRA;" json:"fcmthdepra"  form:"fcmthdepra" `
     FCNAME string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
     FCNAME2 string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
     FCOCCEMPL string `gorm:"column:FCOCCEMPL;" json:"fcoccempl"  form:"fcoccempl" `
     FCORDERNO string `gorm:"column:FCORDERNO;" json:"fcorderno"  form:"fcorderno" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPAYREFNO string `gorm:"column:FCPAYREFNO;" json:"fcpayrefno"  form:"fcpayrefno" `
     FCPLANT string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
     FCQCACCHAR string `gorm:"column:FCQCACCHAR;" json:"fcqcacchar"  form:"fcqcacchar" `
     FCREFPROD string `gorm:"column:FCREFPROD;" json:"fcrefprod"  form:"fcrefprod" `
     FCREMARK string `gorm:"column:FCREMARK;" json:"fcremark"  form:"fcremark" `
     FCREMARK2 string `gorm:"column:FCREMARK2;" json:"fcremark2"  form:"fcremark2" `
     FCRESTTYPE string `gorm:"column:FCRESTTYPE;" json:"fcresttype"  form:"fcresttype" `
     FCSALEREF string `gorm:"column:FCSALEREF;" json:"fcsaleref"  form:"fcsaleref" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSERIALNO string `gorm:"column:FCSERIALNO;" json:"fcserialno"  form:"fcserialno" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCSUBBR string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
     FCSUMOF string `gorm:"column:FCSUMOF;" json:"fcsumof"  form:"fcsumof" `
     FCTOUTON string `gorm:"column:FCTOUTON;" json:"fctouton"  form:"fctouton" `
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
     FCVCHOF string `gorm:"column:FCVCHOF;" json:"fcvchof"  form:"fcvchof" `
     FCWARRANNO string `gorm:"column:FCWARRANNO;" json:"fcwarranno"  form:"fcwarranno" `
     FDBUYDATE time.Time `gorm:"column:FDBUYDATE;" json:"fdbuydate"  form:"fdbuydate" default:"now"`
     FDENDCAL string `gorm:"column:FDENDCAL;" json:"fdendcal"  form:"fdendcal" `
     FDENDINSUR string `gorm:"column:FDENDINSUR;" json:"fdendinsur"  form:"fdendinsur" `
     FDSALEDATE string `gorm:"column:FDSALEDATE;" json:"fdsaledate"  form:"fdsaledate" `
     FDSTARTDAT time.Time `gorm:"column:FDSTARTDAT;" json:"fdstartdat"  form:"fdstartdat" default:"now"`
     FIMILLISEC int64 `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMERRMSG string `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMPICNAME string `gorm:"column:FMPICNAME;" json:"fmpicname"  form:"fmpicname" `
     FNACCMDEPR float64 `gorm:"column:FNACCMDEPR;" json:"fnaccmdepr"  form:"fnaccmdepr" `
     FNACCPRICE float64 `gorm:"column:FNACCPRICE;" json:"fnaccprice"  form:"fnaccprice" `
     FNBUYPRICE float64 `gorm:"column:FNBUYPRICE;" json:"fnbuyprice"  form:"fnbuyprice" `
     FNDEPRRATE float64 `gorm:"column:FNDEPRRATE;" json:"fndeprrate"  form:"fndeprrate" `
     FNDEPRREST float64 `gorm:"column:FNDEPRREST;" json:"fndeprrest"  form:"fndeprrest" `
     FNDPR0PCN float64 `gorm:"column:FNDPR0PCN;" json:"fndpr0pcn"  form:"fndpr0pcn" `
     FNGROSSPRI float64 `gorm:"column:FNGROSSPRI;" json:"fngrosspri"  form:"fngrosspri" `
     FNIMPLCOST float64 `gorm:"column:FNIMPLCOST;" json:"fnimplcost"  form:"fnimplcost" `
     FNINSUMTHP float64 `gorm:"column:FNINSUMTHP;" json:"fninsumthp"  form:"fninsumthp" `
     FNINSURELI float64 `gorm:"column:FNINSURELI;" json:"fninsureli"  form:"fninsureli" `
     FNLIFEYR float64 `gorm:"column:FNLIFEYR;" json:"fnlifeyr"  form:"fnlifeyr" `
     FNMISCCOST float64 `gorm:"column:FNMISCCOST;" json:"fnmisccost"  form:"fnmisccost" `
     FNMKTPRICE float64 `gorm:"column:FNMKTPRICE;" json:"fnmktprice"  form:"fnmktprice" `
     FNPREYRDEP float64 `gorm:"column:FNPREYRDEP;" json:"fnpreyrdep"  form:"fnpreyrdep" `
     FNPREYRVAL float64 `gorm:"column:FNPREYRVAL;" json:"fnpreyrval"  form:"fnpreyrval" `
     FNSALEPRIC float64 `gorm:"column:FNSALEPRIC;" json:"fnsalepric"  form:"fnsalepric" `
     FNSHIPCOST float64 `gorm:"column:FNSHIPCOST;" json:"fnshipcost"  form:"fnshipcost" `
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
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
     FTSRCUPD time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
}
func (Asset) TableName() string{
return "ASSET"
}

func (obj *Asset) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
