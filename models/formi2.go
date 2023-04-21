package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Formi2 struct {
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCATTRSTR  string    `gorm:"column:FCATTRSTR;" json:"fcattrstr"  form:"fcattrstr" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCALCGRP  string    `gorm:"column:FCCALCGRP;" json:"fccalcgrp"  form:"fccalcgrp" `
	FCCENTURYM string    `gorm:"column:FCCENTURYM;" json:"fccenturym"  form:"fccenturym" `
	FCCOLID    string    `gorm:"column:FCCOLID;" json:"fccolid"  form:"fccolid" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDATATYPE string    `gorm:"column:FCDATATYPE;" json:"fcdatatype"  form:"fcdatatype" `
	FCDDELI1   string    `gorm:"column:FCDDELI1;" json:"fcddeli1"  form:"fcddeli1" `
	FCDDELI2   string    `gorm:"column:FCDDELI2;" json:"fcddeli2"  form:"fcddeli2" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCDSTYLE   string    `gorm:"column:FCDSTYLE;" json:"fcdstyle"  form:"fcdstyle" `
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
	FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCFONTNAME string    `gorm:"column:FCFONTNAME;" json:"fcfontname"  form:"fcfontname" `
	FCFONTSTYL string    `gorm:"column:FCFONTSTYL;" json:"fcfontstyl"  form:"fcfontstyl" `
	FCFORMH2   string    `gorm:"column:FCFORMH2;" json:"fcformh2"  form:"fcformh2" `
	FCFORMIOBJ string    `gorm:"column:FCFORMIOBJ;" json:"fcformiobj"  form:"fcformiobj" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCITEMGRP  string    `gorm:"column:FCITEMGRP;" json:"fcitemgrp"  form:"fcitemgrp" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCMONTHMOD string    `gorm:"column:FCMONTHMOD;" json:"fcmonthmod"  form:"fcmonthmod" `
	FCOBJNAME  string    `gorm:"column:FCOBJNAME;" json:"fcobjname"  form:"fcobjname" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCRESOLIST string    `gorm:"column:FCRESOLIST;" json:"fcresolist"  form:"fcresolist" `
	FCROWID    string    `gorm:"column:FCROWID;" json:"fcrowid"  form:"fcrowid" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSET      string    `gorm:"column:FCSET;" json:"fcset"  form:"fcset" `
	FCSIGNFORM string    `gorm:"column:FCSIGNFORM;" json:"fcsignform"  form:"fcsignform" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
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
	FCYEARMODE string    `gorm:"column:FCYEARMODE;" json:"fcyearmode"  form:"fcyearmode" `
	FCZERO     string    `gorm:"column:FCZERO;" json:"fczero"  form:"fczero" `
	FDVALUE    time.Time `gorm:"column:FDVALUE;" json:"fdvalue"  form:"fdvalue" default:"now"`
	FMEXPRESSI string    `gorm:"column:FMEXPRESSI;" json:"fmexpressi"  form:"fmexpressi" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMFORMAT   string    `gorm:"column:FMFORMAT;" json:"fmformat"  form:"fmformat" `
	FMLANG1CAP string    `gorm:"column:FMLANG1CAP;" json:"fmlang1cap"  form:"fmlang1cap" `
	FMLANG2CAP string    `gorm:"column:FMLANG2CAP;" json:"fmlang2cap"  form:"fmlang2cap" `
	FNATNCOPY  float64   `gorm:"column:FNATNCOPY;" json:"fnatncopy"  form:"fnatncopy" `
	FNBACKCOLO float64   `gorm:"column:FNBACKCOLO;" json:"fnbackcolo"  form:"fnbackcolo" `
	FNBACKSTYL float64   `gorm:"column:FNBACKSTYL;" json:"fnbackstyl"  form:"fnbackstyl" `
	FNBORDERCL float64   `gorm:"column:FNBORDERCL;" json:"fnbordercl"  form:"fnbordercl" `
	FNBORDERST float64   `gorm:"column:FNBORDERST;" json:"fnborderst"  form:"fnborderst" `
	FNBORDERWD float64   `gorm:"column:FNBORDERWD;" json:"fnborderwd"  form:"fnborderwd" `
	FNCURVATUR float64   `gorm:"column:FNCURVATUR;" json:"fncurvatur"  form:"fncurvatur" `
	FNDECIPOIN float64   `gorm:"column:FNDECIPOIN;" json:"fndecipoin"  form:"fndecipoin" `
	FNDETLINE  float64   `gorm:"column:FNDETLINE;" json:"fndetline"  form:"fndetline" `
	FNEVERPAGE float64   `gorm:"column:FNEVERPAGE;" json:"fneverpage"  form:"fneverpage" `
	FNFILLCOLO float64   `gorm:"column:FNFILLCOLO;" json:"fnfillcolo"  form:"fnfillcolo" `
	FNFILLSTYL float64   `gorm:"column:FNFILLSTYL;" json:"fnfillstyl"  form:"fnfillstyl" `
	FNFONTSIZE float64   `gorm:"column:FNFONTSIZE;" json:"fnfontsize"  form:"fnfontsize" `
	FNFONTUNDE float64   `gorm:"column:FNFONTUNDE;" json:"fnfontunde"  form:"fnfontunde" `
	FNFORECOLO float64   `gorm:"column:FNFORECOLO;" json:"fnforecolo"  form:"fnforecolo" `
	FNFRSTPAGE float64   `gorm:"column:FNFRSTPAGE;" json:"fnfrstpage"  form:"fnfrstpage" `
	FNFULLHEIG float64   `gorm:"column:FNFULLHEIG;" json:"fnfullheig"  form:"fnfullheig" `
	FNFULLWIDT float64   `gorm:"column:FNFULLWIDT;" json:"fnfullwidt"  form:"fnfullwidt" `
	FNHEIGHT   float64   `gorm:"column:FNHEIGHT;" json:"fnheight"  form:"fnheight" `
	FNLASTPAGE float64   `gorm:"column:FNLASTPAGE;" json:"fnlastpage"  form:"fnlastpage" `
	FNLEFT     float64   `gorm:"column:FNLEFT;" json:"fnleft"  form:"fnleft" `
	FNLEN      float64   `gorm:"column:FNLEN;" json:"fnlen"  form:"fnlen" `
	FNLINE     float64   `gorm:"column:FNLINE;" json:"fnline"  form:"fnline" `
	FNLINESPAC float64   `gorm:"column:FNLINESPAC;" json:"fnlinespac"  form:"fnlinespac" `
	FNOBJBSNAP float64   `gorm:"column:FNOBJBSNAP;" json:"fnobjbsnap"  form:"fnobjbsnap" `
	FNOBJLSNAP float64   `gorm:"column:FNOBJLSNAP;" json:"fnobjlsnap"  form:"fnobjlsnap" `
	FNOBJRSNAP float64   `gorm:"column:FNOBJRSNAP;" json:"fnobjrsnap"  form:"fnobjrsnap" `
	FNOBJTSNAP float64   `gorm:"column:FNOBJTSNAP;" json:"fnobjtsnap"  form:"fnobjtsnap" `
	FNPLAINP   float64   `gorm:"column:FNPLAINP;" json:"fnplainp"  form:"fnplainp" `
	FNPREPRINT float64   `gorm:"column:FNPREPRINT;" json:"fnpreprint"  form:"fnpreprint" `
	FNSERVRESO float64   `gorm:"column:FNSERVRESO;" json:"fnservreso"  form:"fnservreso" `
	FNSHOWMODE float64   `gorm:"column:FNSHOWMODE;" json:"fnshowmode"  form:"fnshowmode" `
	FNTOP      float64   `gorm:"column:FNTOP;" json:"fntop"  form:"fntop" `
	FNTXTBSNAP float64   `gorm:"column:FNTXTBSNAP;" json:"fntxtbsnap"  form:"fntxtbsnap" `
	FNTXTLSNAP float64   `gorm:"column:FNTXTLSNAP;" json:"fntxtlsnap"  form:"fntxtlsnap" `
	FNTXTRSNAP float64   `gorm:"column:FNTXTRSNAP;" json:"fntxtrsnap"  form:"fntxtrsnap" `
	FNTXTTSNAP float64   `gorm:"column:FNTXTTSNAP;" json:"fntxttsnap"  form:"fntxttsnap" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNUSEEXPR  float64   `gorm:"column:FNUSEEXPR;" json:"fnuseexpr"  form:"fnuseexpr" `
	FNUSELANG1 float64   `gorm:"column:FNUSELANG1;" json:"fnuselang1"  form:"fnuselang1" `
	FNUSELANG2 float64   `gorm:"column:FNUSELANG2;" json:"fnuselang2"  form:"fnuselang2" `
	FNUSEVALUE float64   `gorm:"column:FNUSEVALUE;" json:"fnusevalue"  form:"fnusevalue" `
	FNVALUE    float64   `gorm:"column:FNVALUE;" json:"fnvalue"  form:"fnvalue" `
	FNWIDTH    float64   `gorm:"column:FNWIDTH;" json:"fnwidth"  form:"fnwidth" `
	FNWORDWRAP float64   `gorm:"column:FNWORDWRAP;" json:"fnwordwrap"  form:"fnwordwrap" `
	FNZORDER   float64   `gorm:"column:FNZORDER;" json:"fnzorder"  form:"fnzorder" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
}

func (Formi2) TableName() string {
	return "FORMI2"
}

func (obj *Formi2) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
