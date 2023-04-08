package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Trani struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCOOR string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCGAG string `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
     FCGVPOLICY string `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
     FCIOTYPE string `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
     FCLOT string `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMAINPROD string `gorm:"column:FCMAINPROD;" json:"fcmainprod"  form:"fcmainprod" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPLANT string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
     FCPROD string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
     FCREFH string `gorm:"column:FCREFH;" json:"fcrefh"  form:"fcrefh" `
     FCREFI string `gorm:"column:FCREFI;" json:"fcrefi"  form:"fcrefi" `
     FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSHIFT string `gorm:"column:FCSHIFT;" json:"fcshift"  form:"fcshift" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCSTEP string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
     FCSTORE string `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
     FCSUBTXID string `gorm:"column:FCSUBTXID;" json:"fcsubtxid"  form:"fcsubtxid" `
     FCTXID string `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FCWHOUSE string `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
     FDDATE string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNACTQTY string `gorm:"column:FNACTQTY;" json:"fnactqty"  form:"fnactqty" `
     FNALLOCQTY string `gorm:"column:FNALLOCQTY;" json:"fnallocqty"  form:"fnallocqty" `
     FNPLANQTY string `gorm:"column:FNPLANQTY;" json:"fnplanqty"  form:"fnplanqty" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Trani) TableName() string{
return "TRANI"
}

func (obj *Trani) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
