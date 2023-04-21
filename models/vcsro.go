package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vcsro struct {
	FCANS01    string    `gorm:"column:FCANS01;" json:"fcans01"  form:"fcans01" `
	FCANS02    string    `gorm:"column:FCANS02;" json:"fcans02"  form:"fcans02" `
	FCANS03    string    `gorm:"column:FCANS03;" json:"fcans03"  form:"fcans03" `
	FCANS04    string    `gorm:"column:FCANS04;" json:"fcans04"  form:"fcans04" `
	FCANS05    string    `gorm:"column:FCANS05;" json:"fcans05"  form:"fcans05" `
	FCANS06    string    `gorm:"column:FCANS06;" json:"fcans06"  form:"fcans06" `
	FCANS07    string    `gorm:"column:FCANS07;" json:"fcans07"  form:"fcans07" `
	FCANS08    string    `gorm:"column:FCANS08;" json:"fcans08"  form:"fcans08" `
	FCANS09    string    `gorm:"column:FCANS09;" json:"fcans09"  form:"fcans09" `
	FCANS10    string    `gorm:"column:FCANS10;" json:"fcans10"  form:"fcans10" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCONTUSER string    `gorm:"column:FCCONTUSER;" json:"fccontuser"  form:"fccontuser" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEXCUTE   string    `gorm:"column:FCEXCUTE;" json:"fcexcute"  form:"fcexcute" `
	FCISAPV    string    `gorm:"column:FCISAPV;" json:"fcisapv"  form:"fcisapv" `
	FCLIC      string    `gorm:"column:FCLIC;" json:"fclic"  form:"fclic" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCQNCORP   string    `gorm:"column:FCQNCORP;" json:"fcqncorp"  form:"fcqncorp" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREQ01    string    `gorm:"column:FCREQ01;" json:"fcreq01"  form:"fcreq01" `
	FCREQ02    string    `gorm:"column:FCREQ02;" json:"fcreq02"  form:"fcreq02" `
	FCREQ03    string    `gorm:"column:FCREQ03;" json:"fcreq03"  form:"fcreq03" `
	FCREQ04    string    `gorm:"column:FCREQ04;" json:"fcreq04"  form:"fcreq04" `
	FCREQ05    string    `gorm:"column:FCREQ05;" json:"fcreq05"  form:"fcreq05" `
	FCREQ06    string    `gorm:"column:FCREQ06;" json:"fcreq06"  form:"fcreq06" `
	FCREQ07    string    `gorm:"column:FCREQ07;" json:"fcreq07"  form:"fcreq07" `
	FCREQ08    string    `gorm:"column:FCREQ08;" json:"fcreq08"  form:"fcreq08" `
	FCREQ09    string    `gorm:"column:FCREQ09;" json:"fcreq09"  form:"fcreq09" `
	FCREQ10    string    `gorm:"column:FCREQ10;" json:"fcreq10"  form:"fcreq10" `
	FCREQUSER  string    `gorm:"column:FCREQUSER;" json:"fcrequser"  form:"fcrequser" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSUPPORT  string    `gorm:"column:FCSUPPORT;" json:"fcsupport"  form:"fcsupport" `
	FCTOPIC    string    `gorm:"column:FCTOPIC;" json:"fctopic"  form:"fctopic" `
	FCTYPE1    string    `gorm:"column:FCTYPE1;" json:"fctype1"  form:"fctype1" `
	FCTYPE2    string    `gorm:"column:FCTYPE2;" json:"fctype2"  form:"fctype2" `
	FCTYPE3    string    `gorm:"column:FCTYPE3;" json:"fctype3"  form:"fctype3" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDAPVDATE  string    `gorm:"column:FDAPVDATE;" json:"fdapvdate"  form:"fdapvdate" `
	FDFINISH   string    `gorm:"column:FDFINISH;" json:"fdfinish"  form:"fdfinish" `
	FDREQDATE  string    `gorm:"column:FDREQDATE;" json:"fdreqdate"  form:"fdreqdate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vcsro) TableName() string {
	return "VCSRO"
}

func (obj *Vcsro) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
