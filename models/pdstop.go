package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Pdstop struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCRITICAL string `gorm:"column:FCCRITICAL;" json:"fccritical"  form:"fccritical" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDESC string `gorm:"column:FCDESC;" json:"fcdesc"  form:"fcdesc" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCISCOMMEN string `gorm:"column:FCISCOMMEN;" json:"fciscommen"  form:"fciscommen" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMOP string `gorm:"column:FCMOP;" json:"fcmop"  form:"fcmop" `
     FCOPSEQ string `gorm:"column:FCOPSEQ;" json:"fcopseq"  form:"fcopseq" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPDSTH string `gorm:"column:FCPDSTH;" json:"fcpdsth"  form:"fcpdsth" `
     FCPROCTYPE string `gorm:"column:FCPROCTYPE;" json:"fcproctype"  form:"fcproctype" `
     FCPTYPE string `gorm:"column:FCPTYPE;" json:"fcptype"  form:"fcptype" `
     FCQCMEMO string `gorm:"column:FCQCMEMO;" json:"fcqcmemo"  form:"fcqcmemo" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSUBOP string `gorm:"column:FCSUBOP;" json:"fcsubop"  form:"fcsubop" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FCWKCTRH string `gorm:"column:FCWKCTRH;" json:"fcwkctrh"  form:"fcwkctrh" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNBATCHSIZ string `gorm:"column:FNBATCHSIZ;" json:"fnbatchsiz"  form:"fnbatchsiz" `
     FNLT_MOVE string `gorm:"column:FNLT_MOVE;" json:"fnlt_move"  form:"fnlt_move" `
     FNLT_PROC string `gorm:"column:FNLT_PROC;" json:"fnlt_proc"  form:"fnlt_proc" `
     FNLT_QUE string `gorm:"column:FNLT_QUE;" json:"fnlt_que"  form:"fnlt_que" `
     FNLT_SETUP string `gorm:"column:FNLT_SETUP;" json:"fnlt_setup"  form:"fnlt_setup" `
     FNLT_TEAR string `gorm:"column:FNLT_TEAR;" json:"fnlt_tear"  form:"fnlt_tear" `
     FNLT_WAIT string `gorm:"column:FNLT_WAIT;" json:"fnlt_wait"  form:"fnlt_wait" `
     FNSTDTIME string `gorm:"column:FNSTDTIME;" json:"fnstdtime"  form:"fnstdtime" `
     FNWASTETIM string `gorm:"column:FNWASTETIM;" json:"fnwastetim"  form:"fnwastetim" `
     FNWIPCOST1 string `gorm:"column:FNWIPCOST1;" json:"fnwipcost1"  form:"fnwipcost1" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Pdstop) TableName() string{
return "PDSTOP"
}

func (obj *Pdstop) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
