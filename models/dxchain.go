package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Dxchain struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCHAINBR  string    `gorm:"column:FCCHAINBR;" json:"fcchainbr"  form:"fcchainbr" `
	FCCHAINCP  string    `gorm:"column:FCCHAINCP;" json:"fcchaincp"  form:"fcchaincp" `
	FCCHAINDB  string    `gorm:"column:FCCHAINDB;" json:"fcchaindb"  form:"fcchaindb" `
	FCCHAINPSQ string    `gorm:"column:FCCHAINPSQ;" json:"fcchainpsq"  form:"fcchainpsq" `
	FCCHAINSEQ string    `gorm:"column:FCCHAINSEQ;" json:"fcchainseq"  form:"fcchainseq" `
	FCCHAINTAB string    `gorm:"column:FCCHAINTAB;" json:"fcchaintab"  form:"fcchaintab" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDXBRANCH string    `gorm:"column:FCDXBRANCH;" json:"fcdxbranch"  form:"fcdxbranch" `
	FCDXCODE   string    `gorm:"column:FCDXCODE;" json:"fcdxcode"  form:"fcdxcode" `
	FCDXCORP   string    `gorm:"column:FCDXCORP;" json:"fcdxcorp"  form:"fcdxcorp" `
	FCDXCTWITH string    `gorm:"column:FCDXCTWITH;" json:"fcdxctwith"  form:"fcdxctwith" `
	FCDXQCBOOK string    `gorm:"column:FCDXQCBOOK;" json:"fcdxqcbook"  form:"fcdxqcbook" `
	FCDXREFNO  string    `gorm:"column:FCDXREFNO;" json:"fcdxrefno"  form:"fcdxrefno" `
	FCDXREFTYP string    `gorm:"column:FCDXREFTYP;" json:"fcdxreftyp"  form:"fcdxreftyp" `
	FCDXSTAT   string    `gorm:"column:FCDXSTAT;" json:"fcdxstat"  form:"fcdxstat" `
	FCDXSTEP   string    `gorm:"column:FCDXSTEP;" json:"fcdxstep"  form:"fcdxstep" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCRUNNING  string    `gorm:"column:FCRUNNING;" json:"fcrunning"  form:"fcrunning" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDDXDATE   string    `gorm:"column:FDDXDATE;" json:"fddxdate"  form:"fddxdate" `
	FIMILLISEC string    `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Dxchain) TableName() string {
	return "DXCHAIN"
}

func (obj *Dxchain) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
