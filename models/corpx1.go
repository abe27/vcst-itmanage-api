package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Corpx1 struct{
     FCADDAPVBY string `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
     FCCON000 string `gorm:"column:FCCON000;" json:"fccon000"  form:"fccon000" `
     FCCON001 string `gorm:"column:FCCON001;" json:"fccon001"  form:"fccon001" `
     FCCON002 string `gorm:"column:FCCON002;" json:"fccon002"  form:"fccon002" `
     FCCON003 string `gorm:"column:FCCON003;" json:"fccon003"  form:"fccon003" `
     FCCON004 string `gorm:"column:FCCON004;" json:"fccon004"  form:"fccon004" `
     FCCON005 string `gorm:"column:FCCON005;" json:"fccon005"  form:"fccon005" `
     FCCON006 string `gorm:"column:FCCON006;" json:"fccon006"  form:"fccon006" `
     FCCON007 string `gorm:"column:FCCON007;" json:"fccon007"  form:"fccon007" `
     FCCON008 string `gorm:"column:FCCON008;" json:"fccon008"  form:"fccon008" `
     FCCON009 string `gorm:"column:FCCON009;" json:"fccon009"  form:"fccon009" `
     FCCON010 string `gorm:"column:FCCON010;" json:"fccon010"  form:"fccon010" `
     FCCON011 string `gorm:"column:FCCON011;" json:"fccon011"  form:"fccon011" `
     FCCON012 string `gorm:"column:FCCON012;" json:"fccon012"  form:"fccon012" `
     FCCON013 string `gorm:"column:FCCON013;" json:"fccon013"  form:"fccon013" `
     FCCON014 string `gorm:"column:FCCON014;" json:"fccon014"  form:"fccon014" `
     FCCON015 string `gorm:"column:FCCON015;" json:"fccon015"  form:"fccon015" `
     FCCON016 string `gorm:"column:FCCON016;" json:"fccon016"  form:"fccon016" `
     FCCON017 string `gorm:"column:FCCON017;" json:"fccon017"  form:"fccon017" `
     FCCON018 string `gorm:"column:FCCON018;" json:"fccon018"  form:"fccon018" `
     FCCON019 string `gorm:"column:FCCON019;" json:"fccon019"  form:"fccon019" `
     FCCON020 string `gorm:"column:FCCON020;" json:"fccon020"  form:"fccon020" `
     FCCON021 string `gorm:"column:FCCON021;" json:"fccon021"  form:"fccon021" `
     FCCON022 string `gorm:"column:FCCON022;" json:"fccon022"  form:"fccon022" `
     FCCON023 string `gorm:"column:FCCON023;" json:"fccon023"  form:"fccon023" `
     FCCON024 string `gorm:"column:FCCON024;" json:"fccon024"  form:"fccon024" `
     FCCON025 string `gorm:"column:FCCON025;" json:"fccon025"  form:"fccon025" `
     FCCON026 string `gorm:"column:FCCON026;" json:"fccon026"  form:"fccon026" `
     FCCON027 string `gorm:"column:FCCON027;" json:"fccon027"  form:"fccon027" `
     FCCON028 string `gorm:"column:FCCON028;" json:"fccon028"  form:"fccon028" `
     FCCON029 string `gorm:"column:FCCON029;" json:"fccon029"  form:"fccon029" `
     FCCON030 string `gorm:"column:FCCON030;" json:"fccon030"  form:"fccon030" `
     FCCON100 string `gorm:"column:FCCON100;" json:"fccon100"  form:"fccon100" `
     FCCON101 string `gorm:"column:FCCON101;" json:"fccon101"  form:"fccon101" `
     FCCON102 string `gorm:"column:FCCON102;" json:"fccon102"  form:"fccon102" `
     FCCON103 string `gorm:"column:FCCON103;" json:"fccon103"  form:"fccon103" `
     FCCON104 string `gorm:"column:FCCON104;" json:"fccon104"  form:"fccon104" `
     FCCON105 string `gorm:"column:FCCON105;" json:"fccon105"  form:"fccon105" `
     FCCON106 string `gorm:"column:FCCON106;" json:"fccon106"  form:"fccon106" `
     FCCON107 string `gorm:"column:FCCON107;" json:"fccon107"  form:"fccon107" `
     FCCON108 string `gorm:"column:FCCON108;" json:"fccon108"  form:"fccon108" `
     FCCON109 string `gorm:"column:FCCON109;" json:"fccon109"  form:"fccon109" `
     FCCON110 string `gorm:"column:FCCON110;" json:"fccon110"  form:"fccon110" `
     FCCON111 string `gorm:"column:FCCON111;" json:"fccon111"  form:"fccon111" `
     FCCON112 string `gorm:"column:FCCON112;" json:"fccon112"  form:"fccon112" `
     FCCON113 string `gorm:"column:FCCON113;" json:"fccon113"  form:"fccon113" `
     FCCON114 string `gorm:"column:FCCON114;" json:"fccon114"  form:"fccon114" `
     FCCON115 string `gorm:"column:FCCON115;" json:"fccon115"  form:"fccon115" `
     FCCON116 string `gorm:"column:FCCON116;" json:"fccon116"  form:"fccon116" `
     FCCON117 string `gorm:"column:FCCON117;" json:"fccon117"  form:"fccon117" `
     FCCON118 string `gorm:"column:FCCON118;" json:"fccon118"  form:"fccon118" `
     FCCON119 string `gorm:"column:FCCON119;" json:"fccon119"  form:"fccon119" `
     FCCON120 string `gorm:"column:FCCON120;" json:"fccon120"  form:"fccon120" `
     FCCON121 string `gorm:"column:FCCON121;" json:"fccon121"  form:"fccon121" `
     FCCON122 string `gorm:"column:FCCON122;" json:"fccon122"  form:"fccon122" `
     FCCON123 string `gorm:"column:FCCON123;" json:"fccon123"  form:"fccon123" `
     FCCON124 string `gorm:"column:FCCON124;" json:"fccon124"  form:"fccon124" `
     FCCON125 string `gorm:"column:FCCON125;" json:"fccon125"  form:"fccon125" `
     FCCON126 string `gorm:"column:FCCON126;" json:"fccon126"  form:"fccon126" `
     FCCON127 string `gorm:"column:FCCON127;" json:"fccon127"  form:"fccon127" `
     FCCON128 string `gorm:"column:FCCON128;" json:"fccon128"  form:"fccon128" `
     FCCON129 string `gorm:"column:FCCON129;" json:"fccon129"  form:"fccon129" `
     FCCON130 string `gorm:"column:FCCON130;" json:"fccon130"  form:"fccon130" `
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
     FCENCOTYPE string `gorm:"column:FCENCOTYPE;" json:"fcencotype"  form:"fcencotype" `
     FCGID string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
     FCISUSED string `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
     FCLID string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMANFLAG string `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
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
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMWEBDATA string `gorm:"column:FMWEBDATA;" json:"fmwebdata"  form:"fmwebdata" `
     FNU1CNT string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
     FNU2CNT string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
     FNU3CNT string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
     FNU4CNT string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
     FNU5CNT string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
     FNU6CNT string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
     FNU7CNT string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
     FNU8CNT string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
     FNU9CNT string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Corpx1) TableName() string{
return "CORPX1"
}

func (obj *Corpx1) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
