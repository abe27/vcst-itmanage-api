package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Vmcstoph struct{
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMBOOK string `gorm:"column:FCMBOOK;" json:"fcmbook"  form:"fcmbook" `
     FCPLANT string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
     FCREFNO string `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
     FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
     FCRFTYPE string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
     FCROW string `gorm:"column:FCROW;" json:"fcrow"  form:"fcrow" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCSTEP string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
     FCTEAM string `gorm:"column:FCTEAM;" json:"fcteam"  form:"fcteam" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FDDATE string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMREMARK string `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
     FMREMARK10 string `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
     FMREMARK2 string `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
     FMREMARK3 string `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
     FMREMARK4 string `gorm:"column:FMREMARK4;" json:"fmremark4"  form:"fmremark4" `
     FMREMARK5 string `gorm:"column:FMREMARK5;" json:"fmremark5"  form:"fmremark5" `
     FMREMARK6 string `gorm:"column:FMREMARK6;" json:"fmremark6"  form:"fmremark6" `
     FMREMARK7 string `gorm:"column:FMREMARK7;" json:"fmremark7"  form:"fmremark7" `
     FMREMARK8 string `gorm:"column:FMREMARK8;" json:"fmremark8"  form:"fmremark8" `
     FMREMARK9 string `gorm:"column:FMREMARK9;" json:"fmremark9"  form:"fmremark9" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}
func (Vmcstoph) TableName() string{
return "VMCSTOPH"
}

func (obj *Vmcstoph) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
