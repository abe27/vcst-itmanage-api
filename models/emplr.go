package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Emplr struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCCAPMODE string `gorm:"column:FCCAPMODE;" json:"fccapmode"  form:"fccapmode" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCENCOTYPE string `gorm:"column:FCENCOTYPE;" json:"fcencotype"  form:"fcencotype" `
     FCFIELLANG string `gorm:"column:FCFIELLANG;" json:"fcfiellang"  form:"fcfiellang" `
     FCHHPW string `gorm:"column:FCHHPW;" json:"fchhpw"  form:"fchhpw" `
     FCHHPW01 string `gorm:"column:FCHHPW01;" json:"fchhpw01"  form:"fchhpw01" `
     FCHPW string `gorm:"column:FCHPW;" json:"fchpw"  form:"fchpw" `
     FCHPW01 string `gorm:"column:FCHPW01;" json:"fchpw01"  form:"fchpw01" `
     FCLOGIN string `gorm:"column:FCLOGIN;" json:"fclogin"  form:"fclogin" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMENULANG string `gorm:"column:FCMENULANG;" json:"fcmenulang"  form:"fcmenulang" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPTHAICOD string `gorm:"column:FCPTHAICOD;" json:"fcpthaicod"  form:"fcpthaicod" `
     FCPW string `gorm:"column:FCPW;" json:"fcpw"  form:"fcpw" `
     FCRCODE string `gorm:"column:FCRCODE;" json:"fcrcode"  form:"fcrcode" `
     FCREPOLANG string `gorm:"column:FCREPOLANG;" json:"fcrepolang"  form:"fcrepolang" `
     FCSCRNLANG string `gorm:"column:FCSCRNLANG;" json:"fcscrnlang"  form:"fcscrnlang" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUSERTEM string `gorm:"column:FCUSERTEM;" json:"fcusertem"  form:"fcusertem" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FCYEARMODE string `gorm:"column:FCYEARMODE;" json:"fcyearmode"  form:"fcyearmode" `
     FDLSTCHGPA string `gorm:"column:FDLSTCHGPA;" json:"fdlstchgpa"  form:"fdlstchgpa" `
     FDLSTLOGIN time.Time `gorm:"column:FDLSTLOGIN;" json:"fdlstlogin"  form:"fdlstlogin" default:"now"`
     FIMILLISEC int64 `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNMISSPASS float64 `gorm:"column:FNMISSPASS;" json:"fnmisspass"  form:"fnmisspass" `
     FNSTATUS float64 `gorm:"column:FNSTATUS;" json:"fnstatus"  form:"fnstatus" `
     FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}
func (Emplr) TableName() string{
return "EMPLR"
}

func (obj *Emplr) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
