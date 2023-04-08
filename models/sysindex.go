package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Sysindex struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCCDXFILE string `gorm:"column:FCCDXFILE;" json:"fccdxfile"  form:"fccdxfile" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDBFFILE string `gorm:"column:FCDBFFILE;" json:"fcdbffile"  form:"fcdbffile" `
     FCDBSERVER string `gorm:"column:FCDBSERVER;" json:"fcdbserver"  form:"fcdbserver" `
     FCFCHR string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
     FCGROUP string `gorm:"column:FCGROUP;" json:"fcgroup"  form:"fcgroup" `
     FCNAME string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
     FCNAME2 string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCSELECTED string `gorm:"column:FCSELECTED;" json:"fcselected"  form:"fcselected" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCTAGNAME string `gorm:"column:FCTAGNAME;" json:"fctagname"  form:"fctagname" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMINDEXFOR string `gorm:"column:FMINDEXFOR;" json:"fmindexfor"  form:"fmindexfor" `
     FMINDEXKEY string `gorm:"column:FMINDEXKEY;" json:"fmindexkey"  form:"fmindexkey" `
     FNBAKDELEO float64 `gorm:"column:FNBAKDELEO;" json:"fnbakdeleo"  form:"fnbakdeleo" `
     FNBAKRATIO float64 `gorm:"column:FNBAKRATIO;" json:"fnbakratio"  form:"fnbakratio" `
     FNCANDIDAT float64 `gorm:"column:FNCANDIDAT;" json:"fncandidat"  form:"fncandidat" `
     FNISUNIQUE float64 `gorm:"column:FNISUNIQUE;" json:"fnisunique"  form:"fnisunique" `
     FNLASTRECC float64 `gorm:"column:FNLASTRECC;" json:"fnlastrecc"  form:"fnlastrecc" `
     FNUSERSETI float64 `gorm:"column:FNUSERSETI;" json:"fnuserseti"  form:"fnuserseti" `
     FNWARNADDR float64 `gorm:"column:FNWARNADDR;" json:"fnwarnaddr"  form:"fnwarnaddr" `
     FNWARNRATI float64 `gorm:"column:FNWARNRATI;" json:"fnwarnrati"  form:"fnwarnrati" `
     FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
     FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
     FTLASTUPD time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
}
func (Sysindex) TableName() string{
return "SYSINDEX"
}

func (obj *Sysindex) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
