package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Mfgclosi struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBEGWKTIM string    `gorm:"column:FCBEGWKTIM;" json:"fcbegwktim"  form:"fcbegwktim" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCTRLSTOC string    `gorm:"column:FCCTRLSTOC;" json:"fcctrlstoc"  form:"fcctrlstoc" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCENDWKTIM string    `gorm:"column:FCENDWKTIM;" json:"fcendwktim"  form:"fcendwktim" `
	FCGAG      string    `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FCGVPOLICY string    `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
	FCIOTYPE   string    `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
	FCISMAINPD string    `gorm:"column:FCISMAINPD;" json:"fcismainpd"  form:"fcismainpd" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLOT      string    `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMFGCLOSH string    `gorm:"column:FCMFGCLOSH;" json:"fcmfgclosh"  form:"fcmfgclosh" `
	FCMFUM     string    `gorm:"column:FCMFUM;" json:"fcmfum"  form:"fcmfum" `
	FCMLOTH    string    `gorm:"column:FCMLOTH;" json:"fcmloth"  form:"fcmloth" `
	FCMORDERH  string    `gorm:"column:FCMORDERH;" json:"fcmorderh"  form:"fcmorderh" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string    `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFPDTYP string    `gorm:"column:FCREFPDTYP;" json:"fcrefpdtyp"  form:"fcrefpdtyp" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTORE    string    `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
	FCSTUM     string    `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCTOOLUSE  string    `gorm:"column:FCTOOLUSE;" json:"fctooluse"  form:"fctooluse" `
	FCTRAN2WHY string    `gorm:"column:FCTRAN2WHY;" json:"fctran2why"  form:"fctran2why" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string    `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDEXPIRE   string    `gorm:"column:FDEXPIRE;" json:"fdexpire"  form:"fdexpire" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
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
	FNACTMANHR string    `gorm:"column:FNACTMANHR;" json:"fnactmanhr"  form:"fnactmanhr" `
	FNACTMFQTY string    `gorm:"column:FNACTMFQTY;" json:"fnactmfqty"  form:"fnactmfqty" `
	FNACTQTY   string    `gorm:"column:FNACTQTY;" json:"fnactqty"  form:"fnactqty" `
	FNACTSTQTY string    `gorm:"column:FNACTSTQTY;" json:"fnactstqty"  form:"fnactstqty" `
	FNCOSTADJ  string    `gorm:"column:FNCOSTADJ;" json:"fncostadj"  form:"fncostadj" `
	FNCOSTADJ2 string    `gorm:"column:FNCOSTADJ2;" json:"fncostadj2"  form:"fncostadj2" `
	FNCOSTADJ3 string    `gorm:"column:FNCOSTADJ3;" json:"fncostadj3"  form:"fncostadj3" `
	FNCOSTADJ4 string    `gorm:"column:FNCOSTADJ4;" json:"fncostadj4"  form:"fncostadj4" `
	FNCOSTADJ5 string    `gorm:"column:FNCOSTADJ5;" json:"fncostadj5"  form:"fncostadj5" `
	FNCOSTADJ6 string    `gorm:"column:FNCOSTADJ6;" json:"fncostadj6"  form:"fncostadj6" `
	FNCOSTADJ7 string    `gorm:"column:FNCOSTADJ7;" json:"fncostadj7"  form:"fncostadj7" `
	FNCOSTADJ8 string    `gorm:"column:FNCOSTADJ8;" json:"fncostadj8"  form:"fncostadj8" `
	FNCOSTADJ9 string    `gorm:"column:FNCOSTADJ9;" json:"fncostadj9"  form:"fncostadj9" `
	FNCOSTADJA string    `gorm:"column:FNCOSTADJA;" json:"fncostadja"  form:"fncostadja" `
	FNCOSTAMT  string    `gorm:"column:FNCOSTAMT;" json:"fncostamt"  form:"fncostamt" `
	FNMFQTY    string    `gorm:"column:FNMFQTY;" json:"fnmfqty"  form:"fnmfqty" `
	FNMFUMQTY  string    `gorm:"column:FNMFUMQTY;" json:"fnmfumqty"  form:"fnmfumqty" `
	FNMOMFQTY  string    `gorm:"column:FNMOMFQTY;" json:"fnmomfqty"  form:"fnmomfqty" `
	FNMOQTY    string    `gorm:"column:FNMOQTY;" json:"fnmoqty"  form:"fnmoqty" `
	FNMOSTQTY  string    `gorm:"column:FNMOSTQTY;" json:"fnmostqty"  form:"fnmostqty" `
	FNPORTION  string    `gorm:"column:FNPORTION;" json:"fnportion"  form:"fnportion" `
	FNQTY      string    `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNRETMFQTY string    `gorm:"column:FNRETMFQTY;" json:"fnretmfqty"  form:"fnretmfqty" `
	FNRETQTY   string    `gorm:"column:FNRETQTY;" json:"fnretqty"  form:"fnretqty" `
	FNRETSTQTY string    `gorm:"column:FNRETSTQTY;" json:"fnretstqty"  form:"fnretstqty" `
	FNSTQTY    string    `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FNSTUMQTY  string    `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNUMQTY    string    `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNWASTEMFQ string    `gorm:"column:FNWASTEMFQ;" json:"fnwastemfq"  form:"fnwastemfq" `
	FNWASTEQTY string    `gorm:"column:FNWASTEQTY;" json:"fnwasteqty"  form:"fnwasteqty" `
	FNWASTESTQ string    `gorm:"column:FNWASTESTQ;" json:"fnwastestq"  form:"fnwastestq" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Mfgclosi) TableName() string {
	return "MFGCLOSI"
}

func (obj *Mfgclosi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
