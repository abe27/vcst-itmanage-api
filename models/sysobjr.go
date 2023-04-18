package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Sysobjr struct {
	FCACCKEYS  string    `gorm:"column:FCACCKEYS;" json:"fcacckeys"  form:"fcacckeys" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCHEADER   string    `gorm:"column:FCHEADER;" json:"fcheader"  form:"fcheader" `
	FCISODOC   string    `gorm:"column:FCISODOC;" json:"fcisodoc"  form:"fcisodoc" `
	FCISOISSUE string    `gorm:"column:FCISOISSUE;" json:"fcisoissue"  form:"fcisoissue" `
	FCISOREVI  string    `gorm:"column:FCISOREVI;" json:"fcisorevi"  form:"fcisorevi" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPROMPT1  string    `gorm:"column:FCPROMPT1;" json:"fcprompt1"  form:"fcprompt1" `
	FCPROMPT2  string    `gorm:"column:FCPROMPT2;" json:"fcprompt2"  form:"fcprompt2" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSYSOBJ   string    `gorm:"column:FCSYSOBJ;" json:"fcsysobj"  form:"fcsysobj" `
	FCSYSTASK  string    `gorm:"column:FCSYSTASK;" json:"fcsystask"  form:"fcsystask" `
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
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMSG1     string    `gorm:"column:FMMSG1;" json:"fmmsg1"  form:"fmmsg1" `
	FMMSG2     string    `gorm:"column:FMMSG2;" json:"fmmsg2"  form:"fmmsg2" `
	FNACCCANCE float64   `gorm:"column:FNACCCANCE;" json:"fnacccance"  form:"fnacccance" `
	FNACCDELE  float64   `gorm:"column:FNACCDELE;" json:"fnaccdele"  form:"fnaccdele" `
	FNACCEDIT  float64   `gorm:"column:FNACCEDIT;" json:"fnaccedit"  form:"fnaccedit" `
	FNACCENTER float64   `gorm:"column:FNACCENTER;" json:"fnaccenter"  form:"fnaccenter" `
	FNACCEXIT  float64   `gorm:"column:FNACCEXIT;" json:"fnaccexit"  form:"fnaccexit" `
	FNACCFILE  float64   `gorm:"column:FNACCFILE;" json:"fnaccfile"  form:"fnaccfile" `
	FNACCINSER float64   `gorm:"column:FNACCINSER;" json:"fnaccinser"  form:"fnaccinser" `
	FNACCPRINT float64   `gorm:"column:FNACCPRINT;" json:"fnaccprint"  form:"fnaccprint" `
	FNACCPWREQ float64   `gorm:"column:FNACCPWREQ;" json:"fnaccpwreq"  form:"fnaccpwreq" `
	FNACCRIGSE float64   `gorm:"column:FNACCRIGSE;" json:"fnaccrigse"  form:"fnaccrigse" `
	FNACCSHOW  float64   `gorm:"column:FNACCSHOW;" json:"fnaccshow"  form:"fnaccshow" `
	FNACCVIEW  float64   `gorm:"column:FNACCVIEW;" json:"fnaccview"  form:"fnaccview" `
	FNCOMCANCE float64   `gorm:"column:FNCOMCANCE;" json:"fncomcance"  form:"fncomcance" `
	FNCOMDELE  float64   `gorm:"column:FNCOMDELE;" json:"fncomdele"  form:"fncomdele" `
	FNCOMEDIT  float64   `gorm:"column:FNCOMEDIT;" json:"fncomedit"  form:"fncomedit" `
	FNCOMENTER float64   `gorm:"column:FNCOMENTER;" json:"fncomenter"  form:"fncomenter" `
	FNCOMEXIT  float64   `gorm:"column:FNCOMEXIT;" json:"fncomexit"  form:"fncomexit" `
	FNCOMFILE  float64   `gorm:"column:FNCOMFILE;" json:"fncomfile"  form:"fncomfile" `
	FNCOMINSER float64   `gorm:"column:FNCOMINSER;" json:"fncominser"  form:"fncominser" `
	FNCOMPRINT float64   `gorm:"column:FNCOMPRINT;" json:"fncomprint"  form:"fncomprint" `
	FNCOMPWREQ float64   `gorm:"column:FNCOMPWREQ;" json:"fncompwreq"  form:"fncompwreq" `
	FNCOMRIGSE float64   `gorm:"column:FNCOMRIGSE;" json:"fncomrigse"  form:"fncomrigse" `
	FNCOMSHOW  float64   `gorm:"column:FNCOMSHOW;" json:"fncomshow"  form:"fncomshow" `
	FNCOMVIEW  float64   `gorm:"column:FNCOMVIEW;" json:"fncomview"  form:"fncomview" `
	FNCTCANCEL float64   `gorm:"column:FNCTCANCEL;" json:"fnctcancel"  form:"fnctcancel" `
	FNCTDELE   float64   `gorm:"column:FNCTDELE;" json:"fnctdele"  form:"fnctdele" `
	FNCTEDIT   float64   `gorm:"column:FNCTEDIT;" json:"fnctedit"  form:"fnctedit" `
	FNCTENTER  float64   `gorm:"column:FNCTENTER;" json:"fnctenter"  form:"fnctenter" `
	FNCTEXIT   float64   `gorm:"column:FNCTEXIT;" json:"fnctexit"  form:"fnctexit" `
	FNCTFILE   float64   `gorm:"column:FNCTFILE;" json:"fnctfile"  form:"fnctfile" `
	FNCTINSERT float64   `gorm:"column:FNCTINSERT;" json:"fnctinsert"  form:"fnctinsert" `
	FNCTPRINT  float64   `gorm:"column:FNCTPRINT;" json:"fnctprint"  form:"fnctprint" `
	FNCTPWREQ  float64   `gorm:"column:FNCTPWREQ;" json:"fnctpwreq"  form:"fnctpwreq" `
	FNCTRIGSET float64   `gorm:"column:FNCTRIGSET;" json:"fnctrigset"  form:"fnctrigset" `
	FNCTSHOW   float64   `gorm:"column:FNCTSHOW;" json:"fnctshow"  form:"fnctshow" `
	FNCTVIEW   float64   `gorm:"column:FNCTVIEW;" json:"fnctview"  form:"fnctview" `
	FNMDCANCEL float64   `gorm:"column:FNMDCANCEL;" json:"fnmdcancel"  form:"fnmdcancel" `
	FNMDDELE   float64   `gorm:"column:FNMDDELE;" json:"fnmddele"  form:"fnmddele" `
	FNMDEDIT   float64   `gorm:"column:FNMDEDIT;" json:"fnmdedit"  form:"fnmdedit" `
	FNMDENTER  float64   `gorm:"column:FNMDENTER;" json:"fnmdenter"  form:"fnmdenter" `
	FNMDEXIT   float64   `gorm:"column:FNMDEXIT;" json:"fnmdexit"  form:"fnmdexit" `
	FNMDFILE   float64   `gorm:"column:FNMDFILE;" json:"fnmdfile"  form:"fnmdfile" `
	FNMDINSERT float64   `gorm:"column:FNMDINSERT;" json:"fnmdinsert"  form:"fnmdinsert" `
	FNMDPRINT  float64   `gorm:"column:FNMDPRINT;" json:"fnmdprint"  form:"fnmdprint" `
	FNMDPWREQ  float64   `gorm:"column:FNMDPWREQ;" json:"fnmdpwreq"  form:"fnmdpwreq" `
	FNMDRIGSET float64   `gorm:"column:FNMDRIGSET;" json:"fnmdrigset"  form:"fnmdrigset" `
	FNMDSHOW   float64   `gorm:"column:FNMDSHOW;" json:"fnmdshow"  form:"fnmdshow" `
	FNMDVIEW   float64   `gorm:"column:FNMDVIEW;" json:"fnmdview"  form:"fnmdview" `
	FNMNGCANCE float64   `gorm:"column:FNMNGCANCE;" json:"fnmngcance"  form:"fnmngcance" `
	FNMNGDELE  float64   `gorm:"column:FNMNGDELE;" json:"fnmngdele"  form:"fnmngdele" `
	FNMNGEDIT  float64   `gorm:"column:FNMNGEDIT;" json:"fnmngedit"  form:"fnmngedit" `
	FNMNGENTER float64   `gorm:"column:FNMNGENTER;" json:"fnmngenter"  form:"fnmngenter" `
	FNMNGEXIT  float64   `gorm:"column:FNMNGEXIT;" json:"fnmngexit"  form:"fnmngexit" `
	FNMNGFILE  float64   `gorm:"column:FNMNGFILE;" json:"fnmngfile"  form:"fnmngfile" `
	FNMNGINSER float64   `gorm:"column:FNMNGINSER;" json:"fnmnginser"  form:"fnmnginser" `
	FNMNGPRINT float64   `gorm:"column:FNMNGPRINT;" json:"fnmngprint"  form:"fnmngprint" `
	FNMNGPWREQ float64   `gorm:"column:FNMNGPWREQ;" json:"fnmngpwreq"  form:"fnmngpwreq" `
	FNMNGRIGSE float64   `gorm:"column:FNMNGRIGSE;" json:"fnmngrigse"  form:"fnmngrigse" `
	FNMNGSHOW  float64   `gorm:"column:FNMNGSHOW;" json:"fnmngshow"  form:"fnmngshow" `
	FNMNGVIEW  float64   `gorm:"column:FNMNGVIEW;" json:"fnmngview"  form:"fnmngview" `
	FNOPCANCEL float64   `gorm:"column:FNOPCANCEL;" json:"fnopcancel"  form:"fnopcancel" `
	FNOPDELE   float64   `gorm:"column:FNOPDELE;" json:"fnopdele"  form:"fnopdele" `
	FNOPEDIT   float64   `gorm:"column:FNOPEDIT;" json:"fnopedit"  form:"fnopedit" `
	FNOPENTER  float64   `gorm:"column:FNOPENTER;" json:"fnopenter"  form:"fnopenter" `
	FNOPEXIT   float64   `gorm:"column:FNOPEXIT;" json:"fnopexit"  form:"fnopexit" `
	FNOPFILE   float64   `gorm:"column:FNOPFILE;" json:"fnopfile"  form:"fnopfile" `
	FNOPINSERT float64   `gorm:"column:FNOPINSERT;" json:"fnopinsert"  form:"fnopinsert" `
	FNOPPRINT  float64   `gorm:"column:FNOPPRINT;" json:"fnopprint"  form:"fnopprint" `
	FNOPPWREQ  float64   `gorm:"column:FNOPPWREQ;" json:"fnoppwreq"  form:"fnoppwreq" `
	FNOPRIGSET float64   `gorm:"column:FNOPRIGSET;" json:"fnoprigset"  form:"fnoprigset" `
	FNOPSHOW   float64   `gorm:"column:FNOPSHOW;" json:"fnopshow"  form:"fnopshow" `
	FNOPVIEW   float64   `gorm:"column:FNOPVIEW;" json:"fnopview"  form:"fnopview" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Sysobjr) TableName() string {
	return "SYSOBJR"
}

func (obj *Sysobjr) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
