package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Morderi struct {
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
	FCIOTYPE   string `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
	FCISMAINPD string `gorm:"column:FCISMAINPD;" json:"fcismainpd"  form:"fcismainpd" `
	FCJOB      string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLOT      string `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMFUM     string `gorm:"column:FCMFUM;" json:"fcmfum"  form:"fcmfum" `
	FCMORDERH  string `gorm:"column:FCMORDERH;" json:"fcmorderh"  form:"fcmorderh" `
	FCOPSEQ    string `gorm:"column:FCOPSEQ;" json:"fcopseq"  form:"fcopseq" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPDSTH    string `gorm:"column:FCPDSTH;" json:"fcpdsth"  form:"fcpdsth" `
	FCPLANT    string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPROCURE  string `gorm:"column:FCPROCURE;" json:"fcprocure"  form:"fcprocure" `
	FCPROD     string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCPROJ     string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFTYPE  string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTORE    string `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
	FCSTUM     string `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSUBBR    string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCSUBTXID  string `gorm:"column:FCSUBTXID;" json:"fcsubtxid"  form:"fcsubtxid" `
	FCTOOLS    string `gorm:"column:FCTOOLS;" json:"fctools"  form:"fctools" `
	FCTXID     string `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCWHOUSE   string `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FDDATE     string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDFINISH   string `gorm:"column:FDFINISH;" json:"fdfinish"  form:"fdfinish" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMREMARK   string `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FMREMARK10 string `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
	FMREMARK2  string `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FMREMARK3  string `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
	FMREMARK4  string `gorm:"column:FMREMARK4;" json:"fmremark4"  form:"fmremark4" `
	FMREMARK5  string `gorm:"column:FMREMARK5;" json:"fmremark5"  form:"fmremark5" `
	FMREMARK6  string `gorm:"column:FMREMARK6;" json:"fmremark6"  form:"fmremark6" `
	FMREMARK7  string `gorm:"column:FMREMARK7;" json:"fmremark7"  form:"fmremark7" `
	FMREMARK8  string `gorm:"column:FMREMARK8;" json:"fmremark8"  form:"fmremark8" `
	FMREMARK9  string `gorm:"column:FMREMARK9;" json:"fmremark9"  form:"fmremark9" `
	FNCONSQTY  string `gorm:"column:FNCONSQTY;" json:"fnconsqty"  form:"fnconsqty" `
	FNCONSUNIT string `gorm:"column:FNCONSUNIT;" json:"fnconsunit"  form:"fnconsunit" `
	FNISSUEQTY string `gorm:"column:FNISSUEQTY;" json:"fnissueqty"  form:"fnissueqty" `
	FNMFQTY    string `gorm:"column:FNMFQTY;" json:"fnmfqty"  form:"fnmfqty" `
	FNMFUMQTY  string `gorm:"column:FNMFUMQTY;" json:"fnmfumqty"  form:"fnmfumqty" `
	FNPORTION  string `gorm:"column:FNPORTION;" json:"fnportion"  form:"fnportion" `
	FNQTY      string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNQTYBFREL string `gorm:"column:FNQTYBFREL;" json:"fnqtybfrel"  form:"fnqtybfrel" `
	FNSTQTY    string `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FNSTUMQTY  string `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNUMQTY    string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNUSEDQTY  string `gorm:"column:FNUSEDQTY;" json:"fnusedqty"  form:"fnusedqty" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Morderi) TableName() string {
	return "MORDERI"
}

func (obj *Morderi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
