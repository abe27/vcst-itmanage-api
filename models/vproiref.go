package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vproiref struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPACKSTYL string    `gorm:"column:FCPACKSTYL;" json:"fcpackstyl"  form:"fcpackstyl" `
	FCPDSTHRPK string    `gorm:"column:FCPDSTHRPK;" json:"fcpdsthrpk"  form:"fcpdsthrpk" `
	FCPDSTHRWK string    `gorm:"column:FCPDSTHRWK;" json:"fcpdsthrwk"  form:"fcpdsthrwk" `
	FCPDSTHSM  string    `gorm:"column:FCPDSTHSM;" json:"fcpdsthsm"  form:"fcpdsthsm" `
	FCPRODSM   string    `gorm:"column:FCPRODSM;" json:"fcprodsm"  form:"fcprodsm" `
	FCPRODWR   string    `gorm:"column:FCPRODWR;" json:"fcprodwr"  form:"fcprodwr" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTEP2    string    `gorm:"column:FCSTEP2;" json:"fcstep2"  form:"fcstep2" `
	FCSTEPPO   string    `gorm:"column:FCSTEPPO;" json:"fcsteppo"  form:"fcsteppo" `
	FCTABNAME  string    `gorm:"column:FCTABNAME;" json:"fctabname"  form:"fctabname" `
	FCTABSKID  string    `gorm:"column:FCTABSKID;" json:"fctabskid"  form:"fctabskid" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDFGDUE    string    `gorm:"column:FDFGDUE;" json:"fdfgdue"  form:"fdfgdue" `
	FDPACKDUE  string    `gorm:"column:FDPACKDUE;" json:"fdpackdue"  form:"fdpackdue" `
	FDSEMIDUE  string    `gorm:"column:FDSEMIDUE;" json:"fdsemidue"  form:"fdsemidue" `
	FDSTEP     string    `gorm:"column:FDSTEP;" json:"fdstep"  form:"fdstep" `
	FDSTEP2    string    `gorm:"column:FDSTEP2;" json:"fdstep2"  form:"fdstep2" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FMREMARK01 string    `gorm:"column:FMREMARK01;" json:"fmremark01"  form:"fmremark01" `
	FMREMARK02 string    `gorm:"column:FMREMARK02;" json:"fmremark02"  form:"fmremark02" `
	FMREMARK03 string    `gorm:"column:FMREMARK03;" json:"fmremark03"  form:"fmremark03" `
	FMREMARK04 string    `gorm:"column:FMREMARK04;" json:"fmremark04"  form:"fmremark04" `
	FMREMARK05 string    `gorm:"column:FMREMARK05;" json:"fmremark05"  form:"fmremark05" `
	FMREMARK10 string    `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
	FMREMARK2  string    `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FMREMARK3  string    `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
	FMREMARK4  string    `gorm:"column:FMREMARK4;" json:"fmremark4"  form:"fmremark4" `
	FMREMARK5  string    `gorm:"column:FMREMARK5;" json:"fmremark5"  form:"fmremark5" `
	FMREMARK6  string    `gorm:"column:FMREMARK6;" json:"fmremark6"  form:"fmremark6" `
	FMREMARK7  string    `gorm:"column:FMREMARK7;" json:"fmremark7"  form:"fmremark7" `
	FMREMARK8  string    `gorm:"column:FMREMARK8;" json:"fmremark8"  form:"fmremark8" `
	FMREMARK9  string    `gorm:"column:FMREMARK9;" json:"fmremark9"  form:"fmremark9" `
	FNFGALLOC  string    `gorm:"column:FNFGALLOC;" json:"fnfgalloc"  form:"fnfgalloc" `
	FNFGPCN    string    `gorm:"column:FNFGPCN;" json:"fnfgpcn"  form:"fnfgpcn" `
	FNFGQTY    string    `gorm:"column:FNFGQTY;" json:"fnfgqty"  form:"fnfgqty" `
	FNFGUW     string    `gorm:"column:FNFGUW;" json:"fnfguw"  form:"fnfguw" `
	FNNFGQTY   string    `gorm:"column:FNNFGQTY;" json:"fnnfgqty"  form:"fnnfgqty" `
	FNNPACKQTY string    `gorm:"column:FNNPACKQTY;" json:"fnnpackqty"  form:"fnnpackqty" `
	FNNSEMIQTY string    `gorm:"column:FNNSEMIQTY;" json:"fnnsemiqty"  form:"fnnsemiqty" `
	FNPACKPCN  string    `gorm:"column:FNPACKPCN;" json:"fnpackpcn"  form:"fnpackpcn" `
	FNPACKQTY  string    `gorm:"column:FNPACKQTY;" json:"fnpackqty"  form:"fnpackqty" `
	FNPACKUW   string    `gorm:"column:FNPACKUW;" json:"fnpackuw"  form:"fnpackuw" `
	FNREPKALLO string    `gorm:"column:FNREPKALLO;" json:"fnrepkallo"  form:"fnrepkallo" `
	FNREWKALLO string    `gorm:"column:FNREWKALLO;" json:"fnrewkallo"  form:"fnrewkallo" `
	FNSEMIALLO string    `gorm:"column:FNSEMIALLO;" json:"fnsemiallo"  form:"fnsemiallo" `
	FNSEMIPCN  string    `gorm:"column:FNSEMIPCN;" json:"fnsemipcn"  form:"fnsemipcn" `
	FNSEMIQTY  string    `gorm:"column:FNSEMIQTY;" json:"fnsemiqty"  form:"fnsemiqty" `
	FNSEMIUW   string    `gorm:"column:FNSEMIUW;" json:"fnsemiuw"  form:"fnsemiuw" `
	FNSMRWALLO string    `gorm:"column:FNSMRWALLO;" json:"fnsmrwallo"  form:"fnsmrwallo" `
	FNTXIDALLO string    `gorm:"column:FNTXIDALLO;" json:"fntxidallo"  form:"fntxidallo" `
	FNYEILDPCN string    `gorm:"column:FNYEILDPCN;" json:"fnyeildpcn"  form:"fnyeildpcn" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vproiref) TableName() string {
	return "VPROIREF"
}

func (obj *Vproiref) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
