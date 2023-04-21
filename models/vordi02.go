package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vordi02 struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCE1_11    string    `gorm:"column:FCE1_11;" json:"fce1_11"  form:"fce1_11" `
	FCE1_12    string    `gorm:"column:FCE1_12;" json:"fce1_12"  form:"fce1_12" `
	FCE1_13    string    `gorm:"column:FCE1_13;" json:"fce1_13"  form:"fce1_13" `
	FCE1_14    string    `gorm:"column:FCE1_14;" json:"fce1_14"  form:"fce1_14" `
	FCE1_15    string    `gorm:"column:FCE1_15;" json:"fce1_15"  form:"fce1_15" `
	FCE1_16    string    `gorm:"column:FCE1_16;" json:"fce1_16"  form:"fce1_16" `
	FCE1_17    string    `gorm:"column:FCE1_17;" json:"fce1_17"  form:"fce1_17" `
	FCE1_18    string    `gorm:"column:FCE1_18;" json:"fce1_18"  form:"fce1_18" `
	FCE1_19    string    `gorm:"column:FCE1_19;" json:"fce1_19"  form:"fce1_19" `
	FCE1_20    string    `gorm:"column:FCE1_20;" json:"fce1_20"  form:"fce1_20" `
	FCE1_21    string    `gorm:"column:FCE1_21;" json:"fce1_21"  form:"fce1_21" `
	FCE1_22    string    `gorm:"column:FCE1_22;" json:"fce1_22"  form:"fce1_22" `
	FCE1_23    string    `gorm:"column:FCE1_23;" json:"fce1_23"  form:"fce1_23" `
	FCE1_24    string    `gorm:"column:FCE1_24;" json:"fce1_24"  form:"fce1_24" `
	FCE1_25    string    `gorm:"column:FCE1_25;" json:"fce1_25"  form:"fce1_25" `
	FCE1_26    string    `gorm:"column:FCE1_26;" json:"fce1_26"  form:"fce1_26" `
	FCE1_27    string    `gorm:"column:FCE1_27;" json:"fce1_27"  form:"fce1_27" `
	FCE1_28    string    `gorm:"column:FCE1_28;" json:"fce1_28"  form:"fce1_28" `
	FCE1_29    string    `gorm:"column:FCE1_29;" json:"fce1_29"  form:"fce1_29" `
	FCE1_30    string    `gorm:"column:FCE1_30;" json:"fce1_30"  form:"fce1_30" `
	FCE1_31    string    `gorm:"column:FCE1_31;" json:"fce1_31"  form:"fce1_31" `
	FCE1_32    string    `gorm:"column:FCE1_32;" json:"fce1_32"  form:"fce1_32" `
	FCE1_33    string    `gorm:"column:FCE1_33;" json:"fce1_33"  form:"fce1_33" `
	FCE1_34    string    `gorm:"column:FCE1_34;" json:"fce1_34"  form:"fce1_34" `
	FCE1_35    string    `gorm:"column:FCE1_35;" json:"fce1_35"  form:"fce1_35" `
	FCE1_36    string    `gorm:"column:FCE1_36;" json:"fce1_36"  form:"fce1_36" `
	FCE1_37    string    `gorm:"column:FCE1_37;" json:"fce1_37"  form:"fce1_37" `
	FCE1_38    string    `gorm:"column:FCE1_38;" json:"fce1_38"  form:"fce1_38" `
	FCE1_39    string    `gorm:"column:FCE1_39;" json:"fce1_39"  form:"fce1_39" `
	FCE1_40    string    `gorm:"column:FCE1_40;" json:"fce1_40"  form:"fce1_40" `
	FCE1_41    string    `gorm:"column:FCE1_41;" json:"fce1_41"  form:"fce1_41" `
	FCE1_42    string    `gorm:"column:FCE1_42;" json:"fce1_42"  form:"fce1_42" `
	FCE1_43    string    `gorm:"column:FCE1_43;" json:"fce1_43"  form:"fce1_43" `
	FCE1_44    string    `gorm:"column:FCE1_44;" json:"fce1_44"  form:"fce1_44" `
	FCE1_45    string    `gorm:"column:FCE1_45;" json:"fce1_45"  form:"fce1_45" `
	FCE1_46    string    `gorm:"column:FCE1_46;" json:"fce1_46"  form:"fce1_46" `
	FCE1_47    string    `gorm:"column:FCE1_47;" json:"fce1_47"  form:"fce1_47" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCRCVEDI   string    `gorm:"column:FCRCVEDI;" json:"fcrcvedi"  form:"fcrcvedi" `
	FCRENEWAL  string    `gorm:"column:FCRENEWAL;" json:"fcrenewal"  form:"fcrenewal" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVORDH02  string    `gorm:"column:FCVORDH02;" json:"fcvordh02"  form:"fcvordh02" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FTCREATED  string    `gorm:"column:FTCREATED;" json:"ftcreated"  form:"ftcreated" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vordi02) TableName() string {
	return "VORDI02"
}

func (obj *Vordi02) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
