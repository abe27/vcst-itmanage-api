package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Formh2 struct {
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCDEVISIZE string    `gorm:"column:FCDEVISIZE;" json:"fcdevisize"  form:"fcdevisize" `
	FCDIRECTOR string    `gorm:"column:FCDIRECTOR;" json:"fcdirector"  form:"fcdirector" `
	FCDOCTYPE  string    `gorm:"column:FCDOCTYPE;" json:"fcdoctype"  form:"fcdoctype" `
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
	FCFILETYPE string    `gorm:"column:FCFILETYPE;" json:"fcfiletype"  form:"fcfiletype" `
	FCFONTNAME string    `gorm:"column:FCFONTNAME;" json:"fcfontname"  form:"fcfontname" `
	FCFONTSTYL string    `gorm:"column:FCFONTSTYL;" json:"fcfontstyl"  form:"fcfontstyl" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCHRTYPE   string    `gorm:"column:FCHRTYPE;" json:"fchrtype"  form:"fchrtype" `
	FCISODOC   string    `gorm:"column:FCISODOC;" json:"fcisodoc"  form:"fcisodoc" `
	FCISOISSUE string    `gorm:"column:FCISOISSUE;" json:"fcisoissue"  form:"fcisoissue" `
	FCISOREVI  string    `gorm:"column:FCISOREVI;" json:"fcisorevi"  form:"fcisorevi" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCLAYOUTVS string    `gorm:"column:FCLAYOUTVS;" json:"fclayoutvs"  form:"fclayoutvs" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string    `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCOUTFILE  string    `gorm:"column:FCOUTFILE;" json:"fcoutfile"  form:"fcoutfile" `
	FCPGRPITEM string    `gorm:"column:FCPGRPITEM;" json:"fcpgrpitem"  form:"fcpgrpitem" `
	FCPITCH    string    `gorm:"column:FCPITCH;" json:"fcpitch"  form:"fcpitch" `
	FCPORT     string    `gorm:"column:FCPORT;" json:"fcport"  form:"fcport" `
	FCPRNCOND  string    `gorm:"column:FCPRNCOND;" json:"fcprncond"  form:"fcprncond" `
	FCPRNNAME  string    `gorm:"column:FCPRNNAME;" json:"fcprnname"  form:"fcprnname" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCREPSIZE  string    `gorm:"column:FCREPSIZE;" json:"fcrepsize"  form:"fcrepsize" `
	FCRESOLIST string    `gorm:"column:FCRESOLIST;" json:"fcresolist"  form:"fcresolist" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSHOWCOMP string    `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCTHAICODE string    `gorm:"column:FCTHAICODE;" json:"fcthaicode"  form:"fcthaicode" `
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
	FMEXPR     string    `gorm:"column:FMEXPR;" json:"fmexpr"  form:"fmexpr" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMLAYOUTRE string    `gorm:"column:FMLAYOUTRE;" json:"fmlayoutre"  form:"fmlayoutre" `
	FNBEGPAGE  float64   `gorm:"column:FNBEGPAGE;" json:"fnbegpage"  form:"fnbegpage" `
	FNBGCOLOR  float64   `gorm:"column:FNBGCOLOR;" json:"fnbgcolor"  form:"fnbgcolor" `
	FNCANCHANG float64   `gorm:"column:FNCANCHANG;" json:"fncanchang"  form:"fncanchang" `
	FNCHECKLEN float64   `gorm:"column:FNCHECKLEN;" json:"fnchecklen"  form:"fnchecklen" `
	FNCHECKLIN float64   `gorm:"column:FNCHECKLIN;" json:"fnchecklin"  form:"fnchecklin" `
	FNDATAAREA float64   `gorm:"column:FNDATAAREA;" json:"fndataarea"  form:"fndataarea" `
	FNENDPAGE  float64   `gorm:"column:FNENDPAGE;" json:"fnendpage"  form:"fnendpage" `
	FNFILE     float64   `gorm:"column:FNFILE;" json:"fnfile"  form:"fnfile" `
	FNFONTSIZE float64   `gorm:"column:FNFONTSIZE;" json:"fnfontsize"  form:"fnfontsize" `
	FNFONTUNDE float64   `gorm:"column:FNFONTUNDE;" json:"fnfontunde"  form:"fnfontunde" `
	FNFORMPP   float64   `gorm:"column:FNFORMPP;" json:"fnformpp"  form:"fnformpp" `
	FNGRAPH    float64   `gorm:"column:FNGRAPH;" json:"fngraph"  form:"fngraph" `
	FNGRPHPRIN float64   `gorm:"column:FNGRPHPRIN;" json:"fngrphprin"  form:"fngrphprin" `
	FNINCHHEIG float64   `gorm:"column:FNINCHHEIG;" json:"fninchheig"  form:"fninchheig" `
	FNINCHWIDT float64   `gorm:"column:FNINCHWIDT;" json:"fninchwidt"  form:"fninchwidt" `
	FNLEFTMARG float64   `gorm:"column:FNLEFTMARG;" json:"fnleftmarg"  form:"fnleftmarg" `
	FNLINSPACE float64   `gorm:"column:FNLINSPACE;" json:"fnlinspace"  form:"fnlinspace" `
	FNLPP      float64   `gorm:"column:FNLPP;" json:"fnlpp"  form:"fnlpp" `
	FNMAXDLINE float64   `gorm:"column:FNMAXDLINE;" json:"fnmaxdline"  form:"fnmaxdline" `
	FNNCOPY    float64   `gorm:"column:FNNCOPY;" json:"fnncopy"  form:"fnncopy" `
	FNNOCHR12  float64   `gorm:"column:FNNOCHR12;" json:"fnnochr12"  form:"fnnochr12" `
	FNOBJCOLOR float64   `gorm:"column:FNOBJCOLOR;" json:"fnobjcolor"  form:"fnobjcolor" `
	FNPAGEBYPA float64   `gorm:"column:FNPAGEBYPA;" json:"fnpagebypa"  form:"fnpagebypa" `
	FNSCREEN   float64   `gorm:"column:FNSCREEN;" json:"fnscreen"  form:"fnscreen" `
	FNSERVRESO float64   `gorm:"column:FNSERVRESO;" json:"fnservreso"  form:"fnservreso" `
	FNSHOWPAGE float64   `gorm:"column:FNSHOWPAGE;" json:"fnshowpage"  form:"fnshowpage" `
	FNTEXTPRIN float64   `gorm:"column:FNTEXTPRIN;" json:"fntextprin"  form:"fntextprin" `
	FNTXTWIDTH float64   `gorm:"column:FNTXTWIDTH;" json:"fntxtwidth"  form:"fntxtwidth" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNUSEFFEED float64   `gorm:"column:FNUSEFFEED;" json:"fnuseffeed"  form:"fnuseffeed" `
	FNUSELANG1 float64   `gorm:"column:FNUSELANG1;" json:"fnuselang1"  form:"fnuselang1" `
	FNUSELANG2 float64   `gorm:"column:FNUSELANG2;" json:"fnuselang2"  form:"fnuselang2" `
	FNUSESOMEP float64   `gorm:"column:FNUSESOMEP;" json:"fnusesomep"  form:"fnusesomep" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
}

func (Formh2) TableName() string {
	return "FORMH2"
}

func (obj *Formh2) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
