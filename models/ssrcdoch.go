package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Ssrcdoch struct {
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCDOCH  string `gorm:"column:FCSRCDOCH;" json:"fcsrcdoch"  form:"fcsrcdoch" `
	FCSTRKEY1  string `gorm:"column:FCSTRKEY1;" json:"fcstrkey1"  form:"fcstrkey1" `
	FCSTRKEY2  string `gorm:"column:FCSTRKEY2;" json:"fcstrkey2"  form:"fcstrkey2" `
	FCSTRKEY3  string `gorm:"column:FCSTRKEY3;" json:"fcstrkey3"  form:"fcstrkey3" `
	FCSTRKEY4  string `gorm:"column:FCSTRKEY4;" json:"fcstrkey4"  form:"fcstrkey4" `
	FCSTRKEY5  string `gorm:"column:FCSTRKEY5;" json:"fcstrkey5"  form:"fcstrkey5" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Ssrcdoch) TableName() string {
	return "SSRCDOCH"
}

func (obj *Ssrcdoch) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
