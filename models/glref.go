package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Glref struct {
	FCATSTEP   string    `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCBOOK     string    `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" default:"$$$+"`
	FCDELICOOR string    `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCFRWHOUSE string    `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"`
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" default:"I"`
	FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCVATCOOR  string    `gorm:"column:FCVATCOOR;" json:"fcvatcoor"  form:"fcvatcoor" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNAFTDEP   float64   `gorm:"column:FNAFTDEP;" json:"fnaftdep"  form:"fnaftdep" defualt:"0"`
	FNAFTDEPKE float64   `gorm:"column:FNAFTDEPKE;" json:"fnaftdepke"  form:"fnaftdepke" defualt:"0"`
	FNAMT      float64   `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
}

func (Glref) TableName() string {
	return "GLREF"
}

func (obj *Glref) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("G%sF", id)

	uid, _ := g.New(26)
	obj.FCGID = uid
	obj.FIMILLISEC = time.Now().Unix()
	obj.FCLUPDAPP = "$0"
	return
}

type GlrefTable struct {
	FCATSTEP       string    `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCBOOK         string    `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH       string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE         string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR         string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP         string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB     string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEBY     string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY     string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCDATASER      string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" default:"$$$+"`
	FCDELICOOR     string    `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	FCDEPT         string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR      string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCFRWHOUSE     string    `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCGID          string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD       string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCJOB          string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCPROJ         string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFNO        string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE      string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE       string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT         string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSKID         string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"`
	FCSTEP         string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" default:"I"`
	FCTOWHOUSE     string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCVATCOOR      string    `gorm:"column:FCVATCOOR;" json:"fcvatcoor"  form:"fcvatcoor" `
	FDDATE         time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FIMILLISEC     int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNAFTDEP       float64   `gorm:"column:FNAFTDEP;" json:"fnaftdep"  form:"fnaftdep" defualt:"0"`
	FNAFTDEPKE     float64   `gorm:"column:FNAFTDEPKE;" json:"fnaftdepke"  form:"fnaftdepke" defualt:"0"`
	FNAMT          float64   `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FTDATETIME     time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT     time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD      time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FCLUPDAPP      string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FMMEMDATA      string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FCSTATUS       bool      `json:"fcstatus" default:"false"`
	Corp           *Corp     `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	Branch         *Branch   `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	Dept           *Dept     `gorm:"foreignKey:FCDEPT;references:FCSKID;" json:"department"`
	Sect           *Sect     `gorm:"foreignKey:FCSECT;references:FCSKID;" json:"section"`
	Job            *Job      `gorm:"foreignKey:FCJOB;references:FCSKID;" json:"job"`
	Glhead         *Glhead   `gorm:"foreignKey:FCGLHEAD;references:FCSKID;" json:"glhead"`
	Book           *Book     `gorm:"foreignKey:FCBOOK;references:FCSKID;" json:"book"`
	Coor           *Coor     `gorm:"foreignKey:FCCOOR;reference:FCSKID;" json:"coor"`
	FromWhouse     *WHouse   `gorm:"foreignKey:FCFRWHOUSE;reference:FCSKID;" json:"from"`
	ToWhouse       *WHouse   `gorm:"foreignKey:FCTOWHOUSE;reference:FCSKID;" json:"to"`
	CreatedBy      *Empl     `gorm:"foreignKey:FCCREATEBY;references:FCSKID;" json:"created_by"`
	UpdatedBy      *Empl     `gorm:"foreignKey:FCCORRECTB;references:FCSKID;" json:"updated_by"`
	VatCoor        *Coor     `gorm:"foreignKey:FCVATCOOR;reference:FCSKID;" json:"vat_coor"`
	Proj           *Proj     `gorm:"foreignKey:FCPROJ;reference:FCSKID;" json:"proj"`
	DeliveryToCoor *Coor     `gorm:"foreignKey:FCDELICOOR;reference:FCSKID;" json:"delivery_to_coor"`
}

func (GlrefTable) TableName() string {
	return "GLREF"
}

type GlRefForm struct {
	FCBOOK     string       `json:"fcbook" form:"fcbook"`
	FCBRANCH   string       `json:"fcbranch" form:"fcbranch"`
	FROMWHOUSE string       `json:"fromwhouse" form:"fromwhouse"`
	TOWHOUSE   string       `json:"towhouse" form:"towhouse"`
	FCCODE     string       `json:"fccode"  form:"fccode" `
	FCGID      string       `json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string       `json:"fcglhead"  form:"fcglhead" `
	FCJOB      string       `json:"fcjob"  form:"fcjob" `
	FCPROJ     string       `json:"fcproj"  form:"fcproj" `
	FCREFNO    string       `json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string       `json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string       `json:"fcrftype"  form:"fcrftype" `
	FCSTEP     string       `json:"fcstep"  form:"fcstep" default:"I"`
	FCSEC      string       `json:"fcsec" form:"fcsec"`
	FCREMARK   string       `json:"fcremark" form:"fcremark"`
	FDDATE     time.Time    `json:"fddate"  form:"fddate" default:"now"`
	FNAFTDEP   float64      `json:"fnaftdep"  form:"fnaftdep" defualt:"0"`
	FNAFTDEPKE float64      `json:"fnaftdepke"  form:"fnaftdepke" defualt:"0"`
	REFPROD    []FrmRefProd `json:"ref_prod" form:"ref_prod"`
}

type FrmUpdateRefProd struct {
	FCPROD  string  `json:"fcprod"`
	FNQTY   float64 `json:"fnqty"`
	REFPROD string  `json:"refprod"`
}

type GlrefPatchForm struct {
	ORDERH  string             `json:"orderh"`
	REMARK  string             `json:"remark"`
	REFPROD []FrmUpdateRefProd `json:"refprod"`
}

type GlrefHistory struct {
	ID           string    `gorm:"size:21;primaryKey;not null;" json:"id"`
	GLREF        string    `gorm:"size:8;" json:"glref"`
	REFPROD      string    `gorm:"size:8;" json:"refprod"`
	ORDERH       string    `gorm:"size:8;" json:"orderh"`
	ORDERI       string    `gorm:"size:8;" json:"orderi"`
	PRODID       string    `gorm:"size:8;" json:"prodid"`
	REFNO        string    `gorm:"size:25;" json:"refno"`
	PONO         string    `gorm:"size:20;" json:"pono"`
	PRODUCTCODE  string    `gorm:"size:25;" json:"productcode"`
	PRODUCTNAME  string    `gorm:"size:50;" json:"productname"`
	QTY          float64   `json:"qty"`
	RECEIVEQTY   float64   `json:"receiveqty" default:"0"`
	Remark       string    `json:"remark"`
	IsComplete   bool      `json:"is_complete" default:"false"`
	UpdateByID   string    `json:"update_by_id"`
	UpdateByUser *User     `gorm:"foreignKey:UpdateByID;reference:ID;" json:"update_by"`
	CreatedAt    time.Time `json:"created_at" default:"now"`
	UpdatedAt    time.Time `json:"updated_at" default:"now"`
}

func (obj *GlrefHistory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	obj.ID = id
	return
}
