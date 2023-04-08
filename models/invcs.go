package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Invcs struct{
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
     ID int64 `gorm:"column:ID;" json:"id"  form:"id" `
     RUN string `gorm:"column:RUN;" json:"run"  form:"run" `
     TYPE string `gorm:"column:TYPE;" json:"type"  form:"type" `
}
func (Invcs) TableName() string{
return "INVCS"
}

func (obj *Invcs) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
