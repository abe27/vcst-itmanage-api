package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vtufallc struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOOK     string `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBOOKREF  string `gorm:"column:FCBOOKREF;" json:"fcbookref"  form:"fcbookref" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLOT      string `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORDERH   string `gorm:"column:FCORDERH;" json:"fcorderh"  form:"fcorderh" `
	FCORDERI   string `gorm:"column:FCORDERI;" json:"fcorderi"  form:"fcorderi" `
	FCORDHREF  string `gorm:"column:FCORDHREF;" json:"fcordhref"  form:"fcordhref" `
	FCORDIREF  string `gorm:"column:FCORDIREF;" json:"fcordiref"  form:"fcordiref" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPDSTH    string `gorm:"column:FCPDSTH;" json:"fcpdsth"  form:"fcpdsth" `
	FCPROD     string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCREFTYPE  string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTUM     string `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSTUMSTD  string `gorm:"column:FCSTUMSTD;" json:"fcstumstd"  form:"fcstumstd" `
	FCSUBTYPE  string `gorm:"column:FCSUBTYPE;" json:"fcsubtype"  form:"fcsubtype" `
	FCTXID     string `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
	FCTXIDREF  string `gorm:"column:FCTXIDREF;" json:"fctxidref"  form:"fctxidref" `
	FCTYPE     string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCWHOUSE   string `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FDDATE     string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
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
	FNINQTY    string `gorm:"column:FNINQTY;" json:"fninqty"  form:"fninqty" `
	FNINQTYM   string `gorm:"column:FNINQTYM;" json:"fninqtym"  form:"fninqtym" `
	FNINUW     string `gorm:"column:FNINUW;" json:"fninuw"  form:"fninuw" `
	FNINUWM    string `gorm:"column:FNINUWM;" json:"fninuwm"  form:"fninuwm" `
	FNOUTQTY   string `gorm:"column:FNOUTQTY;" json:"fnoutqty"  form:"fnoutqty" `
	FNOUTQTYM  string `gorm:"column:FNOUTQTYM;" json:"fnoutqtym"  form:"fnoutqtym" `
	FNQTY      string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNSTQTY    string `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FNSTUMQTY  string `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNUMQTY    string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Vtufallc) TableName() string {
	return "VTUFALLC"
}

func (obj *Vtufallc) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
