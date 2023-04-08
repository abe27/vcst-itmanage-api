package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Mbappv struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAID string `gorm:"column:FCDATAID;" json:"fcdataid"  form:"fcdataid" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCEMPL string `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTATUS string `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
     FCTABNAME string `gorm:"column:FCTABNAME;" json:"fctabname"  form:"fctabname" `
     FCTABSKID string `gorm:"column:FCTABSKID;" json:"fctabskid"  form:"fctabskid" `
     FCTSTATUS string `gorm:"column:FCTSTATUS;" json:"fctstatus"  form:"fctstatus" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMREMARK string `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
     FTAPV string `gorm:"column:FTAPV;" json:"ftapv"  form:"ftapv" `
     FTCANCEL string `gorm:"column:FTCANCEL;" json:"ftcancel"  form:"ftcancel" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTDISAPV string `gorm:"column:FTDISAPV;" json:"ftdisapv"  form:"ftdisapv" `
     FTHOLD string `gorm:"column:FTHOLD;" json:"fthold"  form:"fthold" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTREAD string `gorm:"column:FTREAD;" json:"ftread"  form:"ftread" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
     FTUNAPV string `gorm:"column:FTUNAPV;" json:"ftunapv"  form:"ftunapv" `
}
func (Mbappv) TableName() string{
return "MBAPPV"
}

func (obj *Mbappv) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
