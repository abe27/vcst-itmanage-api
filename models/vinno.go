package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vinno struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBOOKNO   string    `gorm:"column:FCBOOKNO;" json:"fcbookno"  form:"fcbookno" `
	FCBOOKSECT string    `gorm:"column:FCBOOKSECT;" json:"fcbooksect"  form:"fcbooksect" `
	FCCANCELBY string    `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	FCCARCODE  string    `gorm:"column:FCCARCODE;" json:"fccarcode"  form:"fccarcode" `
	FCCARSTAT  string    `gorm:"column:FCCARSTAT;" json:"fccarstat"  form:"fccarstat" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDRVEMPL  string    `gorm:"column:FCDRVEMPL;" json:"fcdrvempl"  form:"fcdrvempl" `
	FCDRVRMK   string    `gorm:"column:FCDRVRMK;" json:"fcdrvrmk"  form:"fcdrvrmk" `
	FCDRVSTEP  string    `gorm:"column:FCDRVSTEP;" json:"fcdrvstep"  form:"fcdrvstep" `
	FCDRVTYPE  string    `gorm:"column:FCDRVTYPE;" json:"fcdrvtype"  form:"fcdrvtype" `
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
	FCENGINENO string    `gorm:"column:FCENGINENO;" json:"fcengineno"  form:"fcengineno" `
	FCEXDRIVER string    `gorm:"column:FCEXDRIVER;" json:"fcexdriver"  form:"fcexdriver" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCINVNO    string    `gorm:"column:FCINVNO;" json:"fcinvno"  form:"fcinvno" `
	FCISLOCK   string    `gorm:"column:FCISLOCK;" json:"fcislock"  form:"fcislock" `
	FCKAKUDAI  string    `gorm:"column:FCKAKUDAI;" json:"fckakudai"  form:"fckakudai" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMONTYEAR string    `gorm:"column:FCMONTYEAR;" json:"fcmontyear"  form:"fcmontyear" `
	FCORDERH   string    `gorm:"column:FCORDERH;" json:"fcorderh"  form:"fcorderh" `
	FCORDERI   string    `gorm:"column:FCORDERI;" json:"fcorderi"  form:"fcorderi" `
	FCORDERREF string    `gorm:"column:FCORDERREF;" json:"fcorderref"  form:"fcorderref" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCQUOTENO  string    `gorm:"column:FCQUOTENO;" json:"fcquoteno"  form:"fcquoteno" `
	FCQUOTREF  string    `gorm:"column:FCQUOTREF;" json:"fcquotref"  form:"fcquotref" `
	FCRECEIVE  string    `gorm:"column:FCRECEIVE;" json:"fcreceive"  form:"fcreceive" `
	FCRECSTAT  string    `gorm:"column:FCRECSTAT;" json:"fcrecstat"  form:"fcrecstat" `
	FCRECTYPE  string    `gorm:"column:FCRECTYPE;" json:"fcrectype"  form:"fcrectype" `
	FCSECTSOLD string    `gorm:"column:FCSECTSOLD;" json:"fcsectsold"  form:"fcsectsold" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCTOSECT   string    `gorm:"column:FCTOSECT;" json:"fctosect"  form:"fctosect" `
	FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCTRANTYPE string    `gorm:"column:FCTRANTYPE;" json:"fctrantype"  form:"fctrantype" `
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
	FCVINNO    string    `gorm:"column:FCVINNO;" json:"fcvinno"  form:"fcvinno" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FDDISTDATE string    `gorm:"column:FDDISTDATE;" json:"fddistdate"  form:"fddistdate" `
	FDDRVASGN  string    `gorm:"column:FDDRVASGN;" json:"fddrvasgn"  form:"fddrvasgn" `
	FDFINALASG string    `gorm:"column:FDFINALASG;" json:"fdfinalasg"  form:"fdfinalasg" `
	FDINVDATE  string    `gorm:"column:FDINVDATE;" json:"fdinvdate"  form:"fdinvdate" `
	FDKAKUDAI  string    `gorm:"column:FDKAKUDAI;" json:"fdkakudai"  form:"fdkakudai" `
	FDRECDATE  string    `gorm:"column:FDRECDATE;" json:"fdrecdate"  form:"fdrecdate" `
	FDRECEIVE  string    `gorm:"column:FDRECEIVE;" json:"fdreceive"  form:"fdreceive" `
	FDRECINV   string    `gorm:"column:FDRECINV;" json:"fdrecinv"  form:"fdrecinv" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FMMEMDATA2 string    `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
	FMMEMDATA3 string    `gorm:"column:FMMEMDATA3;" json:"fmmemdata3"  form:"fmmemdata3" `
	FMMEMDATA4 string    `gorm:"column:FMMEMDATA4;" json:"fmmemdata4"  form:"fmmemdata4" `
	FMMEMDATA5 string    `gorm:"column:FMMEMDATA5;" json:"fmmemdata5"  form:"fmmemdata5" `
	FNMILE     string    `gorm:"column:FNMILE;" json:"fnmile"  form:"fnmile" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vinno) TableName() string {
	return "VINNO"
}

func (obj *Vinno) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
