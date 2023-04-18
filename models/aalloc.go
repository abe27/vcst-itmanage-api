package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Aalloc struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPSKID  string    `gorm:"column:FCAPPSKID;" json:"fcappskid"  form:"fcappskid" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCENCOTYPE string    `gorm:"column:FCENCOTYPE;" json:"fcencotype"  form:"fcencotype" `
	FCHALLOC   string    `gorm:"column:FCHALLOC;" json:"fchalloc"  form:"fchalloc" `
	FCHALLOC1  string    `gorm:"column:FCHALLOC1;" json:"fchalloc1"  form:"fchalloc1" `
	FCHCNT     string    `gorm:"column:FCHCNT;" json:"fchcnt"  form:"fchcnt" `
	FCHCTRLST1 string    `gorm:"column:FCHCTRLST1;" json:"fchctrlst1"  form:"fchctrlst1" `
	FCHCTRLST2 string    `gorm:"column:FCHCTRLST2;" json:"fchctrlst2"  form:"fchctrlst2" `
	FCHCTRLSTR string    `gorm:"column:FCHCTRLSTR;" json:"fchctrlstr"  form:"fchctrlstr" `
	FCHDISKSK1 string    `gorm:"column:FCHDISKSK1;" json:"fchdisksk1"  form:"fchdisksk1" `
	FCHDISKSKI string    `gorm:"column:FCHDISKSKI;" json:"fchdiskski"  form:"fchdiskski" `
	FCHDLIM    string    `gorm:"column:FCHDLIM;" json:"fchdlim"  form:"fchdlim" `
	FCHDLIM10  string    `gorm:"column:FCHDLIM10;" json:"fchdlim10"  form:"fchdlim10" `
	FCHDLIM11  string    `gorm:"column:FCHDLIM11;" json:"fchdlim11"  form:"fchdlim11" `
	FCHDLIM12  string    `gorm:"column:FCHDLIM12;" json:"fchdlim12"  form:"fchdlim12" `
	FCHDLIM13  string    `gorm:"column:FCHDLIM13;" json:"fchdlim13"  form:"fchdlim13" `
	FCHDLIM14  string    `gorm:"column:FCHDLIM14;" json:"fchdlim14"  form:"fchdlim14" `
	FCHDLIM15  string    `gorm:"column:FCHDLIM15;" json:"fchdlim15"  form:"fchdlim15" `
	FCHDLIM16  string    `gorm:"column:FCHDLIM16;" json:"fchdlim16"  form:"fchdlim16" `
	FCHDLIM17  string    `gorm:"column:FCHDLIM17;" json:"fchdlim17"  form:"fchdlim17" `
	FCHDLIM18  string    `gorm:"column:FCHDLIM18;" json:"fchdlim18"  form:"fchdlim18" `
	FCHDLIM19  string    `gorm:"column:FCHDLIM19;" json:"fchdlim19"  form:"fchdlim19" `
	FCHDLIM2   string    `gorm:"column:FCHDLIM2;" json:"fchdlim2"  form:"fchdlim2" `
	FCHDLIM20  string    `gorm:"column:FCHDLIM20;" json:"fchdlim20"  form:"fchdlim20" `
	FCHDLIM3   string    `gorm:"column:FCHDLIM3;" json:"fchdlim3"  form:"fchdlim3" `
	FCHDLIM4   string    `gorm:"column:FCHDLIM4;" json:"fchdlim4"  form:"fchdlim4" `
	FCHDLIM5   string    `gorm:"column:FCHDLIM5;" json:"fchdlim5"  form:"fchdlim5" `
	FCHDLIM6   string    `gorm:"column:FCHDLIM6;" json:"fchdlim6"  form:"fchdlim6" `
	FCHDLIM7   string    `gorm:"column:FCHDLIM7;" json:"fchdlim7"  form:"fchdlim7" `
	FCHDLIM8   string    `gorm:"column:FCHDLIM8;" json:"fchdlim8"  form:"fchdlim8" `
	FCHDLIM9   string    `gorm:"column:FCHDLIM9;" json:"fchdlim9"  form:"fchdlim9" `
	FCHPACKNAM string    `gorm:"column:FCHPACKNAM;" json:"fchpacknam"  form:"fchpacknam" `
	FCHPACKNM1 string    `gorm:"column:FCHPACKNM1;" json:"fchpacknm1"  form:"fchpacknm1" `
	FCHPACKNM2 string    `gorm:"column:FCHPACKNM2;" json:"fchpacknm2"  form:"fchpacknm2" `
	FCHPRODUC1 string    `gorm:"column:FCHPRODUC1;" json:"fchproduc1"  form:"fchproduc1" `
	FCHPRODUC2 string    `gorm:"column:FCHPRODUC2;" json:"fchproduc2"  form:"fchproduc2" `
	FCHPRODUCT string    `gorm:"column:FCHPRODUCT;" json:"fchproduct"  form:"fchproduct" `
	FCHRLIM    string    `gorm:"column:FCHRLIM;" json:"fchrlim"  form:"fchrlim" `
	FCHRLIM10  string    `gorm:"column:FCHRLIM10;" json:"fchrlim10"  form:"fchrlim10" `
	FCHRLIM11  string    `gorm:"column:FCHRLIM11;" json:"fchrlim11"  form:"fchrlim11" `
	FCHRLIM12  string    `gorm:"column:FCHRLIM12;" json:"fchrlim12"  form:"fchrlim12" `
	FCHRLIM13  string    `gorm:"column:FCHRLIM13;" json:"fchrlim13"  form:"fchrlim13" `
	FCHRLIM14  string    `gorm:"column:FCHRLIM14;" json:"fchrlim14"  form:"fchrlim14" `
	FCHRLIM15  string    `gorm:"column:FCHRLIM15;" json:"fchrlim15"  form:"fchrlim15" `
	FCHRLIM16  string    `gorm:"column:FCHRLIM16;" json:"fchrlim16"  form:"fchrlim16" `
	FCHRLIM17  string    `gorm:"column:FCHRLIM17;" json:"fchrlim17"  form:"fchrlim17" `
	FCHRLIM18  string    `gorm:"column:FCHRLIM18;" json:"fchrlim18"  form:"fchrlim18" `
	FCHRLIM19  string    `gorm:"column:FCHRLIM19;" json:"fchrlim19"  form:"fchrlim19" `
	FCHRLIM2   string    `gorm:"column:FCHRLIM2;" json:"fchrlim2"  form:"fchrlim2" `
	FCHRLIM20  string    `gorm:"column:FCHRLIM20;" json:"fchrlim20"  form:"fchrlim20" `
	FCHRLIM3   string    `gorm:"column:FCHRLIM3;" json:"fchrlim3"  form:"fchrlim3" `
	FCHRLIM4   string    `gorm:"column:FCHRLIM4;" json:"fchrlim4"  form:"fchrlim4" `
	FCHRLIM5   string    `gorm:"column:FCHRLIM5;" json:"fchrlim5"  form:"fchrlim5" `
	FCHRLIM6   string    `gorm:"column:FCHRLIM6;" json:"fchrlim6"  form:"fchrlim6" `
	FCHRLIM7   string    `gorm:"column:FCHRLIM7;" json:"fchrlim7"  form:"fchrlim7" `
	FCHRLIM8   string    `gorm:"column:FCHRLIM8;" json:"fchrlim8"  form:"fchrlim8" `
	FCHRLIM9   string    `gorm:"column:FCHRLIM9;" json:"fchrlim9"  form:"fchrlim9" `
	FCHSERIAL  string    `gorm:"column:FCHSERIAL;" json:"fchserial"  form:"fchserial" `
	FCHSERIAL1 string    `gorm:"column:FCHSERIAL1;" json:"fchserial1"  form:"fchserial1" `
	FCHSUBSYS1 string    `gorm:"column:FCHSUBSYS1;" json:"fchsubsys1"  form:"fchsubsys1" `
	FCHSUBSYS2 string    `gorm:"column:FCHSUBSYS2;" json:"fchsubsys2"  form:"fchsubsys2" `
	FCHSUBSYST string    `gorm:"column:FCHSUBSYST;" json:"fchsubsyst"  form:"fchsubsyst" `
	FCHTSTP    string    `gorm:"column:FCHTSTP;" json:"fchtstp"  form:"fchtstp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSERIAL   string    `gorm:"column:FCSERIAL;" json:"fcserial"  form:"fcserial" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
}

func (Aalloc) TableName() string {
	return "AALLOC"
}

func (obj *Aalloc) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
