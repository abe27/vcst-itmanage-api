package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Rubbin struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBOOK string `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCASE string `gorm:"column:FCCASE;" json:"fccase"  form:"fccase" `
     FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
     FCCOMPT string `gorm:"column:FCCOMPT;" json:"fccompt"  form:"fccompt" `
     FCCOOR string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCTRLSTOC string `gorm:"column:FCCTRLSTOC;" json:"fcctrlstoc"  form:"fcctrlstoc" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDATATYPE string `gorm:"column:FCDATATYPE;" json:"fcdatatype"  form:"fcdatatype" `
     FCDBFFILE string `gorm:"column:FCDBFFILE;" json:"fcdbffile"  form:"fcdbffile" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCEMPL string `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMBOOK string `gorm:"column:FCMBOOK;" json:"fcmbook"  form:"fcmbook" `
     FCOCTRLSTO string `gorm:"column:FCOCTRLSTO;" json:"fcoctrlsto"  form:"fcoctrlsto" `
     FCOLDCODE string `gorm:"column:FCOLDCODE;" json:"fcoldcode"  form:"fcoldcode" `
     FCOLDCOOR string `gorm:"column:FCOLDCOOR;" json:"fcoldcoor"  form:"fcoldcoor" `
     FCOLDEMPL string `gorm:"column:FCOLDEMPL;" json:"fcoldempl"  form:"fcoldempl" `
     FCOLDREFNO string `gorm:"column:FCOLDREFNO;" json:"fcoldrefno"  form:"fcoldrefno" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPLANT string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
     FCQCBOOK string `gorm:"column:FCQCBOOK;" json:"fcqcbook"  form:"fcqcbook" `
     FCQCMBOOK string `gorm:"column:FCQCMBOOK;" json:"fcqcmbook"  form:"fcqcmbook" `
     FCREFNO string `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCTABSKID string `gorm:"column:FCTABSKID;" json:"fctabskid"  form:"fctabskid" `
     FCTYPE string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FDDATE string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
     FIMILLISEC int64 `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNAMT float64 `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
     FNOLDAMT float64 `gorm:"column:FNOLDAMT;" json:"fnoldamt"  form:"fnoldamt" `
     FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
}
func (Rubbin) TableName() string{
return "RUBBIN"
}

func (obj *Rubbin) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
