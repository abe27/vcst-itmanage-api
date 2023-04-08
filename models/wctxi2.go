package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Wctxi2 struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBOICARD string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
     FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCEMPL2RCV string `gorm:"column:FCEMPL2RCV;" json:"fcempl2rcv"  form:"fcempl2rcv" `
     FCFRWKCTRH string `gorm:"column:FCFRWKCTRH;" json:"fcfrwkctrh"  form:"fcfrwkctrh" `
     FCGAG string `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
     FCIOTYPE string `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLOT string `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMACHINE string `gorm:"column:FCMACHINE;" json:"fcmachine"  form:"fcmachine" `
     FCMOP string `gorm:"column:FCMOP;" json:"fcmop"  form:"fcmop" `
     FCMORDERH string `gorm:"column:FCMORDERH;" json:"fcmorderh"  form:"fcmorderh" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPLANT string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
     FCPROD string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
     FCPRODRM string `gorm:"column:FCPRODRM;" json:"fcprodrm"  form:"fcprodrm" `
     FCPRODTYPE string `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
     FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
     FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
     FCRFTYPE string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSHIFT string `gorm:"column:FCSHIFT;" json:"fcshift"  form:"fcshift" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCSTEP string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
     FCSUBBR string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
     FCTIMESTAM string `gorm:"column:FCTIMESTAM;" json:"fctimestam"  form:"fctimestam" `
     FCTOWKCTRH string `gorm:"column:FCTOWKCTRH;" json:"fctowkctrh"  form:"fctowkctrh" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUM string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
     FCUMSTD string `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FCWCTXH2 string `gorm:"column:FCWCTXH2;" json:"fcwctxh2"  form:"fcwctxh2" `
     FCWKCTRH string `gorm:"column:FCWKCTRH;" json:"fcwkctrh"  form:"fcwkctrh" `
     FDDATE string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMREMARK string `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
     FNGOODQTY string `gorm:"column:FNGOODQTY;" json:"fngoodqty"  form:"fngoodqty" `
     FNHOLDQTY string `gorm:"column:FNHOLDQTY;" json:"fnholdqty"  form:"fnholdqty" `
     FNUMQTY string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
     FNWASTEQTY string `gorm:"column:FNWASTEQTY;" json:"fnwasteqty"  form:"fnwasteqty" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTFINISH string `gorm:"column:FTFINISH;" json:"ftfinish"  form:"ftfinish" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
     FTSTART string `gorm:"column:FTSTART;" json:"ftstart"  form:"ftstart" `
}
func (Wctxi2) TableName() string{
return "WCTXI2"
}

func (obj *Wctxi2) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
