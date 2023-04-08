package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vpacki01 struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCGLREF    string `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FCH1_14    string `gorm:"column:FCH1_14;" json:"fch1_14"  form:"fch1_14" `
	FCH1_16    string `gorm:"column:FCH1_16;" json:"fch1_16"  form:"fch1_16" `
	FCH1_17    string `gorm:"column:FCH1_17;" json:"fch1_17"  form:"fch1_17" `
	FCH1_18    string `gorm:"column:FCH1_18;" json:"fch1_18"  form:"fch1_18" `
	FCH1_23    string `gorm:"column:FCH1_23;" json:"fch1_23"  form:"fch1_23" `
	FCH1_24    string `gorm:"column:FCH1_24;" json:"fch1_24"  form:"fch1_24" `
	FCH1_25    string `gorm:"column:FCH1_25;" json:"fch1_25"  form:"fch1_25" `
	FCH1_26    string `gorm:"column:FCH1_26;" json:"fch1_26"  form:"fch1_26" `
	FCH1_27    string `gorm:"column:FCH1_27;" json:"fch1_27"  form:"fch1_27" `
	FCH1_28    string `gorm:"column:FCH1_28;" json:"fch1_28"  form:"fch1_28" `
	FCH1_29    string `gorm:"column:FCH1_29;" json:"fch1_29"  form:"fch1_29" `
	FCH1_30    string `gorm:"column:FCH1_30;" json:"fch1_30"  form:"fch1_30" `
	FCH1_31    string `gorm:"column:FCH1_31;" json:"fch1_31"  form:"fch1_31" `
	FCH1_32    string `gorm:"column:FCH1_32;" json:"fch1_32"  form:"fch1_32" `
	FCH1_33    string `gorm:"column:FCH1_33;" json:"fch1_33"  form:"fch1_33" `
	FCH1_34    string `gorm:"column:FCH1_34;" json:"fch1_34"  form:"fch1_34" `
	FCH1_35    string `gorm:"column:FCH1_35;" json:"fch1_35"  form:"fch1_35" `
	FCH1_36    string `gorm:"column:FCH1_36;" json:"fch1_36"  form:"fch1_36" `
	FCH1_37    string `gorm:"column:FCH1_37;" json:"fch1_37"  form:"fch1_37" `
	FCH1_38    string `gorm:"column:FCH1_38;" json:"fch1_38"  form:"fch1_38" `
	FCH1_39    string `gorm:"column:FCH1_39;" json:"fch1_39"  form:"fch1_39" `
	FCH1_40    string `gorm:"column:FCH1_40;" json:"fch1_40"  form:"fch1_40" `
	FCH1_41    string `gorm:"column:FCH1_41;" json:"fch1_41"  form:"fch1_41" `
	FCH1_42    string `gorm:"column:FCH1_42;" json:"fch1_42"  form:"fch1_42" `
	FCH2_14    string `gorm:"column:FCH2_14;" json:"fch2_14"  form:"fch2_14" `
	FCH2_15    string `gorm:"column:FCH2_15;" json:"fch2_15"  form:"fch2_15" `
	FCH2_16    string `gorm:"column:FCH2_16;" json:"fch2_16"  form:"fch2_16" `
	FCH2_17    string `gorm:"column:FCH2_17;" json:"fch2_17"  form:"fch2_17" `
	FCH2_18    string `gorm:"column:FCH2_18;" json:"fch2_18"  form:"fch2_18" `
	FCH2_19    string `gorm:"column:FCH2_19;" json:"fch2_19"  form:"fch2_19" `
	FCH2_20    string `gorm:"column:FCH2_20;" json:"fch2_20"  form:"fch2_20" `
	FCH2_21    string `gorm:"column:FCH2_21;" json:"fch2_21"  form:"fch2_21" `
	FCH2_22    string `gorm:"column:FCH2_22;" json:"fch2_22"  form:"fch2_22" `
	FCH2_23    string `gorm:"column:FCH2_23;" json:"fch2_23"  form:"fch2_23" `
	FCH2_24    string `gorm:"column:FCH2_24;" json:"fch2_24"  form:"fch2_24" `
	FCH2_30    string `gorm:"column:FCH2_30;" json:"fch2_30"  form:"fch2_30" `
	FCH2_31    string `gorm:"column:FCH2_31;" json:"fch2_31"  form:"fch2_31" `
	FCH2_32    string `gorm:"column:FCH2_32;" json:"fch2_32"  form:"fch2_32" `
	FCH2_37    string `gorm:"column:FCH2_37;" json:"fch2_37"  form:"fch2_37" `
	FCH3_14    string `gorm:"column:FCH3_14;" json:"fch3_14"  form:"fch3_14" `
	FCH3_15    string `gorm:"column:FCH3_15;" json:"fch3_15"  form:"fch3_15" `
	FCH3_16    string `gorm:"column:FCH3_16;" json:"fch3_16"  form:"fch3_16" `
	FCH3_23    string `gorm:"column:FCH3_23;" json:"fch3_23"  form:"fch3_23" `
	FCH3_25    string `gorm:"column:FCH3_25;" json:"fch3_25"  form:"fch3_25" `
	FCH3_26    string `gorm:"column:FCH3_26;" json:"fch3_26"  form:"fch3_26" `
	FCH3_27    string `gorm:"column:FCH3_27;" json:"fch3_27"  form:"fch3_27" `
	FCH3_28    string `gorm:"column:FCH3_28;" json:"fch3_28"  form:"fch3_28" `
	FCH3_29    string `gorm:"column:FCH3_29;" json:"fch3_29"  form:"fch3_29" `
	FCH3_30    string `gorm:"column:FCH3_30;" json:"fch3_30"  form:"fch3_30" `
	FCH3_31    string `gorm:"column:FCH3_31;" json:"fch3_31"  form:"fch3_31" `
	FCH3_32    string `gorm:"column:FCH3_32;" json:"fch3_32"  form:"fch3_32" `
	FCH3_33    string `gorm:"column:FCH3_33;" json:"fch3_33"  form:"fch3_33" `
	FCH3_34    string `gorm:"column:FCH3_34;" json:"fch3_34"  form:"fch3_34" `
	FCH3_35    string `gorm:"column:FCH3_35;" json:"fch3_35"  form:"fch3_35" `
	FCH3_36    string `gorm:"column:FCH3_36;" json:"fch3_36"  form:"fch3_36" `
	FCH3_37    string `gorm:"column:FCH3_37;" json:"fch3_37"  form:"fch3_37" `
	FCH3_39    string `gorm:"column:FCH3_39;" json:"fch3_39"  form:"fch3_39" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPACKH    string `gorm:"column:FCPACKH;" json:"fcpackh"  form:"fcpackh" `
	FCPACKI    string `gorm:"column:FCPACKI;" json:"fcpacki"  form:"fcpacki" `
	FCRCVEDI   string `gorm:"column:FCRCVEDI;" json:"fcrcvedi"  form:"fcrcvedi" `
	FCRECSTAT  string `gorm:"column:FCRECSTAT;" json:"fcrecstat"  form:"fcrecstat" `
	FCREFPROD  string `gorm:"column:FCREFPROD;" json:"fcrefprod"  form:"fcrefprod" `
	FCRENEWAL  string `gorm:"column:FCRENEWAL;" json:"fcrenewal"  form:"fcrenewal" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSENDSTAT string `gorm:"column:FCSENDSTAT;" json:"fcsendstat"  form:"fcsendstat" `
	FCSEQ      string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVPACKH01 string `gorm:"column:FCVPACKH01;" json:"fcvpackh01"  form:"fcvpackh01" `
	FIMILISEC  string `gorm:"column:FIMILISEC;" json:"fimilisec"  form:"fimilisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNH1_11    string `gorm:"column:FNH1_11;" json:"fnh1_11"  form:"fnh1_11" `
	FNH1_12    string `gorm:"column:FNH1_12;" json:"fnh1_12"  form:"fnh1_12" `
	FNH1_13    string `gorm:"column:FNH1_13;" json:"fnh1_13"  form:"fnh1_13" `
	FNH1_15    string `gorm:"column:FNH1_15;" json:"fnh1_15"  form:"fnh1_15" `
	FNH1_19    string `gorm:"column:FNH1_19;" json:"fnh1_19"  form:"fnh1_19" `
	FNH1_20    string `gorm:"column:FNH1_20;" json:"fnh1_20"  form:"fnh1_20" `
	FNH1_21    string `gorm:"column:FNH1_21;" json:"fnh1_21"  form:"fnh1_21" `
	FNH1_22    string `gorm:"column:FNH1_22;" json:"fnh1_22"  form:"fnh1_22" `
	FNH2_11    string `gorm:"column:FNH2_11;" json:"fnh2_11"  form:"fnh2_11" `
	FNH2_12    string `gorm:"column:FNH2_12;" json:"fnh2_12"  form:"fnh2_12" `
	FNH2_13    string `gorm:"column:FNH2_13;" json:"fnh2_13"  form:"fnh2_13" `
	FNH2_25    string `gorm:"column:FNH2_25;" json:"fnh2_25"  form:"fnh2_25" `
	FNH2_26    string `gorm:"column:FNH2_26;" json:"fnh2_26"  form:"fnh2_26" `
	FNH2_27    string `gorm:"column:FNH2_27;" json:"fnh2_27"  form:"fnh2_27" `
	FNH2_28    string `gorm:"column:FNH2_28;" json:"fnh2_28"  form:"fnh2_28" `
	FNH2_29    string `gorm:"column:FNH2_29;" json:"fnh2_29"  form:"fnh2_29" `
	FNH2_33    string `gorm:"column:FNH2_33;" json:"fnh2_33"  form:"fnh2_33" `
	FNH2_34    string `gorm:"column:FNH2_34;" json:"fnh2_34"  form:"fnh2_34" `
	FNH2_35    string `gorm:"column:FNH2_35;" json:"fnh2_35"  form:"fnh2_35" `
	FNH2_36    string `gorm:"column:FNH2_36;" json:"fnh2_36"  form:"fnh2_36" `
	FNH3_11    string `gorm:"column:FNH3_11;" json:"fnh3_11"  form:"fnh3_11" `
	FNH3_12    string `gorm:"column:FNH3_12;" json:"fnh3_12"  form:"fnh3_12" `
	FNH3_13    string `gorm:"column:FNH3_13;" json:"fnh3_13"  form:"fnh3_13" `
	FNH3_17    string `gorm:"column:FNH3_17;" json:"fnh3_17"  form:"fnh3_17" `
	FNH3_18    string `gorm:"column:FNH3_18;" json:"fnh3_18"  form:"fnh3_18" `
	FNH3_19    string `gorm:"column:FNH3_19;" json:"fnh3_19"  form:"fnh3_19" `
	FNH3_20    string `gorm:"column:FNH3_20;" json:"fnh3_20"  form:"fnh3_20" `
	FNH3_21    string `gorm:"column:FNH3_21;" json:"fnh3_21"  form:"fnh3_21" `
	FNH3_22    string `gorm:"column:FNH3_22;" json:"fnh3_22"  form:"fnh3_22" `
	FNH3_24    string `gorm:"column:FNH3_24;" json:"fnh3_24"  form:"fnh3_24" `
	FNH3_38    string `gorm:"column:FNH3_38;" json:"fnh3_38"  form:"fnh3_38" `
	FTCREATED  string `gorm:"column:FTCREATED;" json:"ftcreated"  form:"ftcreated" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vpacki01) TableName() string {
	return "VPACKI01"
}

func (obj *Vpacki01) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
