package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vbciprii struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCHKDT    string    `gorm:"column:FCCHKDT;" json:"fcchkdt"  form:"fcchkdt" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVBCIPRIH string    `gorm:"column:FCVBCIPRIH;" json:"fcvbciprih"  form:"fcvbciprih" `
	FDNCIUPD   string    `gorm:"column:FDNCIUPD;" json:"fdnciupd"  form:"fdnciupd" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNPRICE0   string    `gorm:"column:FNPRICE0;" json:"fnprice0"  form:"fnprice0" `
	FNPRICE1   string    `gorm:"column:FNPRICE1;" json:"fnprice1"  form:"fnprice1" `
	FNPRICE2   string    `gorm:"column:FNPRICE2;" json:"fnprice2"  form:"fnprice2" `
	FNPRICE3   string    `gorm:"column:FNPRICE3;" json:"fnprice3"  form:"fnprice3" `
	FNPRICE4   string    `gorm:"column:FNPRICE4;" json:"fnprice4"  form:"fnprice4" `
	FNPRICE5   string    `gorm:"column:FNPRICE5;" json:"fnprice5"  form:"fnprice5" `
	FNPRICE6   string    `gorm:"column:FNPRICE6;" json:"fnprice6"  form:"fnprice6" `
	FNPRICE7   string    `gorm:"column:FNPRICE7;" json:"fnprice7"  form:"fnprice7" `
	FNPRICE8   string    `gorm:"column:FNPRICE8;" json:"fnprice8"  form:"fnprice8" `
	FNPRICE9   string    `gorm:"column:FNPRICE9;" json:"fnprice9"  form:"fnprice9" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vbciprii) TableName() string {
	return "VBCIPRII"
}

func (obj *Vbciprii) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
