package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Budget struct{
     FCACCHART string `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
     FCACGROUP string `gorm:"column:FCACGROUP;" json:"fcacgroup"  form:"fcacgroup" `
     FCACTYPE string `gorm:"column:FCACTYPE;" json:"fcactype"  form:"fcactype" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBOICARD string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
     FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCBUDGRP string `gorm:"column:FCBUDGRP;" json:"fcbudgrp"  form:"fcbudgrp" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
     FCDOCFLOWI string `gorm:"column:FCDOCFLOWI;" json:"fcdocflowi"  form:"fcdocflowi" `
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
     FCGID string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
     FCISCENTER string `gorm:"column:FCISCENTER;" json:"fciscenter"  form:"fciscenter" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLID string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSUBBR string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
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
     FCWARNING string `gorm:"column:FCWARNING;" json:"fcwarning"  form:"fcwarning" `
     FCYEAR string `gorm:"column:FCYEAR;" json:"fcyear"  form:"fcyear" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMDOCFLOW string `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNAMT01 string `gorm:"column:FNAMT01;" json:"fnamt01"  form:"fnamt01" `
     FNAMT02 string `gorm:"column:FNAMT02;" json:"fnamt02"  form:"fnamt02" `
     FNAMT03 string `gorm:"column:FNAMT03;" json:"fnamt03"  form:"fnamt03" `
     FNAMT04 string `gorm:"column:FNAMT04;" json:"fnamt04"  form:"fnamt04" `
     FNAMT05 string `gorm:"column:FNAMT05;" json:"fnamt05"  form:"fnamt05" `
     FNAMT06 string `gorm:"column:FNAMT06;" json:"fnamt06"  form:"fnamt06" `
     FNAMT07 string `gorm:"column:FNAMT07;" json:"fnamt07"  form:"fnamt07" `
     FNAMT08 string `gorm:"column:FNAMT08;" json:"fnamt08"  form:"fnamt08" `
     FNAMT09 string `gorm:"column:FNAMT09;" json:"fnamt09"  form:"fnamt09" `
     FNAMT10 string `gorm:"column:FNAMT10;" json:"fnamt10"  form:"fnamt10" `
     FNAMT11 string `gorm:"column:FNAMT11;" json:"fnamt11"  form:"fnamt11" `
     FNAMT12 string `gorm:"column:FNAMT12;" json:"fnamt12"  form:"fnamt12" `
     FNGLAMT string `gorm:"column:FNGLAMT;" json:"fnglamt"  form:"fnglamt" `
     FNINVAMT string `gorm:"column:FNINVAMT;" json:"fninvamt"  form:"fninvamt" `
     FNPCT01 string `gorm:"column:FNPCT01;" json:"fnpct01"  form:"fnpct01" `
     FNPCT02 string `gorm:"column:FNPCT02;" json:"fnpct02"  form:"fnpct02" `
     FNPCT03 string `gorm:"column:FNPCT03;" json:"fnpct03"  form:"fnpct03" `
     FNPCT04 string `gorm:"column:FNPCT04;" json:"fnpct04"  form:"fnpct04" `
     FNPCT05 string `gorm:"column:FNPCT05;" json:"fnpct05"  form:"fnpct05" `
     FNPCT06 string `gorm:"column:FNPCT06;" json:"fnpct06"  form:"fnpct06" `
     FNPCT07 string `gorm:"column:FNPCT07;" json:"fnpct07"  form:"fnpct07" `
     FNPCT08 string `gorm:"column:FNPCT08;" json:"fnpct08"  form:"fnpct08" `
     FNPCT09 string `gorm:"column:FNPCT09;" json:"fnpct09"  form:"fnpct09" `
     FNPCT10 string `gorm:"column:FNPCT10;" json:"fnpct10"  form:"fnpct10" `
     FNPCT11 string `gorm:"column:FNPCT11;" json:"fnpct11"  form:"fnpct11" `
     FNPCT12 string `gorm:"column:FNPCT12;" json:"fnpct12"  form:"fnpct12" `
     FNPOAMT string `gorm:"column:FNPOAMT;" json:"fnpoamt"  form:"fnpoamt" `
     FNPRAMT string `gorm:"column:FNPRAMT;" json:"fnpramt"  form:"fnpramt" `
     FNRESERVE string `gorm:"column:FNRESERVE;" json:"fnreserve"  form:"fnreserve" `
     FNTOTAL string `gorm:"column:FNTOTAL;" json:"fntotal"  form:"fntotal" `
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
func (Budget) TableName() string{
return "BUDGET"
}

func (obj *Budget) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
