package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Book struct {
	FCACCBOOK string `gorm:"column:FCACCBOOK;" json:"fcaccbook"  form:"fcaccbook" `
	// FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	// FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	// FCAPVTYPE  string    `gorm:"column:FCAPVTYPE;" json:"fcapvtype"  form:"fcapvtype" `
	FCATSTEP string `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	// FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	// FCBANKACCT string    `gorm:"column:FCBANKACCT;" json:"fcbankacct"  form:"fcbankacct" `
	// FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	// FCCHGSIGNE string    `gorm:"column:FCCHGSIGNE;" json:"fcchgsigne"  form:"fcchgsigne" `
	FCCODE     string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	// FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCCTRLBOQ  string    `gorm:"column:FCCTRLBOQ;" json:"fcctrlboq"  form:"fcctrlboq" `
	// FCCTRLINV  string    `gorm:"column:FCCTRLINV;" json:"fcctrlinv"  form:"fcctrlinv" `
	// FCCTRLORD  string    `gorm:"column:FCCTRLORD;" json:"fcctrlord"  form:"fcctrlord" `
	// FCCTRLREF  string    `gorm:"column:FCCTRLREF;" json:"fcctrlref"  form:"fcctrlref" `
	// FCCTRLREQ  string    `gorm:"column:FCCTRLREQ;" json:"fcctrlreq"  form:"fcctrlreq" `
	// FCCTRLTXID string    `gorm:"column:FCCTRLTXID;" json:"fcctrltxid"  form:"fcctrltxid" `
	// FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	// FCDAYLOCK  string    `gorm:"column:FCDAYLOCK;" json:"fcdaylock"  form:"fcdaylock" `
	// FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	// FCDEPTYPE  string    `gorm:"column:FCDEPTYPE;" json:"fcdeptype"  form:"fcdeptype" `
	// FCDOCRUNFM string    `gorm:"column:FCDOCRUNFM;" json:"fcdocrunfm"  form:"fcdocrunfm" `
	// FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	// FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCFCHR string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	// FCFOOTSTR1 string    `gorm:"column:FCFOOTSTR1;" json:"fcfootstr1"  form:"fcfootstr1" `
	// FCFOOTSTR2 string    `gorm:"column:FCFOOTSTR2;" json:"fcfootstr2"  form:"fcfootstr2" `
	// FCFOOTSTR3 string    `gorm:"column:FCFOOTSTR3;" json:"fcfootstr3"  form:"fcfootstr3" `
	// FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	// FCGLHEADRE string    `gorm:"column:FCGLHEADRE;" json:"fcglheadre"  form:"fcglheadre" `
	FCGMODHEAD string `gorm:"column:FCGMODHEAD;" json:"fcgmodhead"  form:"fcgmodhead" `
	// FCHEADSTR1 string    `gorm:"column:FCHEADSTR1;" json:"fcheadstr1"  form:"fcheadstr1" `
	// FCHEADSTR2 string    `gorm:"column:FCHEADSTR2;" json:"fcheadstr2"  form:"fcheadstr2" `
	// FCHEADSTR3 string    `gorm:"column:FCHEADSTR3;" json:"fcheadstr3"  form:"fcheadstr3" `
	// FCINCDEC   string    `gorm:"column:FCINCDEC;" json:"fcincdec"  form:"fcincdec" `
	// FCISBORROW string    `gorm:"column:FCISBORROW;" json:"fcisborrow"  form:"fcisborrow" `
	// FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	// FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	// FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	// FCLOCATION string    `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
	// FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	// FCMACHINE  string    `gorm:"column:FCMACHINE;" json:"fcmachine"  form:"fcmachine" `
	// FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	// FCMNMASTEP string    `gorm:"column:FCMNMASTEP;" json:"fcmnmastep"  form:"fcmnmastep" `
	FCNAME  string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2 string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	// FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCPDCAT    string    `gorm:"column:FCPDCAT;" json:"fcpdcat"  form:"fcpdcat" `
	// FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	// FCPREFIX   string    `gorm:"column:FCPREFIX;" json:"fcprefix"  form:"fcprefix" `
	// FCPSSTYLE  string    `gorm:"column:FCPSSTYLE;" json:"fcpsstyle"  form:"fcpsstyle" `
	FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	// FCRIGSHOW  string    `gorm:"column:FCRIGSHOW;" json:"fcrigshow"  form:"fcrigshow" `
	// FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	// FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	// FCSENDETAX string    `gorm:"column:FCSENDETAX;" json:"fcsendetax"  form:"fcsendetax" `
	// FCSIGNETAX string    `gorm:"column:FCSIGNETAX;" json:"fcsignetax"  form:"fcsignetax" `
	FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	// FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FCSTARTCOD string    `gorm:"column:FCSTARTCOD;" json:"fcstartcod"  form:"fcstartcod" `
	// FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	// FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCTYPELINK string `gorm:"column:FCTYPELINK;" json:"fctypelink"  form:"fctypelink" `
	// FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCU1STATDF string    `gorm:"column:FCU1STATDF;" json:"fcu1statdf"  form:"fcu1statdf" `
	// FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	// FCU2STATDF string    `gorm:"column:FCU2STATDF;" json:"fcu2statdf"  form:"fcu2statdf" `
	// FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	// FCU3STATDF string    `gorm:"column:FCU3STATDF;" json:"fcu3statdf"  form:"fcu3statdf" `
	// FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	// FCU4STATDF string    `gorm:"column:FCU4STATDF;" json:"fcu4statdf"  form:"fcu4statdf" `
	// FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	// FCU5STATDF string    `gorm:"column:FCU5STATDF;" json:"fcu5statdf"  form:"fcu5statdf" `
	// FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	// FCU6STATDF string    `gorm:"column:FCU6STATDF;" json:"fcu6statdf"  form:"fcu6statdf" `
	// FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	// FCU7STATDF string    `gorm:"column:FCU7STATDF;" json:"fcu7statdf"  form:"fcu7statdf" `
	// FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	// FCU8STATDF string    `gorm:"column:FCU8STATDF;" json:"fcu8statdf"  form:"fcu8statdf" `
	// FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	// FCU9STATDF string    `gorm:"column:FCU9STATDF;" json:"fcu9statdf"  form:"fcu9statdf" `
	// FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	// FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	// FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	// FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	// FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FCWHOUSE string `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	// FCWHTFORM  string    `gorm:"column:FCWHTFORM;" json:"fcwhtform"  form:"fcwhtform" `
	// FDLOCKED   string    `gorm:"column:FDLOCKED;" json:"fdlocked"  form:"fdlocked" `
	// FDSTETAX   string    `gorm:"column:FDSTETAX;" json:"fdstetax"  form:"fdstetax" `
	// FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	// FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	// FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	// FNBARCODE  float64   `gorm:"column:FNBARCODE;" json:"fnbarcode"  form:"fnbarcode" `
	// FNNITEM    float64   `gorm:"column:FNNITEM;" json:"fnnitem"  form:"fnnitem" `
	// FNRUNLEN   float64   `gorm:"column:FNRUNLEN;" json:"fnrunlen"  form:"fnrunlen" `
	// FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	// FTLASTEDIT time.Time    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD  time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
	CreatedBy *Empl     `gorm:"foreignKey:FCCREATEBY;references:FCSKID;" json:"created_by"`
	Corp      *Corp     `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	WHouse    *WHouse   `gorm:"foreignKey:FCWHOUSE;references:FCSKID;" json:"whouse"`
	Accbook   *Accbook  `gorm:"foreignKey:FCACCBOOK;references:FCSKID;" json:"accbook"`
	Gmodhead  *Gmodhead `gorm:"foreignKey:FCGMODHEAD;references:FCSKID;" json:"gmodhead"`
	Branch    *Branch   `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
}

func (Book) TableName() string {
	return "BOOK"
}
func (obj *Book) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
