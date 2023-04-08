package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Layipd struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCOOR     string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOUNTER  string `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEMPL     string `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCFORMULAS string `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
	FCGLREF    string `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FCGPSTR    string `gorm:"column:FCGPSTR;" json:"fcgpstr"  form:"fcgpstr" `
	FCLAYH     string `gorm:"column:FCLAYH;" json:"fclayh"  form:"fclayh" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPCNGPCON string `gorm:"column:FCPCNGPCON;" json:"fcpcngpcon"  form:"fcpcngpcon" `
	FCPFORMULA string `gorm:"column:FCPFORMULA;" json:"fcpformula"  form:"fcpformula" `
	FCPROD     string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCREFPDTYP string `gorm:"column:FCREFPDTYP;" json:"fcrefpdtyp"  form:"fcrefpdtyp" `
	FCREFPROD  string `gorm:"column:FCREFPROD;" json:"fcrefprod"  form:"fcrefprod" `
	FCREFTYPE  string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCROOTSEQ  string `gorm:"column:FCROOTSEQ;" json:"fcrootseq"  form:"fcrootseq" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSHOWCOMP string `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDDATE     string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDSELLDATE string `gorm:"column:FDSELLDATE;" json:"fdselldate"  form:"fdselldate" `
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
	FNAMT      string `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNCNAMT    string `gorm:"column:FNCNAMT;" json:"fncnamt"  form:"fncnamt" `
	FNDISCAMT  string `gorm:"column:FNDISCAMT;" json:"fndiscamt"  form:"fndiscamt" `
	FNDISCPCN  string `gorm:"column:FNDISCPCN;" json:"fndiscpcn"  form:"fndiscpcn" `
	FNGP       string `gorm:"column:FNGP;" json:"fngp"  form:"fngp" `
	FNNETRECAM string `gorm:"column:FNNETRECAM;" json:"fnnetrecam"  form:"fnnetrecam" `
	FNPRICE    string `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNQTY      string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNRECAMT   string `gorm:"column:FNRECAMT;" json:"fnrecamt"  form:"fnrecamt" `
	FNUMQTY    string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNVATAMT   string `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Layipd) TableName() string {
	return "LAYIPD"
}

func (obj *Layipd) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
