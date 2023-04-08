package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Layh struct{
     FCACSTEP string `gorm:"column:FCACSTEP;" json:"fcacstep"  form:"fcacstep" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBOOK string `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCANCELBY string `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
     FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
     FCCOOR string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCOUNTER string `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDISCTYPE string `gorm:"column:FCDISCTYPE;" json:"fcdisctype"  form:"fcdisctype" `
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
     FCEMPL string `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
     FCEMPLSM string `gorm:"column:FCEMPLSM;" json:"fcemplsm"  form:"fcemplsm" `
     FCGID string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
     FCGLHEAD string `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
     FCLID string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPDBRAND string `gorm:"column:FCPDBRAND;" json:"fcpdbrand"  form:"fcpdbrand" `
     FCRECALBY string `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
     FCREFNO string `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
     FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
     FCRFTYPE string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCSTEP string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
     FCTAXTYPE string `gorm:"column:FCTAXTYPE;" json:"fctaxtype"  form:"fctaxtype" `
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
     FDDATE time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
     FDDUEDATE time.Time `gorm:"column:FDDUEDATE;" json:"fdduedate"  form:"fdduedate" default:"now"`
     FDRECEDATE time.Time `gorm:"column:FDRECEDATE;" json:"fdrecedate"  form:"fdrecedate" default:"now"`
     FIMILLISEC int64 `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMDOCFLOW string `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMMEMDATA string `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
     FMMEMDATA2 string `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
     FMMEMDATA3 string `gorm:"column:FMMEMDATA3;" json:"fmmemdata3"  form:"fmmemdata3" `
     FMMEMDATA4 string `gorm:"column:FMMEMDATA4;" json:"fmmemdata4"  form:"fmmemdata4" `
     FMMEMDATA5 string `gorm:"column:FMMEMDATA5;" json:"fmmemdata5"  form:"fmmemdata5" `
     FNAMT float64 `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
     FNAMT2 float64 `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
     FNCREDTERM float64 `gorm:"column:FNCREDTERM;" json:"fncredterm"  form:"fncredterm" `
     FNU1CNT float64 `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
     FNU2CNT float64 `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
     FNU3CNT float64 `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
     FNU4CNT float64 `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
     FNU5CNT float64 `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
     FNU6CNT float64 `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
     FNU7CNT float64 `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
     FNU8CNT float64 `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
     FNU9CNT float64 `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
     FNVATAMT float64 `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
     FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
     FTLASRECAL string `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Layh) TableName() string{
return "LAYH"
}

func (obj *Layh) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
