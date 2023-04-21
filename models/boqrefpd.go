package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Boqrefpd struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPARENT   string    `gorm:"column:FCPARENT;" json:"fcparent"  form:"fcparent" `
	FCPDUMSTD  string    `gorm:"column:FCPDUMSTD;" json:"fcpdumstd"  form:"fcpdumstd" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string    `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCPRODUM   string    `gorm:"column:FCPRODUM;" json:"fcprodum"  form:"fcprodum" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSUBJOBH  string    `gorm:"column:FCSUBJOBH;" json:"fcsubjobh"  form:"fcsubjobh" `
	FCTABNAME  string    `gorm:"column:FCTABNAME;" json:"fctabname"  form:"fctabname" `
	FCTABREF   string    `gorm:"column:FCTABREF;" json:"fctabref"  form:"fctabref" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FMREMARK10 string    `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
	FMREMARK2  string    `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FMREMARK3  string    `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
	FMREMARK4  string    `gorm:"column:FMREMARK4;" json:"fmremark4"  form:"fmremark4" `
	FMREMARK5  string    `gorm:"column:FMREMARK5;" json:"fmremark5"  form:"fmremark5" `
	FMREMARK6  string    `gorm:"column:FMREMARK6;" json:"fmremark6"  form:"fmremark6" `
	FMREMARK7  string    `gorm:"column:FMREMARK7;" json:"fmremark7"  form:"fmremark7" `
	FMREMARK8  string    `gorm:"column:FMREMARK8;" json:"fmremark8"  form:"fmremark8" `
	FMREMARK9  string    `gorm:"column:FMREMARK9;" json:"fmremark9"  form:"fmremark9" `
	FNAMT      string    `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNINVAMT   string    `gorm:"column:FNINVAMT;" json:"fninvamt"  form:"fninvamt" `
	FNINVQTY   string    `gorm:"column:FNINVQTY;" json:"fninvqty"  form:"fninvqty" `
	FNLBORCOST string    `gorm:"column:FNLBORCOST;" json:"fnlborcost"  form:"fnlborcost" `
	FNPDUMQTY  string    `gorm:"column:FNPDUMQTY;" json:"fnpdumqty"  form:"fnpdumqty" `
	FNPOAMT    string    `gorm:"column:FNPOAMT;" json:"fnpoamt"  form:"fnpoamt" `
	FNPOQTY    string    `gorm:"column:FNPOQTY;" json:"fnpoqty"  form:"fnpoqty" `
	FNPRAMT    string    `gorm:"column:FNPRAMT;" json:"fnpramt"  form:"fnpramt" `
	FNPRODCOST string    `gorm:"column:FNPRODCOST;" json:"fnprodcost"  form:"fnprodcost" `
	FNPRODQTY  string    `gorm:"column:FNPRODQTY;" json:"fnprodqty"  form:"fnprodqty" `
	FNPRQTY    string    `gorm:"column:FNPRQTY;" json:"fnprqty"  form:"fnprqty" `
	FNQTY      string    `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Boqrefpd) TableName() string {
	return "BOQREFPD"
}

func (obj *Boqrefpd) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
