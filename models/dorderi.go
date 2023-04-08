package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Dorderi struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPROVEB string `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCBOICARD  string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCOOR     string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOUNTER  string `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCTRLSTOC string `gorm:"column:FCCTRLSTOC;" json:"fcctrlstoc"  form:"fcctrlstoc" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCDISCSTR  string `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCDORDERH  string `gorm:"column:FCDORDERH;" json:"fcdorderh"  form:"fcdorderh" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCFORMULAS string `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
	FCGAG      string `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FCGLREF    string `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FCGVPOLICY string `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
	FCIOTYPE   string `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
	FCISCHG    string `gorm:"column:FCISCHG;" json:"fcischg"  form:"fcischg" `
	FCJOB      string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLOCATION string `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
	FCLOT      string `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMSTEP    string `gorm:"column:FCMSTEP;" json:"fcmstep"  form:"fcmstep" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPFORMULA string `gorm:"column:FCPFORMULA;" json:"fcpformula"  form:"fcpformula" `
	FCPLANT    string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPRICEBY  string `gorm:"column:FCPRICEBY;" json:"fcpriceby"  form:"fcpriceby" `
	FCPROD     string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCPROJ     string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFPDTYP string `gorm:"column:FCREFPDTYP;" json:"fcrefpdtyp"  form:"fcrefpdtyp" `
	FCREFPROD  string `gorm:"column:FCREFPROD;" json:"fcrefprod"  form:"fcrefprod" `
	FCREFTYPE  string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCREMLOCAT string `gorm:"column:FCREMLOCAT;" json:"fcremlocat"  form:"fcremlocat" `
	FCROOTSEQ  string `gorm:"column:FCROOTSEQ;" json:"fcrootseq"  form:"fcrootseq" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSHOWCOMP string `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTEP2    string `gorm:"column:FCSTEP2;" json:"fcstep2"  form:"fcstep2" `
	FCSTORE    string `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
	FCSTUM     string `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSTUMSTD  string `gorm:"column:FCSTUMSTD;" json:"fcstumstd"  form:"fcstumstd" `
	FCSUBBR    string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVATISOUT string `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FCWHOUSE   string `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FDAPPROVE  string `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" `
	FDDATE     string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDDELIVERY string `gorm:"column:FDDELIVERY;" json:"fddelivery"  form:"fddelivery" `
	FDEXPIRE   string `gorm:"column:FDEXPIRE;" json:"fdexpire"  form:"fdexpire" `
	FDMFGDATE  string `gorm:"column:FDMFGDATE;" json:"fdmfgdate"  form:"fdmfgdate" `
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
	FMTEXTKE   string `gorm:"column:FMTEXTKE;" json:"fmtextke"  form:"fmtextke" `
	FNBACKQTY  string `gorm:"column:FNBACKQTY;" json:"fnbackqty"  form:"fnbackqty" `
	FNDISCAMT  string `gorm:"column:FNDISCAMT;" json:"fndiscamt"  form:"fndiscamt" `
	FNDISCAMTK string `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	FNDISCPCN  string `gorm:"column:FNDISCPCN;" json:"fndiscpcn"  form:"fndiscpcn" `
	FNINVQTY   string `gorm:"column:FNINVQTY;" json:"fninvqty"  form:"fninvqty" `
	FNPRICE    string `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNPRICEKE  string `gorm:"column:FNPRICEKE;" json:"fnpriceke"  form:"fnpriceke" `
	FNQTY      string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNSTQTY    string `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FNSTUMQTY  string `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNUMQTY    string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNVATAMT   string `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATRATE  string `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNXRATE    string `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Dorderi) TableName() string {
	return "DORDERI"
}

func (obj *Dorderi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
