package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Saladjts struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBASETYPE string    `gorm:"column:FCBASETYPE;" json:"fcbasetype"  form:"fcbasetype" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCDETAXCON string    `gorm:"column:FCDETAXCON;" json:"fcdetaxcon"  form:"fcdetaxcon" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEMPL     string    `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCEMPLTYPE string    `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
	FCEMPLX1   string    `gorm:"column:FCEMPLX1;" json:"fcemplx1"  form:"fcemplx1" `
	FCHSALARY  string    `gorm:"column:FCHSALARY;" json:"fchsalary"  form:"fchsalary" `
	FCINCRULE  string    `gorm:"column:FCINCRULE;" json:"fcincrule"  form:"fcincrule" `
	FCISCALFUN string    `gorm:"column:FCISCALFUN;" json:"fciscalfun"  form:"fciscalfun" `
	FCISCALSOC string    `gorm:"column:FCISCALSOC;" json:"fciscalsoc"  form:"fciscalsoc" `
	FCISCALTAX string    `gorm:"column:FCISCALTAX;" json:"fciscaltax"  form:"fciscaltax" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCMTH      string    `gorm:"column:FCMTH;" json:"fcmth"  form:"fcmth" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPAYTIMES string    `gorm:"column:FCPAYTIMES;" json:"fcpaytimes"  form:"fcpaytimes" `
	FCPRATEBY  string    `gorm:"column:FCPRATEBY;" json:"fcprateby"  form:"fcprateby" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCPYPERIOD string    `gorm:"column:FCPYPERIOD;" json:"fcpyperiod"  form:"fcpyperiod" `
	FCRANK     string    `gorm:"column:FCRANK;" json:"fcrank"  form:"fcrank" `
	FCRATEBY   string    `gorm:"column:FCRATEBY;" json:"fcrateby"  form:"fcrateby" `
	FCREMARK   string    `gorm:"column:FCREMARK;" json:"fcremark"  form:"fcremark" `
	FCSALADJ   string    `gorm:"column:FCSALADJ;" json:"fcsaladj"  form:"fcsaladj" `
	FCSALADJDF string    `gorm:"column:FCSALADJDF;" json:"fcsaladjdf"  form:"fcsaladjdf" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSYSTYPE  string    `gorm:"column:FCSYSTYPE;" json:"fcsystype"  form:"fcsystype" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCYR       string    `gorm:"column:FCYR;" json:"fcyr"  form:"fcyr" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNAMT      string    `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNPCNBY    string    `gorm:"column:FNPCNBY;" json:"fnpcnby"  form:"fnpcnby" `
	FNQTY      string    `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNSALRATE  string    `gorm:"column:FNSALRATE;" json:"fnsalrate"  form:"fnsalrate" `
	FNWTAXPCN  string    `gorm:"column:FNWTAXPCN;" json:"fnwtaxpcn"  form:"fnwtaxpcn" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Saladjts) TableName() string {
	return "SALADJTS"
}

func (obj *Saladjts) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
