package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Orderi struct {
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCORDERH   string    `gorm:"column:FCORDERH;" json:"fcorderh"  form:"fcorderh" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCREFPDTYP string    `gorm:"column:FCREFPDTYP;" json:"fcrefpdtyp"  form:"fcrefpdtyp" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCPRODTYPE string    `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCSHOWCOMP string    `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
	FCROOTSEQ  string    `gorm:"column:FCROOTSEQ;" json:"fcrootseq"  form:"fcrootseq" `
	FCFORMULAS string    `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
	FCPFORMULA string    `gorm:"column:FCPFORMULA;" json:"fcpformula"  form:"fcpformula" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCUM       string    `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string    `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FNUMQTY    float64   `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNQTY      float64   `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNBACKQTY  float64   `gorm:"column:FNBACKQTY;" json:"fnbackqty"  form:"fnbackqty" `
	FNPRICE    float64   `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNDISCAMT  float64   `gorm:"column:FNDISCAMT;" json:"fndiscamt"  form:"fndiscamt" `
	FNDISCPCN  float64   `gorm:"column:FNDISCPCN;" json:"fndiscpcn"  form:"fndiscpcn" `
	FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNVATAMT   float64   `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FCLOT      string    `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FCISCHG    string    `gorm:"column:FCISCHG;" json:"fcischg"  form:"fcischg" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FCCTRLSTOC string    `gorm:"column:FCCTRLSTOC;" json:"fcctrlstoc"  form:"fcctrlstoc" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FNPRICEKE  float64   `gorm:"column:FNPRICEKE;" json:"fnpriceke"  form:"fnpriceke" `
	FNDISCAMTK float64   `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	FNXRATE    float64   `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FMREMARK2  string    `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FMREMARK3  string    `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
	FMREMARK4  string    `gorm:"column:FMREMARK4;" json:"fmremark4"  form:"fmremark4" `
	FMREMARK5  string    `gorm:"column:FMREMARK5;" json:"fmremark5"  form:"fmremark5" `
	FMREMARK6  string    `gorm:"column:FMREMARK6;" json:"fmremark6"  form:"fmremark6" `
	FMREMARK7  string    `gorm:"column:FMREMARK7;" json:"fmremark7"  form:"fmremark7" `
	FMREMARK8  string    `gorm:"column:FMREMARK8;" json:"fmremark8"  form:"fmremark8" `
	FMREMARK9  string    `gorm:"column:FMREMARK9;" json:"fmremark9"  form:"fmremark9" `
	FMREMARK10 string    `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
	FCIOTYPE   string    `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
	FCSTUM     string    `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSTUMSTD  string    `gorm:"column:FCSTUMSTD;" json:"fcstumstd"  form:"fcstumstd" `
	FNSTUMQTY  float64   `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNSTQTY    float64   `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FDDELIVERY time.Time `gorm:"column:FDDELIVERY;" json:"fddelivery"  form:"fddelivery" default:"now"`
	FCGVPOLICY string    `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FDAPPROVE  time.Time `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" default:"now"`
	FCSTEP2    string    `gorm:"column:FCSTEP2;" json:"fcstep2"  form:"fcstep2" `
	FCMSTEP    string    `gorm:"column:FCMSTEP;" json:"fcmstep"  form:"fcmstep" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FCPROCHAIN string    `gorm:"column:FCPROCHAIN;" json:"fcprochain"  form:"fcprochain" `
	FCSTORE    string    `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
	FCPRIORITY string    `gorm:"column:FCPRIORITY;" json:"fcpriority"  form:"fcpriority" `
	FCGAG      string    `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FMTEXTKE   string    `gorm:"column:FMTEXTKE;" json:"fmtextke"  form:"fmtextke" `
	FNPRQTY    float64   `gorm:"column:FNPRQTY;" json:"fnprqty"  form:"fnprqty" `
	FCTXID     string    `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
	FNPLANQTY  float64   `gorm:"column:FNPLANQTY;" json:"fnplanqty"  form:"fnplanqty" `
	FCPDSTHPAC string    `gorm:"column:FCPDSTHPAC;" json:"fcpdsthpac"  form:"fcpdsthpac" `
	FNDEFAPRIC float64   `gorm:"column:FNDEFAPRIC;" json:"fndefapric"  form:"fndefapric" `
	FCSUBTXID  string    `gorm:"column:FCSUBTXID;" json:"fcsubtxid"  form:"fcsubtxid" `
	FNPACKQTY  float64   `gorm:"column:FNPACKQTY;" json:"fnpackqty"  form:"fnpackqty" `
	FNDOQTY    float64   `gorm:"column:FNDOQTY;" json:"fndoqty"  form:"fndoqty" `
	FCACCHART  string    `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FCPRICE    string    `gorm:"column:FCPRICE;" json:"fcprice"  form:"fcprice" `
	FCDELITIME string    `gorm:"column:FCDELITIME;" json:"fcdelitime"  form:"fcdelitime" `
	FCBUDCHART string    `gorm:"column:FCBUDCHART;" json:"fcbudchart"  form:"fcbudchart" `
	FCTXIDLINK string    `gorm:"column:FCTXIDLINK;" json:"fctxidlink"  form:"fctxidlink" `
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FCMAINJOBH string    `gorm:"column:FCMAINJOBH;" json:"fcmainjobh"  form:"fcmainjobh" `
	FCSUBJOBH  string    `gorm:"column:FCSUBJOBH;" json:"fcsubjobh"  form:"fcsubjobh" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCOUNTER  string    `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCRECALBY  string    `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FNWTAXRATE float64   `gorm:"column:FNWTAXRATE;" json:"fnwtaxrate"  form:"fnwtaxrate" `
	FNWTAXAMT  float64   `gorm:"column:FNWTAXAMT;" json:"fnwtaxamt"  form:"fnwtaxamt" `
	FNWTAXAMTK float64   `gorm:"column:FNWTAXAMTK;" json:"fnwtaxamtk"  form:"fnwtaxamtk" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCPRICEBY  string    `gorm:"column:FCPRICEBY;" json:"fcpriceby"  form:"fcpriceby" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FDMFGDATE  string    `gorm:"column:FDMFGDATE;" json:"fdmfgdate"  form:"fdmfgdate" `
	FDEXPIRE   string    `gorm:"column:FDEXPIRE;" json:"fdexpire"  form:"fdexpire" `
	FCREMLOCAT string    `gorm:"column:FCREMLOCAT;" json:"fcremlocat"  form:"fcremlocat" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Orderi) TableName() string {
	return "ORDERI"
}
func (obj *Orderi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
