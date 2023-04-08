package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Mnmgl struct{
     FCACCHART string `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
     FCACLINK string `gorm:"column:FCACLINK;" json:"fcaclink"  form:"fcaclink" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCATSTEP string `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
     FCBOICARD string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
     FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCLOSED string `gorm:"column:FCCLOSED;" json:"fcclosed"  form:"fcclosed" `
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
     FCGRP string `gorm:"column:FCGRP;" json:"fcgrp"  form:"fcgrp" `
     FCISINBROW string `gorm:"column:FCISINBROW;" json:"fcisinbrow"  form:"fcisinbrow" `
     FCISSUMOFB string `gorm:"column:FCISSUMOFB;" json:"fcissumofb"  form:"fcissumofb" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMNMGLH string `gorm:"column:FCMNMGLH;" json:"fcmnmglh"  form:"fcmnmglh" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPAYMENT string `gorm:"column:FCPAYMENT;" json:"fcpayment"  form:"fcpayment" `
     FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSUBBR string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FDDATE string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMDETAIL string `gorm:"column:FMDETAIL;" json:"fmdetail"  form:"fmdetail" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNAMT string `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
     FNCALLINE string `gorm:"column:FNCALLINE;" json:"fncalline"  form:"fncalline" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Mnmgl) TableName() string{
return "MNMGL"
}

func (obj *Mnmgl) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
