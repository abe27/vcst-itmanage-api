package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Refprod struct {
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCGLHEAD   string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCGLREF    string    `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCREFPDTYP string    `gorm:"column:FCREFPDTYP;" json:"fcrefpdtyp"  form:"fcrefpdtyp" `
	FCIOTYPE   string    `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCPRODTYPE string    `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCSHOWCOMP string    `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
	FCROOTSEQ  string    `gorm:"column:FCROOTSEQ;" json:"fcrootseq"  form:"fcrootseq" `
	FCPFORMULA string    `gorm:"column:FCPFORMULA;" json:"fcpformula"  form:"fcpformula" `
	FCFORMULAS string    `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCUM       string    `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string    `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FNUMQTY    float64   `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNQTY      float64   `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FCSTUM     string    `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSTUMSTD  string    `gorm:"column:FCSTUMSTD;" json:"fcstumstd"  form:"fcstumstd" `
	FNSTUMQTY  float64   `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNSTQTY    float64   `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FNPRICE    float64   `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNCOST     float64   `gorm:"column:FNCOST;" json:"fncost"  form:"fncost" `
	FNDISCAMT  float64   `gorm:"column:FNDISCAMT;" json:"fndiscamt"  form:"fndiscamt" `
	FNDISCPCN  float64   `gorm:"column:FNDISCPCN;" json:"fndiscpcn"  form:"fndiscpcn" `
	FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FNQTYATDAT float64   `gorm:"column:FNQTYATDAT;" json:"fnqtyatdat"  form:"fnqtyatdat" `
	FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNVATAMT   float64   `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCLOT      string    `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FNCOMMISSI float64   `gorm:"column:FNCOMMISSI;" json:"fncommissi"  form:"fncommissi" `
	FCPROMOTE  string    `gorm:"column:FCPROMOTE;" json:"fcpromote"  form:"fcpromote" `
	FCACCHART  string    `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNVATPORT  float64   `gorm:"column:FNVATPORT;" json:"fnvatport"  form:"fnvatport" `
	FNVATPORTA float64   `gorm:"column:FNVATPORTA;" json:"fnvatporta"  form:"fnvatporta" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FNCOSTAMT  float64   `gorm:"column:FNCOSTAMT;" json:"fncostamt"  form:"fncostamt" `
	FNPRICEKE  float64   `gorm:"column:FNPRICEKE;" json:"fnpriceke"  form:"fnpriceke" `
	FNDISCAMTK float64   `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	FNXRATE    float64   `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FNCOSTADJ  float64   `gorm:"column:FNCOSTADJ;" json:"fncostadj"  form:"fncostadj" `
	FMREMARK2  string    `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FMREMARK3  string    `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
	FMREMARK4  string    `gorm:"column:FMREMARK4;" json:"fmremark4"  form:"fmremark4" `
	FMREMARK5  string    `gorm:"column:FMREMARK5;" json:"fmremark5"  form:"fmremark5" `
	FMREMARK6  string    `gorm:"column:FMREMARK6;" json:"fmremark6"  form:"fmremark6" `
	FMREMARK7  string    `gorm:"column:FMREMARK7;" json:"fmremark7"  form:"fmremark7" `
	FMREMARK8  string    `gorm:"column:FMREMARK8;" json:"fmremark8"  form:"fmremark8" `
	FMREMARK9  string    `gorm:"column:FMREMARK9;" json:"fmremark9"  form:"fmremark9" `
	FMREMARK10 string    `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FCTRANWHY  string    `gorm:"column:FCTRANWHY;" json:"fctranwhy"  form:"fctranwhy" `
	FCTIMESTAM string    `gorm:"column:FCTIMESTAM;" json:"fctimestam"  form:"fctimestam" `
	FCQCPASS   string    `gorm:"column:FCQCPASS;" json:"fcqcpass"  form:"fcqcpass" `
	FCGVPOLICY string    `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
	FCMFGRUNB  string    `gorm:"column:FCMFGRUNB;" json:"fcmfgrunb"  form:"fcmfgrunb" `
	FCMFGRUNE  string    `gorm:"column:FCMFGRUNE;" json:"fcmfgrune"  form:"fcmfgrune" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FCLAYSTEP  string    `gorm:"column:FCLAYSTEP;" json:"fclaystep"  form:"fclaystep" `
	FCPROCHAIN string    `gorm:"column:FCPROCHAIN;" json:"fcprochain"  form:"fcprochain" `
	FCXFERSTEP string    `gorm:"column:FCXFERSTEP;" json:"fcxferstep"  form:"fcxferstep" `
	FTXFER     string    `gorm:"column:FTXFER;" json:"ftxfer"  form:"ftxfer" `
	FCREFDO    string    `gorm:"column:FCREFDO;" json:"fcrefdo"  form:"fcrefdo" `
	FDEXPIRE   string    `gorm:"column:FDEXPIRE;" json:"fdexpire"  form:"fdexpire" `
	FCSTORE    string    `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
	FCGAG      string    `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FMTEXTKE   string    `gorm:"column:FMTEXTKE;" json:"fmtextke"  form:"fmtextke" `
	FDMFGDATE  string    `gorm:"column:FDMFGDATE;" json:"fdmfgdate"  form:"fdmfgdate" `
	FCTXNO     string    `gorm:"column:FCTXNO;" json:"fctxno"  form:"fctxno" `
	FCRECOSTTY string    `gorm:"column:FCRECOSTTY;" json:"fcrecostty"  form:"fcrecostty" `
	FNDEFAPRIC float64   `gorm:"column:FNDEFAPRIC;" json:"fndefapric"  form:"fndefapric" `
	FCGL       string    `gorm:"column:FCGL;" json:"fcgl"  form:"fcgl" `
	FNREFQTY   float64   `gorm:"column:FNREFQTY;" json:"fnrefqty"  form:"fnrefqty" `
	FCTXID     string    `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
	FCSUBTXID  string    `gorm:"column:FCSUBTXID;" json:"fcsubtxid"  form:"fcsubtxid" `
	FDDELIVERY string    `gorm:"column:FDDELIVERY;" json:"fddelivery"  form:"fddelivery" `
	FCMORDERH  string    `gorm:"column:FCMORDERH;" json:"fcmorderh"  form:"fcmorderh" `
	FCSTEPD    string    `gorm:"column:FCSTEPD;" json:"fcstepd"  form:"fcstepd" `
	FTRECEIVE  string    `gorm:"column:FTRECEIVE;" json:"ftreceive"  form:"ftreceive" `
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FCMAINJOBH string    `gorm:"column:FCMAINJOBH;" json:"fcmainjobh"  form:"fcmainjobh" `
	FCSUBJOBH  string    `gorm:"column:FCSUBJOBH;" json:"fcsubjobh"  form:"fcsubjobh" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCCOUNTER  string    `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCMORDERI  string    `gorm:"column:FCMORDERI;" json:"fcmorderi"  form:"fcmorderi" `
	FCRECALBY  string    `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FCMNMGLH   string    `gorm:"column:FCMNMGLH;" json:"fcmnmglh"  form:"fcmnmglh" `
	FCMNMGL    string    `gorm:"column:FCMNMGL;" json:"fcmnmgl"  form:"fcmnmgl" `
	FNPAYAMT   float64   `gorm:"column:FNPAYAMT;" json:"fnpayamt"  form:"fnpayamt" `
	FNPAYAMTKE float64   `gorm:"column:FNPAYAMTKE;" json:"fnpayamtke"  form:"fnpayamtke" `
	FNWTAXRATE float64   `gorm:"column:FNWTAXRATE;" json:"fnwtaxrate"  form:"fnwtaxrate" `
	FNWTAXAMT  float64   `gorm:"column:FNWTAXAMT;" json:"fnwtaxamt"  form:"fnwtaxamt" `
	FNWTAXAMTK float64   `gorm:"column:FNWTAXAMTK;" json:"fnwtaxamtk"  form:"fnwtaxamtk" `
	FNCOSTAVG  float64   `gorm:"column:FNCOSTAVG;" json:"fncostavg"  form:"fncostavg" `
	FNCOSTFIFO float64   `gorm:"column:FNCOSTFIFO;" json:"fncostfifo"  form:"fncostfifo" `
	FNCOSTLOT  float64   `gorm:"column:FNCOSTLOT;" json:"fncostlot"  form:"fncostlot" `
	FCLOCATION string    `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
	FCPRICEBY  string    `gorm:"column:FCPRICEBY;" json:"fcpriceby"  form:"fcpriceby" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FNPRODAGE  float64   `gorm:"column:FNPRODAGE;" json:"fnprodage"  form:"fnprodage" `
	FNCOSTSTD  float64   `gorm:"column:FNCOSTSTD;" json:"fncoststd"  form:"fncoststd" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCREMLOCAT string    `gorm:"column:FCREMLOCAT;" json:"fcremlocat"  form:"fcremlocat" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Refprod) TableName() string {
	return "REFPROD"
}

func (obj *Refprod) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
