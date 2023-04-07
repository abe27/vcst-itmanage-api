package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Orderh struct {
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FDDUEDATE  time.Time `gorm:"column:FDDUEDATE;" json:"fdduedate"  form:"fdduedate" default:"now"`
	FDRECEDATE time.Time `gorm:"column:FDRECEDATE;" json:"fdrecedate"  form:"fdrecedate" default:"now"`
	FCBOOK     string    `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FNDISCAMT1 float64   `gorm:"column:FNDISCAMT1;" json:"fndiscamt1"  form:"fndiscamt1" `
	FNDISCAMT2 float64   `gorm:"column:FNDISCAMT2;" json:"fndiscamt2"  form:"fndiscamt2" `
	FNDISCAMTI float64   `gorm:"column:FNDISCAMTI;" json:"fndiscamti"  form:"fndiscamti" `
	FNDISCPCN1 float64   `gorm:"column:FNDISCPCN1;" json:"fndiscpcn1"  form:"fndiscpcn1" `
	FNDISCPCN2 float64   `gorm:"column:FNDISCPCN2;" json:"fndiscpcn2"  form:"fndiscpcn2" `
	FNDISCPCN3 float64   `gorm:"column:FNDISCPCN3;" json:"fndiscpcn3"  form:"fndiscpcn3" `
	FNAMT      float64   `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNAMT2     float64   `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
	FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNVATAMT   float64   `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCEMPL     string    `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCEMZONE   string    `gorm:"column:FCEMZONE;" json:"fcemzone"  form:"fcemzone" `
	FCISCASH   string    `gorm:"column:FCISCASH;" json:"fciscash"  form:"fciscash" `
	FNCREDTERM float64   `gorm:"column:FNCREDTERM;" json:"fncredterm"  form:"fncredterm" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCANCELBY string    `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	FCACSTEP   string    `gorm:"column:FCACSTEP;" json:"fcacstep"  form:"fcacstep" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FCHASRET   string    `gorm:"column:FCHASRET;" json:"fchasret"  form:"fchasret" `
	FCVATDUE   string    `gorm:"column:FCVATDUE;" json:"fcvatdue"  form:"fcvatdue" `
	FCRECVMAN  string    `gorm:"column:FCRECVMAN;" json:"fcrecvman"  form:"fcrecvman" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FCSTEP2    string    `gorm:"column:FCSTEP2;" json:"fcstep2"  form:"fcstep2" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCCURRENCY string    `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	FNAMTKE    float64   `gorm:"column:FNAMTKE;" json:"fnamtke"  form:"fnamtke" `
	FNVATAMTKE float64   `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
	FNDISCAMTK float64   `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	FNXRATE    float64   `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FCFRWHOUSE string    `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCDELICOOR string    `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	FDAPPROVE  time.Time `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" default:"now"`
	FCMSTEP    string    `gorm:"column:FCMSTEP;" json:"fcmstep"  form:"fcmstep" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FCPRIORITY string    `gorm:"column:FCPRIORITY;" json:"fcpriority"  form:"fcpriority" `
	FCHASCHAIN string    `gorm:"column:FCHASCHAIN;" json:"fchaschain"  form:"fchaschain" `
	FCXFERSTEP string    `gorm:"column:FCXFERSTEP;" json:"fcxferstep"  form:"fcxferstep" `
	FTXFER     string    `gorm:"column:FTXFER;" json:"ftxfer"  form:"ftxfer" `
	FMMEMDATA2 string    `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
	FMMEMDATA3 string    `gorm:"column:FMMEMDATA3;" json:"fmmemdata3"  form:"fmmemdata3" `
	FMMEMDATA4 string    `gorm:"column:FMMEMDATA4;" json:"fmmemdata4"  form:"fmmemdata4" `
	FCSHOWPLED string    `gorm:"column:FCSHOWPLED;" json:"fcshowpled"  form:"fcshowpled" `
	FCSTEPX1   string    `gorm:"column:FCSTEPX1;" json:"fcstepx1"  form:"fcstepx1" `
	FMMEMDATA5 string    `gorm:"column:FMMEMDATA5;" json:"fmmemdata5"  form:"fmmemdata5" `
	FCTRADETRM string    `gorm:"column:FCTRADETRM;" json:"fctradetrm"  form:"fctradetrm" `
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FCREASONRQ string    `gorm:"column:FCREASONRQ;" json:"fcreasonrq"  form:"fcreasonrq" `
	FCPAYTERM  string    `gorm:"column:FCPAYTERM;" json:"fcpayterm"  form:"fcpayterm" `
	FDREQDATE  time.Time `gorm:"column:FDREQDATE;" json:"fdreqdate"  form:"fdreqdate" default:"now"`
	FCREQMEN   string    `gorm:"column:FCREQMEN;" json:"fcreqmen"  form:"fcreqmen" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCREQPERSN string    `gorm:"column:FCREQPERSN;" json:"fcreqpersn"  form:"fcreqpersn" `
	FNDAYASSUR float64   `gorm:"column:FNDAYASSUR;" json:"fndayassur"  form:"fndayassur" `
	FCISPDPART string    `gorm:"column:FCISPDPART;" json:"fcispdpart"  form:"fcispdpart" `
	FCPERSON   string    `gorm:"column:FCPERSON;" json:"fcperson"  form:"fcperson" `
	FCLINKAPP1 string    `gorm:"column:FCLINKAPP1;" json:"fclinkapp1"  form:"fclinkapp1" `
	FCLINKSTP1 string    `gorm:"column:FCLINKSTP1;" json:"fclinkstp1"  form:"fclinkstp1" `
	FCLINKAPP2 string    `gorm:"column:FCLINKAPP2;" json:"fclinkapp2"  form:"fclinkapp2" `
	FCLINKSTP2 string    `gorm:"column:FCLINKSTP2;" json:"fclinkstp2"  form:"fclinkstp2" `
	FCCOUNTER  string    `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	FCPDBRAND  string    `gorm:"column:FCPDBRAND;" json:"fcpdbrand"  form:"fcpdbrand" `
	FCRNGTIME  string    `gorm:"column:FCRNGTIME;" json:"fcrngtime"  form:"fcrngtime" `
	FCTRANTYPE string    `gorm:"column:FCTRANTYPE;" json:"fctrantype"  form:"fctrantype" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCRECOMENB string    `gorm:"column:FCRECOMENB;" json:"fcrecomenb"  form:"fcrecomenb" `
	FDRECOMENB string    `gorm:"column:FDRECOMENB;" json:"fdrecomenb"  form:"fdrecomenb" `
	FMMEMDATA6 string    `gorm:"column:FMMEMDATA6;" json:"fmmemdata6"  form:"fmmemdata6" `
	FMMEMDATA7 string    `gorm:"column:FMMEMDATA7;" json:"fmmemdata7"  form:"fmmemdata7" `
	FMMEMDATA8 string    `gorm:"column:FMMEMDATA8;" json:"fmmemdata8"  form:"fmmemdata8" `
	FMMEMDATA9 string    `gorm:"column:FMMEMDATA9;" json:"fmmemdata9"  form:"fmmemdata9" `
	FMMEMDATAA string    `gorm:"column:FMMEMDATAA;" json:"fmmemdataa"  form:"fmmemdataa" `
	FCRECALBY  string    `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FCCARTYPE  string    `gorm:"column:FCCARTYPE;" json:"fccartype"  form:"fccartype" `
	FCDOCSTEP  string    `gorm:"column:FCDOCSTEP;" json:"fcdocstep"  form:"fcdocstep" `
	FCDOCAPRVB string    `gorm:"column:FCDOCAPRVB;" json:"fcdocaprvb"  form:"fcdocaprvb" `
	FDDOCAPRV  string    `gorm:"column:FDDOCAPRV;" json:"fddocaprv"  form:"fddocaprv" `
	FCCRSTEP   string    `gorm:"column:FCCRSTEP;" json:"fccrstep"  form:"fccrstep" `
	FCCRAPRVB  string    `gorm:"column:FCCRAPRVB;" json:"fccraprvb"  form:"fccraprvb" `
	FDCRAPRV   string    `gorm:"column:FDCRAPRV;" json:"fdcraprv"  form:"fdcraprv" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCDELIBY   string    `gorm:"column:FCDELIBY;" json:"fcdeliby"  form:"fcdeliby" `
	FCCVEHICLE string    `gorm:"column:FCCVEHICLE;" json:"fccvehicle"  form:"fccvehicle" `
	FCTHRDPTY  string    `gorm:"column:FCTHRDPTY;" json:"fcthrdpty"  form:"fcthrdpty" `
	FCROUTEDEL string    `gorm:"column:FCROUTEDEL;" json:"fcroutedel"  form:"fcroutedel" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FCDOCFLOWI string    `gorm:"column:FCDOCFLOWI;" json:"fcdocflowi"  form:"fcdocflowi" `
	FMDOCFLOW  string    `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCCONTACTN string    `gorm:"column:FCCONTACTN;" json:"fccontactn"  form:"fccontactn" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCCARRIER  string    `gorm:"column:FCCARRIER;" json:"fccarrier"  form:"fccarrier" `
	FCCARRYTYP string    `gorm:"column:FCCARRYTYP;" json:"fccarrytyp"  form:"fccarrytyp" `
	FCDELIEMPL string    `gorm:"column:FCDELIEMPL;" json:"fcdeliempl"  form:"fcdeliempl" `
	FCSTATAPPV string    `gorm:"column:FCSTATAPPV;" json:"fcstatappv"  form:"fcstatappv" `
	FTSTATAPPV string    `gorm:"column:FTSTATAPPV;" json:"ftstatappv"  form:"ftstatappv" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Orderh) TableName() string {
	return "ORDERH"
}
func (obj *Orderh) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
