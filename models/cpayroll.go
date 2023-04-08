package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Cpayroll struct{
     FCADDAPVBY string `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBANKBR string `gorm:"column:FCBANKBR;" json:"fcbankbr"  form:"fcbankbr" `
     FCBANKCORP string `gorm:"column:FCBANKCORP;" json:"fcbankcorp"  form:"fcbankcorp" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCALSADAY string `gorm:"column:FCCALSADAY;" json:"fccalsaday"  form:"fccalsaday" `
     FCCALSOC string `gorm:"column:FCCALSOC;" json:"fccalsoc"  form:"fccalsoc" `
     FCCALTMETH string `gorm:"column:FCCALTMETH;" json:"fccaltmeth"  form:"fccaltmeth" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDELAPVBY string `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
     FCDSTAXCON string `gorm:"column:FCDSTAXCON;" json:"fcdstaxcon"  form:"fcdstaxcon" `
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
     FCFRACFUN string `gorm:"column:FCFRACFUN;" json:"fcfracfun"  form:"fcfracfun" `
     FCFRACFUN2 string `gorm:"column:FCFRACFUN2;" json:"fcfracfun2"  form:"fcfracfun2" `
     FCFRACNET string `gorm:"column:FCFRACNET;" json:"fcfracnet"  form:"fcfracnet" `
     FCFRACSOC string `gorm:"column:FCFRACSOC;" json:"fcfracsoc"  form:"fcfracsoc" `
     FCFRACTAX string `gorm:"column:FCFRACTAX;" json:"fcfractax"  form:"fcfractax" `
     FCGID string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
     FCISUSED string `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
     FCLID string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
     FCMANFLAG string `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPAYNAME string `gorm:"column:FCPAYNAME;" json:"fcpayname"  form:"fcpayname" `
     FCPAYNAME2 string `gorm:"column:FCPAYNAME2;" json:"fcpayname2"  form:"fcpayname2" `
     FCPAYPOST string `gorm:"column:FCPAYPOST;" json:"fcpaypost"  form:"fcpaypost" `
     FCPAYPOST2 string `gorm:"column:FCPAYPOST2;" json:"fcpaypost2"  form:"fcpaypost2" `
     FCPOSTPKS1 string `gorm:"column:FCPOSTPKS1;" json:"fcpostpks1"  form:"fcpostpks1" `
     FCPOSTPKS2 string `gorm:"column:FCPOSTPKS2;" json:"fcpostpks2"  form:"fcpostpks2" `
     FCQCBANKAC string `gorm:"column:FCQCBANKAC;" json:"fcqcbankac"  form:"fcqcbankac" `
     FCSIGNPKS1 string `gorm:"column:FCSIGNPKS1;" json:"fcsignpks1"  form:"fcsignpks1" `
     FCSIGNPKS2 string `gorm:"column:FCSIGNPKS2;" json:"fcsignpks2"  form:"fcsignpks2" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSOCBRSEQ string `gorm:"column:FCSOCBRSEQ;" json:"fcsocbrseq"  form:"fcsocbrseq" `
     FCSOCNO string `gorm:"column:FCSOCNO;" json:"fcsocno"  form:"fcsocno" `
     FCSOCPVNAM string `gorm:"column:FCSOCPVNAM;" json:"fcsocpvnam"  form:"fcsocpvnam" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCTRADEREG string `gorm:"column:FCTRADEREG;" json:"fctradereg"  form:"fctradereg" `
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
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMMEMDATA string `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
     FNBEGTIME1 string `gorm:"column:FNBEGTIME1;" json:"fnbegtime1"  form:"fnbegtime1" `
     FNBEGTIME2 string `gorm:"column:FNBEGTIME2;" json:"fnbegtime2"  form:"fnbegtime2" `
     FNDAYINMTH string `gorm:"column:FNDAYINMTH;" json:"fndayinmth"  form:"fndayinmth" `
     FNDAYINYR string `gorm:"column:FNDAYINYR;" json:"fndayinyr"  form:"fndayinyr" `
     FNDAYPAY1 string `gorm:"column:FNDAYPAY1;" json:"fndaypay1"  form:"fndaypay1" `
     FNDAYPAY2 string `gorm:"column:FNDAYPAY2;" json:"fndaypay2"  form:"fndaypay2" `
     FNDAYPAY3 string `gorm:"column:FNDAYPAY3;" json:"fndaypay3"  form:"fndaypay3" `
     FNDAYPAY4 string `gorm:"column:FNDAYPAY4;" json:"fndaypay4"  form:"fndaypay4" `
     FNENDTIME1 string `gorm:"column:FNENDTIME1;" json:"fnendtime1"  form:"fnendtime1" `
     FNENDTIME2 string `gorm:"column:FNENDTIME2;" json:"fnendtime2"  form:"fnendtime2" `
     FNFCALLAST string `gorm:"column:FNFCALLAST;" json:"fnfcallast"  form:"fnfcallast" `
     FNHLYCAL string `gorm:"column:FNHLYCAL;" json:"fnhlycal"  form:"fnhlycal" `
     FNHOURPAY1 string `gorm:"column:FNHOURPAY1;" json:"fnhourpay1"  form:"fnhourpay1" `
     FNHOURPAY2 string `gorm:"column:FNHOURPAY2;" json:"fnhourpay2"  form:"fnhourpay2" `
     FNHOURPAY3 string `gorm:"column:FNHOURPAY3;" json:"fnhourpay3"  form:"fnhourpay3" `
     FNHOURPAY4 string `gorm:"column:FNHOURPAY4;" json:"fnhourpay4"  form:"fnhourpay4" `
     FNHRINDAY string `gorm:"column:FNHRINDAY;" json:"fnhrinday"  form:"fnhrinday" `
     FNMXBUSILV string `gorm:"column:FNMXBUSILV;" json:"fnmxbusilv"  form:"fnmxbusilv" `
     FNMXCASULV string `gorm:"column:FNMXCASULV;" json:"fnmxcasulv"  form:"fnmxcasulv" `
     FNMXSICKLV string `gorm:"column:FNMXSICKLV;" json:"fnmxsicklv"  form:"fnmxsicklv" `
     FNMXVACALV string `gorm:"column:FNMXVACALV;" json:"fnmxvacalv"  form:"fnmxvacalv" `
     FNPYPERIOD string `gorm:"column:FNPYPERIOD;" json:"fnpyperiod"  form:"fnpyperiod" `
     FNSALMANLV string `gorm:"column:FNSALMANLV;" json:"fnsalmanlv"  form:"fnsalmanlv" `
     FNSCALLAST string `gorm:"column:FNSCALLAST;" json:"fnscallast"  form:"fnscallast" `
     FNTAXDECI string `gorm:"column:FNTAXDECI;" json:"fntaxdeci"  form:"fntaxdeci" `
     FNU1CNT string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
     FNU2CNT string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
     FNU3CNT string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
     FNU4CNT string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
     FNU5CNT string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
     FNU6CNT string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
     FNU7CNT string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
     FNU8CNT string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
     FNU9CNT string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
     FNWCALLAST string `gorm:"column:FNWCALLAST;" json:"fnwcallast"  form:"fnwcallast" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Cpayroll) TableName() string{
return "CPAYROLL"
}

func (obj *Cpayroll) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
