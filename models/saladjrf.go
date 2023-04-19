package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Saladjrf struct {
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEMPL     string    `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCEMPLTYPE string    `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMTH      string    `gorm:"column:FCMTH;" json:"fcmth"  form:"fcmth" `
	FCPRATEBY  string    `gorm:"column:FCPRATEBY;" json:"fcprateby"  form:"fcprateby" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCPYPERIOD string    `gorm:"column:FCPYPERIOD;" json:"fcpyperiod"  form:"fcpyperiod" `
	FCPYPERIOO string    `gorm:"column:FCPYPERIOO;" json:"fcpyperioo"  form:"fcpyperioo" `
	FCRANK     string    `gorm:"column:FCRANK;" json:"fcrank"  form:"fcrank" `
	FCREFSKID  string    `gorm:"column:FCREFSKID;" json:"fcrefskid"  form:"fcrefskid" `
	FCREMARK   string    `gorm:"column:FCREMARK;" json:"fcremark"  form:"fcremark" `
	FCREMARKO  string    `gorm:"column:FCREMARKO;" json:"fcremarko"  form:"fcremarko" `
	FCSALADJ   string    `gorm:"column:FCSALADJ;" json:"fcsaladj"  form:"fcsaladj" `
	FCSALADJDF string    `gorm:"column:FCSALADJDF;" json:"fcsaladjdf"  form:"fcsaladjdf" `
	FCSALADJO  string    `gorm:"column:FCSALADJO;" json:"fcsaladjo"  form:"fcsaladjo" `
	FCSALADJTS string    `gorm:"column:FCSALADJTS;" json:"fcsaladjts"  form:"fcsaladjts" `
	FCSALADJTY string    `gorm:"column:FCSALADJTY;" json:"fcsaladjty"  form:"fcsaladjty" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSYSTYPE  string    `gorm:"column:FCSYSTYPE;" json:"fcsystype"  form:"fcsystype" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCYR       string    `gorm:"column:FCYR;" json:"fcyr"  form:"fcyr" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNAMT      string    `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNAMTO     string    `gorm:"column:FNAMTO;" json:"fnamto"  form:"fnamto" `
	FNPCNBY    string    `gorm:"column:FNPCNBY;" json:"fnpcnby"  form:"fnpcnby" `
	FNQTY      string    `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNSALRATE  string    `gorm:"column:FNSALRATE;" json:"fnsalrate"  form:"fnsalrate" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Saladjrf) TableName() string {
	return "SALADJRF"
}

func (obj *Saladjrf) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
