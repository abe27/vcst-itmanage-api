package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Book struct {
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string    `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCSTARTCOD string    `gorm:"column:FCSTARTCOD;" json:"fcstartcod"  form:"fcstartcod" `
	FCGMODHEAD string    `gorm:"column:FCGMODHEAD;" json:"fcgmodhead"  form:"fcgmodhead" `
	FCACCBOOK  string    `gorm:"column:FCACCBOOK;" json:"fcaccbook"  form:"fcaccbook" `
	FCGLHEADRE string    `gorm:"column:FCGLHEADRE;" json:"fcglheadre"  form:"fcglheadre" `
	FCTYPELINK string    `gorm:"column:FCTYPELINK;" json:"fctypelink"  form:"fctypelink" `
	FNNITEM    float64   `gorm:"column:FNNITEM;" json:"fnnitem"  form:"fnnitem" `
	FCHEADSTR1 string    `gorm:"column:FCHEADSTR1;" json:"fcheadstr1"  form:"fcheadstr1" `
	FCHEADSTR2 string    `gorm:"column:FCHEADSTR2;" json:"fcheadstr2"  form:"fcheadstr2" `
	FCHEADSTR3 string    `gorm:"column:FCHEADSTR3;" json:"fcheadstr3"  form:"fcheadstr3" `
	FCFOOTSTR1 string    `gorm:"column:FCFOOTSTR1;" json:"fcfootstr1"  form:"fcfootstr1" `
	FCFOOTSTR2 string    `gorm:"column:FCFOOTSTR2;" json:"fcfootstr2"  form:"fcfootstr2" `
	FCFOOTSTR3 string    `gorm:"column:FCFOOTSTR3;" json:"fcfootstr3"  form:"fcfootstr3" `
	FCLOCATION string    `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCPSSTYLE  string    `gorm:"column:FCPSSTYLE;" json:"fcpsstyle"  form:"fcpsstyle" `
	FCINCDEC   string    `gorm:"column:FCINCDEC;" json:"fcincdec"  form:"fcincdec" `
	FDLOCKED   string    `gorm:"column:FDLOCKED;" json:"fdlocked"  form:"fdlocked" `
	FNBARCODE  float64   `gorm:"column:FNBARCODE;" json:"fnbarcode"  form:"fnbarcode" `
	FCCTRLORD  string    `gorm:"column:FCCTRLORD;" json:"fcctrlord"  form:"fcctrlord" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCMACHINE  string    `gorm:"column:FCMACHINE;" json:"fcmachine"  form:"fcmachine" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FCPREFIX   string    `gorm:"column:FCPREFIX;" json:"fcprefix"  form:"fcprefix" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCATSTEP   string    `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCCTRLINV  string    `gorm:"column:FCCTRLINV;" json:"fcctrlinv"  form:"fcctrlinv" `
	FCCTRLTXID string    `gorm:"column:FCCTRLTXID;" json:"fcctrltxid"  form:"fcctrltxid" `
	FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCAPVTYPE  string    `gorm:"column:FCAPVTYPE;" json:"fcapvtype"  form:"fcapvtype" `
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FCISBORROW string    `gorm:"column:FCISBORROW;" json:"fcisborrow"  form:"fcisborrow" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCU1STATDF string    `gorm:"column:FCU1STATDF;" json:"fcu1statdf"  form:"fcu1statdf" `
	FCU2STATDF string    `gorm:"column:FCU2STATDF;" json:"fcu2statdf"  form:"fcu2statdf" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCRIGSHOW  string    `gorm:"column:FCRIGSHOW;" json:"fcrigshow"  form:"fcrigshow" `
	FCDOCRUNFM string    `gorm:"column:FCDOCRUNFM;" json:"fcdocrunfm"  form:"fcdocrunfm" `
	FNRUNLEN   float64   `gorm:"column:FNRUNLEN;" json:"fnrunlen"  form:"fnrunlen" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCBANKACCT string    `gorm:"column:FCBANKACCT;" json:"fcbankacct"  form:"fcbankacct" `
	FCCTRLBOQ  string    `gorm:"column:FCCTRLBOQ;" json:"fcctrlboq"  form:"fcctrlboq" `
	FCPDCAT    string    `gorm:"column:FCPDCAT;" json:"fcpdcat"  form:"fcpdcat" `
	FCCTRLREF  string    `gorm:"column:FCCTRLREF;" json:"fcctrlref"  form:"fcctrlref" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCMNMASTEP string    `gorm:"column:FCMNMASTEP;" json:"fcmnmastep"  form:"fcmnmastep" `
	FCCTRLREQ  string    `gorm:"column:FCCTRLREQ;" json:"fcctrlreq"  form:"fcctrlreq" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCDEPTYPE  string    `gorm:"column:FCDEPTYPE;" json:"fcdeptype"  form:"fcdeptype" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
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
	FCU3STATDF string    `gorm:"column:FCU3STATDF;" json:"fcu3statdf"  form:"fcu3statdf" `
	FCU4STATDF string    `gorm:"column:FCU4STATDF;" json:"fcu4statdf"  form:"fcu4statdf" `
	FCU5STATDF string    `gorm:"column:FCU5STATDF;" json:"fcu5statdf"  form:"fcu5statdf" `
	FCU6STATDF string    `gorm:"column:FCU6STATDF;" json:"fcu6statdf"  form:"fcu6statdf" `
	FCU7STATDF string    `gorm:"column:FCU7STATDF;" json:"fcu7statdf"  form:"fcu7statdf" `
	FCU8STATDF string    `gorm:"column:FCU8STATDF;" json:"fcu8statdf"  form:"fcu8statdf" `
	FCU9STATDF string    `gorm:"column:FCU9STATDF;" json:"fcu9statdf"  form:"fcu9statdf" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCDAYLOCK  string    `gorm:"column:FCDAYLOCK;" json:"fcdaylock"  form:"fcdaylock" `
	FCWHTFORM  string    `gorm:"column:FCWHTFORM;" json:"fcwhtform"  form:"fcwhtform" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCSENDETAX string    `gorm:"column:FCSENDETAX;" json:"fcsendetax"  form:"fcsendetax" `
	FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FDSTETAX   string    `gorm:"column:FDSTETAX;" json:"fdstetax"  form:"fdstetax" `
	FCSIGNETAX string    `gorm:"column:FCSIGNETAX;" json:"fcsignetax"  form:"fcsignetax" `
	FCCHGSIGNE string    `gorm:"column:FCCHGSIGNE;" json:"fcchgsigne"  form:"fcchgsigne" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Book) TableName() string {
	return "BOOK"
}
func (obj *Book) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
