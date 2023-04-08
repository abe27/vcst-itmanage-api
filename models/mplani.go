package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Mplani struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOICARD  string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCOOR     string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCGAG      string `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FCGMACHINE string `gorm:"column:FCGMACHINE;" json:"fcgmachine"  form:"fcgmachine" `
	FCGVPOLICY string `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
	FCJOB      string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMACHINE  string `gorm:"column:FCMACHINE;" json:"fcmachine"  form:"fcmachine" `
	FCMFUM     string `gorm:"column:FCMFUM;" json:"fcmfum"  form:"fcmfum" `
	FCMFUMSTD  string `gorm:"column:FCMFUMSTD;" json:"fcmfumstd"  form:"fcmfumstd" `
	FCMPLANH   string `gorm:"column:FCMPLANH;" json:"fcmplanh"  form:"fcmplanh" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPDSTH    string `gorm:"column:FCPDSTH;" json:"fcpdsth"  form:"fcpdsth" `
	FCPLANT    string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPRIORITY string `gorm:"column:FCPRIORITY;" json:"fcpriority"  form:"fcpriority" `
	FCPROD     string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPROJ     string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTUM     string `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSTUMSTD  string `gorm:"column:FCSTUMSTD;" json:"fcstumstd"  form:"fcstumstd" `
	FCSUBBR    string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDDELIVERY string `gorm:"column:FDDELIVERY;" json:"fddelivery"  form:"fddelivery" `
	FDETDDATE  string `gorm:"column:FDETDDATE;" json:"fdetddate"  form:"fdetddate" `
	FDFINISH   string `gorm:"column:FDFINISH;" json:"fdfinish"  form:"fdfinish" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNFQTY     string `gorm:"column:FNFQTY;" json:"fnfqty"  form:"fnfqty" `
	FNMFFQTY   string `gorm:"column:FNMFFQTY;" json:"fnmffqty"  form:"fnmffqty" `
	FNMFQTY    string `gorm:"column:FNMFQTY;" json:"fnmfqty"  form:"fnmfqty" `
	FNMFUMQTY  string `gorm:"column:FNMFUMQTY;" json:"fnmfumqty"  form:"fnmfumqty" `
	FNQTY      string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNSTFQTY   string `gorm:"column:FNSTFQTY;" json:"fnstfqty"  form:"fnstfqty" `
	FNSTQTY    string `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FNSTUMQTY  string `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNUMQTY    string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Mplani) TableName() string {
	return "MPLANI"
}

func (obj *Mplani) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
