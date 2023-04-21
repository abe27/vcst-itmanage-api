package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Qreport struct {
	FCA1STATUS string    `gorm:"column:FCA1STATUS;" json:"fca1status"  form:"fca1status" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
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
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCOWNER    string    `gorm:"column:FCOWNER;" json:"fcowner"  form:"fcowner" `
	FCREPEXT   string    `gorm:"column:FCREPEXT;" json:"fcrepext"  form:"fcrepext" `
	FCREPNAME  string    `gorm:"column:FCREPNAME;" json:"fcrepname"  form:"fcrepname" `
	FCSDNAME   string    `gorm:"column:FCSDNAME;" json:"fcsdname"  form:"fcsdname" `
	FCSENDAUTO string    `gorm:"column:FCSENDAUTO;" json:"fcsendauto"  form:"fcsendauto" `
	FCSENDTO   string    `gorm:"column:FCSENDTO;" json:"fcsendto"  form:"fcsendto" `
	FCSID      string    `gorm:"column:FCSID;" json:"fcsid"  form:"fcsid" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSOURCE   string    `gorm:"column:FCSOURCE;" json:"fcsource"  form:"fcsource" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCSYSTASK  string    `gorm:"column:FCSYSTASK;" json:"fcsystask"  form:"fcsystask" `
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
	FMFPARA1   string    `gorm:"column:FMFPARA1;" json:"fmfpara1"  form:"fmfpara1" `
	FMFPARA2   string    `gorm:"column:FMFPARA2;" json:"fmfpara2"  form:"fmfpara2" `
	FMFPARA3   string    `gorm:"column:FMFPARA3;" json:"fmfpara3"  form:"fmfpara3" `
	FMFPARA4   string    `gorm:"column:FMFPARA4;" json:"fmfpara4"  form:"fmfpara4" `
	FMFPARA5   string    `gorm:"column:FMFPARA5;" json:"fmfpara5"  form:"fmfpara5" `
	FMWPARA1   string    `gorm:"column:FMWPARA1;" json:"fmwpara1"  form:"fmwpara1" `
	FMWPARA2   string    `gorm:"column:FMWPARA2;" json:"fmwpara2"  form:"fmwpara2" `
	FMWPARA3   string    `gorm:"column:FMWPARA3;" json:"fmwpara3"  form:"fmwpara3" `
	FMWPARA4   string    `gorm:"column:FMWPARA4;" json:"fmwpara4"  form:"fmwpara4" `
	FMWPARA5   string    `gorm:"column:FMWPARA5;" json:"fmwpara5"  form:"fmwpara5" `
	FNPRIORITY string    `gorm:"column:FNPRIORITY;" json:"fnpriority"  form:"fnpriority" `
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
	FTSTARTAF  string    `gorm:"column:FTSTARTAF;" json:"ftstartaf"  form:"ftstartaf" `
	QCCORP     string    `gorm:"column:QCCORP;" json:"qccorp"  form:"qccorp" `
	QNCORP     string    `gorm:"column:QNCORP;" json:"qncorp"  form:"qncorp" `
}

func (Qreport) TableName() string {
	return "QREPORT"
}

func (obj *Qreport) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
