package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Dorderh struct {
	FCACSTEP   string    `gorm:"column:FCACSTEP;" json:"fcacstep"  form:"fcacstep" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCATSTEP   string    `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBOOK     string    `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCANCELBY string    `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	FCCARRIER  string    `gorm:"column:FCCARRIER;" json:"fccarrier"  form:"fccarrier" `
	FCCARRYTYP string    `gorm:"column:FCCARRYTYP;" json:"fccarrytyp"  form:"fccarrytyp" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCONTACTN string    `gorm:"column:FCCONTACTN;" json:"fccontactn"  form:"fccontactn" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOUNTER  string    `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCCURRENCY string    `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELICOOR string    `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	FCDELIEMPL string    `gorm:"column:FCDELIEMPL;" json:"fcdeliempl"  form:"fcdeliempl" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCDOCFLOWI string    `gorm:"column:FCDOCFLOWI;" json:"fcdocflowi"  form:"fcdocflowi" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEMPL     string    `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCEMZONE   string    `gorm:"column:FCEMZONE;" json:"fcemzone"  form:"fcemzone" `
	FCFRWHOUSE string    `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCGLHEADCA string    `gorm:"column:FCGLHEADCA;" json:"fcglheadca"  form:"fcglheadca" `
	FCGLREF    string    `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FCHASRET   string    `gorm:"column:FCHASRET;" json:"fchasret"  form:"fchasret" `
	FCISCASH   string    `gorm:"column:FCISCASH;" json:"fciscash"  form:"fciscash" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMSTEP    string    `gorm:"column:FCMSTEP;" json:"fcmstep"  form:"fcmstep" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPDBRAND  string    `gorm:"column:FCPDBRAND;" json:"fcpdbrand"  form:"fcpdbrand" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCRECVMAN  string    `gorm:"column:FCRECVMAN;" json:"fcrecvman"  form:"fcrecvman" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCROUTEDEL string    `gorm:"column:FCROUTEDEL;" json:"fcroutedel"  form:"fcroutedel" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTEP2    string    `gorm:"column:FCSTEP2;" json:"fcstep2"  form:"fcstep2" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCTHRDPTY  string    `gorm:"column:FCTHRDPTY;" json:"fcthrdpty"  form:"fcthrdpty" `
	FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVATDUE   string    `gorm:"column:FCVATDUE;" json:"fcvatdue"  form:"fcvatdue" `
	FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FDAPPROVE  string    `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" `
	FDDATE     string    `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
	FDDUEDATE  string    `gorm:"column:FDDUEDATE;" json:"fdduedate"  form:"fdduedate" `
	FDRECEDATE string    `gorm:"column:FDRECEDATE;" json:"fdrecedate"  form:"fdrecedate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMDOCFLOW  string    `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FMMEMDATA2 string    `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
	FMMEMDATA3 string    `gorm:"column:FMMEMDATA3;" json:"fmmemdata3"  form:"fmmemdata3" `
	FMMEMDATA4 string    `gorm:"column:FMMEMDATA4;" json:"fmmemdata4"  form:"fmmemdata4" `
	FMMEMDATA5 string    `gorm:"column:FMMEMDATA5;" json:"fmmemdata5"  form:"fmmemdata5" `
	FMMEMDATA6 string    `gorm:"column:FMMEMDATA6;" json:"fmmemdata6"  form:"fmmemdata6" `
	FMMEMDATA7 string    `gorm:"column:FMMEMDATA7;" json:"fmmemdata7"  form:"fmmemdata7" `
	FMMEMDATA8 string    `gorm:"column:FMMEMDATA8;" json:"fmmemdata8"  form:"fmmemdata8" `
	FMMEMDATA9 string    `gorm:"column:FMMEMDATA9;" json:"fmmemdata9"  form:"fmmemdata9" `
	FMMEMDATAA string    `gorm:"column:FMMEMDATAA;" json:"fmmemdataa"  form:"fmmemdataa" `
	FNAMT      string    `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNAMT2     string    `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
	FNAMTKE    string    `gorm:"column:FNAMTKE;" json:"fnamtke"  form:"fnamtke" `
	FNCREDTERM string    `gorm:"column:FNCREDTERM;" json:"fncredterm"  form:"fncredterm" `
	FNDISCAMT1 string    `gorm:"column:FNDISCAMT1;" json:"fndiscamt1"  form:"fndiscamt1" `
	FNDISCAMT2 string    `gorm:"column:FNDISCAMT2;" json:"fndiscamt2"  form:"fndiscamt2" `
	FNDISCAMTI string    `gorm:"column:FNDISCAMTI;" json:"fndiscamti"  form:"fndiscamti" `
	FNDISCAMTK string    `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	FNDISCPCN1 string    `gorm:"column:FNDISCPCN1;" json:"fndiscpcn1"  form:"fndiscpcn1" `
	FNDISCPCN2 string    `gorm:"column:FNDISCPCN2;" json:"fndiscpcn2"  form:"fndiscpcn2" `
	FNDISCPCN3 string    `gorm:"column:FNDISCPCN3;" json:"fndiscpcn3"  form:"fndiscpcn3" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNVATAMT   string    `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATAMTKE string    `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
	FNVATRATE  string    `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNXRATE    string    `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Dorderh) TableName() string {
	return "DORDERH"
}

func (obj *Dorderh) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
