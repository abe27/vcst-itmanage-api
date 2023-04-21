package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Refpdpl struct {
	FCACCHART  string    `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOUNTER  string    `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCFORMULAS string    `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
	FCGAG      string    `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FCGL       string    `gorm:"column:FCGL;" json:"fcgl"  form:"fcgl" `
	FCGLHEAD   string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCGLREFPL  string    `gorm:"column:FCGLREFPL;" json:"fcglrefpl"  form:"fcglrefpl" `
	FCGVPOLICY string    `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
	FCIOTYPE   string    `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLAYSTEP  string    `gorm:"column:FCLAYSTEP;" json:"fclaystep"  form:"fclaystep" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLOCATION string    `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
	FCLOT      string    `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMAINJOBH string    `gorm:"column:FCMAINJOBH;" json:"fcmainjobh"  form:"fcmainjobh" `
	FCMFGRUNB  string    `gorm:"column:FCMFGRUNB;" json:"fcmfgrunb"  form:"fcmfgrunb" `
	FCMFGRUNE  string    `gorm:"column:FCMFGRUNE;" json:"fcmfgrune"  form:"fcmfgrune" `
	FCMNMGL    string    `gorm:"column:FCMNMGL;" json:"fcmnmgl"  form:"fcmnmgl" `
	FCMNMGLH   string    `gorm:"column:FCMNMGLH;" json:"fcmnmglh"  form:"fcmnmglh" `
	FCMORDERH  string    `gorm:"column:FCMORDERH;" json:"fcmorderh"  form:"fcmorderh" `
	FCMORDERI  string    `gorm:"column:FCMORDERI;" json:"fcmorderi"  form:"fcmorderi" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPFORMULA string    `gorm:"column:FCPFORMULA;" json:"fcpformula"  form:"fcpformula" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPRICEBY  string    `gorm:"column:FCPRICEBY;" json:"fcpriceby"  form:"fcpriceby" `
	FCPROCHAIN string    `gorm:"column:FCPROCHAIN;" json:"fcprochain"  form:"fcprochain" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string    `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCPROMOTE  string    `gorm:"column:FCPROMOTE;" json:"fcpromote"  form:"fcpromote" `
	FCQCPASS   string    `gorm:"column:FCQCPASS;" json:"fcqcpass"  form:"fcqcpass" `
	FCRECALBY  string    `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	FCRECOSTTY string    `gorm:"column:FCRECOSTTY;" json:"fcrecostty"  form:"fcrecostty" `
	FCREFDO    string    `gorm:"column:FCREFDO;" json:"fcrefdo"  form:"fcrefdo" `
	FCREFPDTYP string    `gorm:"column:FCREFPDTYP;" json:"fcrefpdtyp"  form:"fcrefpdtyp" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCREFTYPE2 string    `gorm:"column:FCREFTYPE2;" json:"fcreftype2"  form:"fcreftype2" `
	FCREMLOCAT string    `gorm:"column:FCREMLOCAT;" json:"fcremlocat"  form:"fcremlocat" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCROOTSEQ  string    `gorm:"column:FCROOTSEQ;" json:"fcrootseq"  form:"fcrootseq" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSHOWCOMP string    `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEPD    string    `gorm:"column:FCSTEPD;" json:"fcstepd"  form:"fcstepd" `
	FCSTORE    string    `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
	FCSTUM     string    `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSTUMSTD  string    `gorm:"column:FCSTUMSTD;" json:"fcstumstd"  form:"fcstumstd" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCSUBJOBH  string    `gorm:"column:FCSUBJOBH;" json:"fcsubjobh"  form:"fcsubjobh" `
	FCSUBTXID  string    `gorm:"column:FCSUBTXID;" json:"fcsubtxid"  form:"fcsubtxid" `
	FCTIMESTAM string    `gorm:"column:FCTIMESTAM;" json:"fctimestam"  form:"fctimestam" `
	FCTRANWHY  string    `gorm:"column:FCTRANWHY;" json:"fctranwhy"  form:"fctranwhy" `
	FCTXID     string    `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
	FCTXNO     string    `gorm:"column:FCTXNO;" json:"fctxno"  form:"fctxno" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string    `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string    `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FCXFERSTEP string    `gorm:"column:FCXFERSTEP;" json:"fcxferstep"  form:"fcxferstep" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDDELIVERY string    `gorm:"column:FDDELIVERY;" json:"fddelivery"  form:"fddelivery" `
	FDEXPIRE   string    `gorm:"column:FDEXPIRE;" json:"fdexpire"  form:"fdexpire" `
	FDMFGDATE  string    `gorm:"column:FDMFGDATE;" json:"fdmfgdate"  form:"fdmfgdate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
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
	FMTEXTKE   string    `gorm:"column:FMTEXTKE;" json:"fmtextke"  form:"fmtextke" `
	FNCOMMISSI string    `gorm:"column:FNCOMMISSI;" json:"fncommissi"  form:"fncommissi" `
	FNCOST     string    `gorm:"column:FNCOST;" json:"fncost"  form:"fncost" `
	FNCOSTADJ  string    `gorm:"column:FNCOSTADJ;" json:"fncostadj"  form:"fncostadj" `
	FNCOSTAMT  string    `gorm:"column:FNCOSTAMT;" json:"fncostamt"  form:"fncostamt" `
	FNCOSTAVG  string    `gorm:"column:FNCOSTAVG;" json:"fncostavg"  form:"fncostavg" `
	FNCOSTFIFO string    `gorm:"column:FNCOSTFIFO;" json:"fncostfifo"  form:"fncostfifo" `
	FNCOSTLOT  string    `gorm:"column:FNCOSTLOT;" json:"fncostlot"  form:"fncostlot" `
	FNCOSTSTD  string    `gorm:"column:FNCOSTSTD;" json:"fncoststd"  form:"fncoststd" `
	FNDEFAPRIC string    `gorm:"column:FNDEFAPRIC;" json:"fndefapric"  form:"fndefapric" `
	FNDISCAMT  string    `gorm:"column:FNDISCAMT;" json:"fndiscamt"  form:"fndiscamt" `
	FNDISCAMTK string    `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	FNDISCPCN  string    `gorm:"column:FNDISCPCN;" json:"fndiscpcn"  form:"fndiscpcn" `
	FNPAYAMT   string    `gorm:"column:FNPAYAMT;" json:"fnpayamt"  form:"fnpayamt" `
	FNPAYAMTKE string    `gorm:"column:FNPAYAMTKE;" json:"fnpayamtke"  form:"fnpayamtke" `
	FNPRICE    string    `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNPRICEKE  string    `gorm:"column:FNPRICEKE;" json:"fnpriceke"  form:"fnpriceke" `
	FNPRODAGE  string    `gorm:"column:FNPRODAGE;" json:"fnprodage"  form:"fnprodage" `
	FNQTY      string    `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNQTYATDAT string    `gorm:"column:FNQTYATDAT;" json:"fnqtyatdat"  form:"fnqtyatdat" `
	FNREFQTY   string    `gorm:"column:FNREFQTY;" json:"fnrefqty"  form:"fnrefqty" `
	FNSTQTY    string    `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FNSTUMQTY  string    `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNUMQTY    string    `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNVATAMT   string    `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATPORT  string    `gorm:"column:FNVATPORT;" json:"fnvatport"  form:"fnvatport" `
	FNVATPORTA string    `gorm:"column:FNVATPORTA;" json:"fnvatporta"  form:"fnvatporta" `
	FNVATRATE  string    `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNWTAXAMT  string    `gorm:"column:FNWTAXAMT;" json:"fnwtaxamt"  form:"fnwtaxamt" `
	FNWTAXAMTK string    `gorm:"column:FNWTAXAMTK;" json:"fnwtaxamtk"  form:"fnwtaxamtk" `
	FNWTAXRATE string    `gorm:"column:FNWTAXRATE;" json:"fnwtaxrate"  form:"fnwtaxrate" `
	FNXRATE    string    `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTRECEIVE  string    `gorm:"column:FTRECEIVE;" json:"ftreceive"  form:"ftreceive" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FTXFER     string    `gorm:"column:FTXFER;" json:"ftxfer"  form:"ftxfer" `
}

func (Refpdpl) TableName() string {
	return "REFPDPL"
}

func (obj *Refpdpl) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
