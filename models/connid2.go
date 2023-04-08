package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Connid2 struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCFCHR string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
     FCNAME string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}
func (Connid2) TableName() string{
return "CONNID2"
}

func (obj *Connid2) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}