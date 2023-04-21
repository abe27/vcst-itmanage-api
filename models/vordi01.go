package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vordi01 struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDESTCITY string    `gorm:"column:FCDESTCITY;" json:"fcdestcity"  form:"fcdestcity" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORDERH   string    `gorm:"column:FCORDERH;" json:"fcorderh"  form:"fcorderh" `
	FCORDERI   string    `gorm:"column:FCORDERI;" json:"fcorderi"  form:"fcorderi" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPACKSTYL string    `gorm:"column:FCPACKSTYL;" json:"fcpackstyl"  form:"fcpackstyl" `
	FCPDCODE   string    `gorm:"column:FCPDCODE;" json:"fcpdcode"  form:"fcpdcode" `
	FCPDMAP    string    `gorm:"column:FCPDMAP;" json:"fcpdmap"  form:"fcpdmap" `
	FCPDNAME   string    `gorm:"column:FCPDNAME;" json:"fcpdname"  form:"fcpdname" `
	FCRCVEDI   string    `gorm:"column:FCRCVEDI;" json:"fcrcvedi"  form:"fcrcvedi" `
	FCRECSTAT  string    `gorm:"column:FCRECSTAT;" json:"fcrecstat"  form:"fcrecstat" `
	FCRENEWAL  string    `gorm:"column:FCRENEWAL;" json:"fcrenewal"  form:"fcrenewal" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSENDSTAT string    `gorm:"column:FCSENDSTAT;" json:"fcsendstat"  form:"fcsendstat" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSHIPMODE string    `gorm:"column:FCSHIPMODE;" json:"fcshipmode"  form:"fcshipmode" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUMCODE   string    `gorm:"column:FCUMCODE;" json:"fcumcode"  form:"fcumcode" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVORDH01  string    `gorm:"column:FCVORDH01;" json:"fcvordh01"  form:"fcvordh01" `
	FDDELIVERY string    `gorm:"column:FDDELIVERY;" json:"fddelivery"  form:"fddelivery" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNPRICE    string    `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNQTY      string    `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FTCREATED  string    `gorm:"column:FTCREATED;" json:"ftcreated"  form:"ftcreated" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vordi01) TableName() string {
	return "VORDI01"
}

func (obj *Vordi01) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
