package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Saladj struct{
     FCACCHART string `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
     FCADDAPVBY string `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
     FCBASETYPE string `gorm:"column:FCBASETYPE;" json:"fcbasetype"  form:"fcbasetype" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCALTYPE string `gorm:"column:FCCALTYPE;" json:"fccaltype"  form:"fccaltype" `
     FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
     FCCODE2 string `gorm:"column:FCCODE2;" json:"fccode2"  form:"fccode2" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDELAPVBY string `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
     FCDETAXCON string `gorm:"column:FCDETAXCON;" json:"fcdetaxcon"  form:"fcdetaxcon" `
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
     FCEMPLTYPE string `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
     FCFCHR string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
     FCFRACTION string `gorm:"column:FCFRACTION;" json:"fcfraction"  form:"fcfraction" `
     FCGID string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
     FCINCRULE string `gorm:"column:FCINCRULE;" json:"fcincrule"  form:"fcincrule" `
     FCISCALFUN string `gorm:"column:FCISCALFUN;" json:"fciscalfun"  form:"fciscalfun" `
     FCISCALSOC string `gorm:"column:FCISCALSOC;" json:"fciscalsoc"  form:"fciscalsoc" `
     FCISCALTAX string `gorm:"column:FCISCALTAX;" json:"fciscaltax"  form:"fciscaltax" `
     FCISFIX string `gorm:"column:FCISFIX;" json:"fcisfix"  form:"fcisfix" `
     FCISUSED string `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
     FCLID string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
     FCMANFLAG string `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
     FCNAME string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
     FCNAME2 string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPAYTIMES string `gorm:"column:FCPAYTIMES;" json:"fcpaytimes"  form:"fcpaytimes" `
     FCRATEBY string `gorm:"column:FCRATEBY;" json:"fcrateby"  form:"fcrateby" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSHOW string `gorm:"column:FCSHOW;" json:"fcshow"  form:"fcshow" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSYSTYPE string `gorm:"column:FCSYSTYPE;" json:"fcsystype"  form:"fcsystype" `
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
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNRATE string `gorm:"column:FNRATE;" json:"fnrate"  form:"fnrate" `
     FNU1CNT string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
     FNU2CNT string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
     FNU3CNT string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
     FNU4CNT string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
     FNU5CNT string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
     FNU6CNT string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
     FNU7CNT string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
     FNU8CNT string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
     FNU9CNT string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
     FNWTAXPCN string `gorm:"column:FNWTAXPCN;" json:"fnwtaxpcn"  form:"fnwtaxpcn" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Saladj) TableName() string{
return "SALADJ"
}

func (obj *Saladj) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
