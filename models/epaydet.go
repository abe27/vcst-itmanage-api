package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Epaydet struct {
	FCADVICEBY string    `gorm:"column:FCADVICEBY;" json:"fcadviceby"  form:"fcadviceby" `
	FCADVICETO string    `gorm:"column:FCADVICETO;" json:"fcadviceto"  form:"fcadviceto" `
	FCBATCHID  string    `gorm:"column:FCBATCHID;" json:"fcbatchid"  form:"fcbatchid" `
	FCBICCODE  string    `gorm:"column:FCBICCODE;" json:"fcbiccode"  form:"fcbiccode" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCHEQTYP  string    `gorm:"column:FCCHEQTYP;" json:"fccheqtyp"  form:"fccheqtyp" `
	FCCHGBEAR  string    `gorm:"column:FCCHGBEAR;" json:"fcchgbear"  form:"fcchgbear" `
	FCCLEARZON string    `gorm:"column:FCCLEARZON;" json:"fcclearzon"  form:"fcclearzon" `
	FCCOBANK   string    `gorm:"column:FCCOBANK;" json:"fccobank"  form:"fccobank" `
	FCCOBKACCT string    `gorm:"column:FCCOBKACCT;" json:"fccobkacct"  form:"fccobkacct" `
	FCCOBRANCH string    `gorm:"column:FCCOBRANCH;" json:"fccobranch"  form:"fccobranch" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOTAXID  string    `gorm:"column:FCCOTAXID;" json:"fccotaxid"  form:"fccotaxid" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELIBY   string    `gorm:"column:FCDELIBY;" json:"fcdeliby"  form:"fcdeliby" `
	FCDOC2PICK string    `gorm:"column:FCDOC2PICK;" json:"fcdoc2pick"  form:"fcdoc2pick" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLOCALINS string    `gorm:"column:FCLOCALINS;" json:"fclocalins"  form:"fclocalins" `
	FCNCHRGBR  string    `gorm:"column:FCNCHRGBR;" json:"fcnchrgbr"  form:"fcnchrgbr" `
	FCNPYMETHD string    `gorm:"column:FCNPYMETHD;" json:"fcnpymethd"  form:"fcnpymethd" `
	FCP2ERRMSG string    `gorm:"column:FCP2ERRMSG;" json:"fcp2errmsg"  form:"fcp2errmsg" `
	FCP2ERRNO  string    `gorm:"column:FCP2ERRNO;" json:"fcp2errno"  form:"fcp2errno" `
	FCP2STEP   string    `gorm:"column:FCP2STEP;" json:"fcp2step"  form:"fcp2step" `
	FCPAYID    string    `gorm:"column:FCPAYID;" json:"fcpayid"  form:"fcpayid" `
	FCPAYMENT  string    `gorm:"column:FCPAYMENT;" json:"fcpayment"  form:"fcpayment" `
	FCPICKLOCA string    `gorm:"column:FCPICKLOCA;" json:"fcpickloca"  form:"fcpickloca" `
	FCPRPAYTAX string    `gorm:"column:FCPRPAYTAX;" json:"fcprpaytax"  form:"fcprpaytax" `
	FCPYMETHOD string    `gorm:"column:FCPYMETHOD;" json:"fcpymethod"  form:"fcpymethod" `
	FCRECID    string    `gorm:"column:FCRECID;" json:"fcrecid"  form:"fcrecid" `
	FCREFNO1   string    `gorm:"column:FCREFNO1;" json:"fcrefno1"  form:"fcrefno1" `
	FCREFNO2   string    `gorm:"column:FCREFNO2;" json:"fcrefno2"  form:"fcrefno2" `
	FCSERVTYPE string    `gorm:"column:FCSERVTYPE;" json:"fcservtype"  form:"fcservtype" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCSUCCFILE string    `gorm:"column:FCSUCCFILE;" json:"fcsuccfile"  form:"fcsuccfile" `
	FCSVLEVEL  string    `gorm:"column:FCSVLEVEL;" json:"fcsvlevel"  form:"fcsvlevel" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FDP2DATE   string    `gorm:"column:FDP2DATE;" json:"fdp2date"  form:"fdp2date" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Epaydet) TableName() string {
	return "EPAYDET"
}

func (obj *Epaydet) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
