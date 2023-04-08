package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Pfmah struct {
	FCACSTEP   string `gorm:"column:FCACSTEP;" json:"fcacstep"  form:"fcacstep" `
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPROVEB string `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCATSTEP   string `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCBOICARD  string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBOOK     string `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCANCELBY string `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	FCCASHIER  string `gorm:"column:FCCASHIER;" json:"fccashier"  form:"fccashier" `
	FCCODE     string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCCURRENCY string `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDEBTCODE string `gorm:"column:FCDEBTCODE;" json:"fcdebtcode"  form:"fcdebtcode" `
	FCDEPT     string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCDISCSTR  string `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCDTYPE1   string `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEMPL     string `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCFRWHOUSE string `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCGLHEADCA string `gorm:"column:FCGLHEADCA;" json:"fcglheadca"  form:"fcglheadca" `
	FCHASRET   string `gorm:"column:FCHASRET;" json:"fchasret"  form:"fchasret" `
	FCISCASH   string `gorm:"column:FCISCASH;" json:"fciscash"  form:"fciscash" `
	FCJOB      string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLAYH     string `gorm:"column:FCLAYH;" json:"fclayh"  form:"fclayh" `
	FCLAYSEQ   string `gorm:"column:FCLAYSEQ;" json:"fclayseq"  form:"fclayseq" `
	FCLID      string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMACHINE  string `gorm:"column:FCMACHINE;" json:"fcmachine"  form:"fcmachine" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPERIODS  string `gorm:"column:FCPERIODS;" json:"fcperiods"  form:"fcperiods" `
	FCPROJ     string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCPROMOTE  string `gorm:"column:FCPROMOTE;" json:"fcpromote"  form:"fcpromote" `
	FCRECVMAN  string `gorm:"column:FCRECVMAN;" json:"fcrecvman"  form:"fcrecvman" `
	FCREFNO    string `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTEPX1   string `gorm:"column:FCSTEPX1;" json:"fcstepx1"  form:"fcstepx1" `
	FCSTEPX2   string `gorm:"column:FCSTEPX2;" json:"fcstepx2"  form:"fcstepx2" `
	FCTOWHOUSE string `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVATCOOR  string `gorm:"column:FCVATCOOR;" json:"fcvatcoor"  form:"fcvatcoor" `
	FCVATDUE   string `gorm:"column:FCVATDUE;" json:"fcvatdue"  form:"fcvatdue" `
	FCVATISOUT string `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATSEQ   string `gorm:"column:FCVATSEQ;" json:"fcvatseq"  form:"fcvatseq" `
	FCVATTYPE  string `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FDDATE     string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDDEBTDATE string `gorm:"column:FDDEBTDATE;" json:"fddebtdate"  form:"fddebtdate" `
	FDDUEDATE  string `gorm:"column:FDDUEDATE;" json:"fdduedate"  form:"fdduedate" `
	FDFULLFILL string `gorm:"column:FDFULLFILL;" json:"fdfullfill"  form:"fdfullfill" `
	FDLAYDATE  string `gorm:"column:FDLAYDATE;" json:"fdlaydate"  form:"fdlaydate" `
	FDRECEDATE string `gorm:"column:FDRECEDATE;" json:"fdrecedate"  form:"fdrecedate" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FNAMT      string `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNAMT2     string `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
	FNAMTKE    string `gorm:"column:FNAMTKE;" json:"fnamtke"  form:"fnamtke" `
	FNAMTNOVAT string `gorm:"column:FNAMTNOVAT;" json:"fnamtnovat"  form:"fnamtnovat" `
	FNBEFOAMT  string `gorm:"column:FNBEFOAMT;" json:"fnbefoamt"  form:"fnbefoamt" `
	FNBPAYINV  string `gorm:"column:FNBPAYINV;" json:"fnbpayinv"  form:"fnbpayinv" `
	FNCREDTERM string `gorm:"column:FNCREDTERM;" json:"fncredterm"  form:"fncredterm" `
	FNDEBTZERO string `gorm:"column:FNDEBTZERO;" json:"fndebtzero"  form:"fndebtzero" `
	FNDISCAMT1 string `gorm:"column:FNDISCAMT1;" json:"fndiscamt1"  form:"fndiscamt1" `
	FNDISCAMT2 string `gorm:"column:FNDISCAMT2;" json:"fndiscamt2"  form:"fndiscamt2" `
	FNDISCAMTA string `gorm:"column:FNDISCAMTA;" json:"fndiscamta"  form:"fndiscamta" `
	FNDISCAMTI string `gorm:"column:FNDISCAMTI;" json:"fndiscamti"  form:"fndiscamti" `
	FNDISCAMTK string `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	FNDISCPCN1 string `gorm:"column:FNDISCPCN1;" json:"fndiscpcn1"  form:"fndiscpcn1" `
	FNDISCPCN2 string `gorm:"column:FNDISCPCN2;" json:"fndiscpcn2"  form:"fndiscpcn2" `
	FNDISCPCN3 string `gorm:"column:FNDISCPCN3;" json:"fndiscpcn3"  form:"fndiscpcn3" `
	FNPAYAMT   string `gorm:"column:FNPAYAMT;" json:"fnpayamt"  form:"fnpayamt" `
	FNPAYAMTKE string `gorm:"column:FNPAYAMTKE;" json:"fnpayamtke"  form:"fnpayamtke" `
	FNSPAYAMT  string `gorm:"column:FNSPAYAMT;" json:"fnspayamt"  form:"fnspayamt" `
	FNU1CNT    string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNVATAMT   string `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATAMTKE string `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
	FNVATPORT  string `gorm:"column:FNVATPORT;" json:"fnvatport"  form:"fnvatport" `
	FNVATPORTA string `gorm:"column:FNVATPORTA;" json:"fnvatporta"  form:"fnvatporta" `
	FNVATRATE  string `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNXRATE    string `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Pfmah) TableName() string {
	return "PFMAH"
}

func (obj *Pfmah) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
