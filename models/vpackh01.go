package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Vpackh01 struct {
	FCA1_14    string    `gorm:"column:FCA1_14;" json:"fca1_14"  form:"fca1_14" `
	FCA1_15    string    `gorm:"column:FCA1_15;" json:"fca1_15"  form:"fca1_15" `
	FCA1_16    string    `gorm:"column:FCA1_16;" json:"fca1_16"  form:"fca1_16" `
	FCA1_17    string    `gorm:"column:FCA1_17;" json:"fca1_17"  form:"fca1_17" `
	FCA1_18    string    `gorm:"column:FCA1_18;" json:"fca1_18"  form:"fca1_18" `
	FCA1_19    string    `gorm:"column:FCA1_19;" json:"fca1_19"  form:"fca1_19" `
	FCA1_20    string    `gorm:"column:FCA1_20;" json:"fca1_20"  form:"fca1_20" `
	FCA1_21    string    `gorm:"column:FCA1_21;" json:"fca1_21"  form:"fca1_21" `
	FCA1_22    string    `gorm:"column:FCA1_22;" json:"fca1_22"  form:"fca1_22" `
	FCA1_23    string    `gorm:"column:FCA1_23;" json:"fca1_23"  form:"fca1_23" `
	FCA1_24    string    `gorm:"column:FCA1_24;" json:"fca1_24"  form:"fca1_24" `
	FCA1_25    string    `gorm:"column:FCA1_25;" json:"fca1_25"  form:"fca1_25" `
	FCA1_26    string    `gorm:"column:FCA1_26;" json:"fca1_26"  form:"fca1_26" `
	FCA1_27    string    `gorm:"column:FCA1_27;" json:"fca1_27"  form:"fca1_27" `
	FCA1_28    string    `gorm:"column:FCA1_28;" json:"fca1_28"  form:"fca1_28" `
	FCA1_29    string    `gorm:"column:FCA1_29;" json:"fca1_29"  form:"fca1_29" `
	FCA1_30    string    `gorm:"column:FCA1_30;" json:"fca1_30"  form:"fca1_30" `
	FCA1_31    string    `gorm:"column:FCA1_31;" json:"fca1_31"  form:"fca1_31" `
	FCA1_32    string    `gorm:"column:FCA1_32;" json:"fca1_32"  form:"fca1_32" `
	FCA1_34    string    `gorm:"column:FCA1_34;" json:"fca1_34"  form:"fca1_34" `
	FCA1_37    string    `gorm:"column:FCA1_37;" json:"fca1_37"  form:"fca1_37" `
	FCA1_38    string    `gorm:"column:FCA1_38;" json:"fca1_38"  form:"fca1_38" `
	FCA1_39    string    `gorm:"column:FCA1_39;" json:"fca1_39"  form:"fca1_39" `
	FCA1_40    string    `gorm:"column:FCA1_40;" json:"fca1_40"  form:"fca1_40" `
	FCA1_41    string    `gorm:"column:FCA1_41;" json:"fca1_41"  form:"fca1_41" `
	FCA1_42    string    `gorm:"column:FCA1_42;" json:"fca1_42"  form:"fca1_42" `
	FCA1_43    string    `gorm:"column:FCA1_43;" json:"fca1_43"  form:"fca1_43" `
	FCA1_44    string    `gorm:"column:FCA1_44;" json:"fca1_44"  form:"fca1_44" `
	FCA1_49    string    `gorm:"column:FCA1_49;" json:"fca1_49"  form:"fca1_49" `
	FCA1_50    string    `gorm:"column:FCA1_50;" json:"fca1_50"  form:"fca1_50" `
	FCA1_51    string    `gorm:"column:FCA1_51;" json:"fca1_51"  form:"fca1_51" `
	FCA1_52    string    `gorm:"column:FCA1_52;" json:"fca1_52"  form:"fca1_52" `
	FCA1_53    string    `gorm:"column:FCA1_53;" json:"fca1_53"  form:"fca1_53" `
	FCA1_54    string    `gorm:"column:FCA1_54;" json:"fca1_54"  form:"fca1_54" `
	FCA1_55    string    `gorm:"column:FCA1_55;" json:"fca1_55"  form:"fca1_55" `
	FCA2_14    string    `gorm:"column:FCA2_14;" json:"fca2_14"  form:"fca2_14" `
	FCA2_15    string    `gorm:"column:FCA2_15;" json:"fca2_15"  form:"fca2_15" `
	FCA2_16    string    `gorm:"column:FCA2_16;" json:"fca2_16"  form:"fca2_16" `
	FCA2_17    string    `gorm:"column:FCA2_17;" json:"fca2_17"  form:"fca2_17" `
	FCA2_18    string    `gorm:"column:FCA2_18;" json:"fca2_18"  form:"fca2_18" `
	FCA2_19    string    `gorm:"column:FCA2_19;" json:"fca2_19"  form:"fca2_19" `
	FCA2_20    string    `gorm:"column:FCA2_20;" json:"fca2_20"  form:"fca2_20" `
	FCA2_21    string    `gorm:"column:FCA2_21;" json:"fca2_21"  form:"fca2_21" `
	FCA2_25    string    `gorm:"column:FCA2_25;" json:"fca2_25"  form:"fca2_25" `
	FCA2_29    string    `gorm:"column:FCA2_29;" json:"fca2_29"  form:"fca2_29" `
	FCA2_30    string    `gorm:"column:FCA2_30;" json:"fca2_30"  form:"fca2_30" `
	FCA3_14    string    `gorm:"column:FCA3_14;" json:"fca3_14"  form:"fca3_14" `
	FCA3_15    string    `gorm:"column:FCA3_15;" json:"fca3_15"  form:"fca3_15" `
	FCA3_16    string    `gorm:"column:FCA3_16;" json:"fca3_16"  form:"fca3_16" `
	FCA3_17    string    `gorm:"column:FCA3_17;" json:"fca3_17"  form:"fca3_17" `
	FCA3_18    string    `gorm:"column:FCA3_18;" json:"fca3_18"  form:"fca3_18" `
	FCA4_14    string    `gorm:"column:FCA4_14;" json:"fca4_14"  form:"fca4_14" `
	FCA4_15    string    `gorm:"column:FCA4_15;" json:"fca4_15"  form:"fca4_15" `
	FCA4_16    string    `gorm:"column:FCA4_16;" json:"fca4_16"  form:"fca4_16" `
	FCA4_17    string    `gorm:"column:FCA4_17;" json:"fca4_17"  form:"fca4_17" `
	FCA4_18    string    `gorm:"column:FCA4_18;" json:"fca4_18"  form:"fca4_18" `
	FCA5_14    string    `gorm:"column:FCA5_14;" json:"fca5_14"  form:"fca5_14" `
	FCA5_15    string    `gorm:"column:FCA5_15;" json:"fca5_15"  form:"fca5_15" `
	FCA5_16    string    `gorm:"column:FCA5_16;" json:"fca5_16"  form:"fca5_16" `
	FCA5_17    string    `gorm:"column:FCA5_17;" json:"fca5_17"  form:"fca5_17" `
	FCA5_18    string    `gorm:"column:FCA5_18;" json:"fca5_18"  form:"fca5_18" `
	FCA6_14    string    `gorm:"column:FCA6_14;" json:"fca6_14"  form:"fca6_14" `
	FCA6_15    string    `gorm:"column:FCA6_15;" json:"fca6_15"  form:"fca6_15" `
	FCA6_16    string    `gorm:"column:FCA6_16;" json:"fca6_16"  form:"fca6_16" `
	FCA6_17    string    `gorm:"column:FCA6_17;" json:"fca6_17"  form:"fca6_17" `
	FCA6_18    string    `gorm:"column:FCA6_18;" json:"fca6_18"  form:"fca6_18" `
	FCA6_19    string    `gorm:"column:FCA6_19;" json:"fca6_19"  form:"fca6_19" `
	FCA6_20    string    `gorm:"column:FCA6_20;" json:"fca6_20"  form:"fca6_20" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCB1_14    string    `gorm:"column:FCB1_14;" json:"fcb1_14"  form:"fcb1_14" `
	FCB1_15    string    `gorm:"column:FCB1_15;" json:"fcb1_15"  form:"fcb1_15" `
	FCB1_16    string    `gorm:"column:FCB1_16;" json:"fcb1_16"  form:"fcb1_16" `
	FCB1_17    string    `gorm:"column:FCB1_17;" json:"fcb1_17"  form:"fcb1_17" `
	FCB1_18    string    `gorm:"column:FCB1_18;" json:"fcb1_18"  form:"fcb1_18" `
	FCB1_19    string    `gorm:"column:FCB1_19;" json:"fcb1_19"  form:"fcb1_19" `
	FCB1_20    string    `gorm:"column:FCB1_20;" json:"fcb1_20"  form:"fcb1_20" `
	FCB1_21    string    `gorm:"column:FCB1_21;" json:"fcb1_21"  form:"fcb1_21" `
	FCB1_22    string    `gorm:"column:FCB1_22;" json:"fcb1_22"  form:"fcb1_22" `
	FCB1_23    string    `gorm:"column:FCB1_23;" json:"fcb1_23"  form:"fcb1_23" `
	FCB2_14    string    `gorm:"column:FCB2_14;" json:"fcb2_14"  form:"fcb2_14" `
	FCB2_15    string    `gorm:"column:FCB2_15;" json:"fcb2_15"  form:"fcb2_15" `
	FCB2_16    string    `gorm:"column:FCB2_16;" json:"fcb2_16"  form:"fcb2_16" `
	FCB2_17    string    `gorm:"column:FCB2_17;" json:"fcb2_17"  form:"fcb2_17" `
	FCB2_18    string    `gorm:"column:FCB2_18;" json:"fcb2_18"  form:"fcb2_18" `
	FCB2_19    string    `gorm:"column:FCB2_19;" json:"fcb2_19"  form:"fcb2_19" `
	FCBOOK     string    `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCC1_14    string    `gorm:"column:FCC1_14;" json:"fcc1_14"  form:"fcc1_14" `
	FCC1_15    string    `gorm:"column:FCC1_15;" json:"fcc1_15"  form:"fcc1_15" `
	FCC1_16    string    `gorm:"column:FCC1_16;" json:"fcc1_16"  form:"fcc1_16" `
	FCC1_17    string    `gorm:"column:FCC1_17;" json:"fcc1_17"  form:"fcc1_17" `
	FCC1_18    string    `gorm:"column:FCC1_18;" json:"fcc1_18"  form:"fcc1_18" `
	FCC1_19    string    `gorm:"column:FCC1_19;" json:"fcc1_19"  form:"fcc1_19" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
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
	FCGLREF    string    `gorm:"column:FCGLREF;" json:"fcglref"  form:"fcglref" `
	FCI1_24    string    `gorm:"column:FCI1_24;" json:"fci1_24"  form:"fci1_24" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPACKH    string    `gorm:"column:FCPACKH;" json:"fcpackh"  form:"fcpackh" `
	FCRCVEDI   string    `gorm:"column:FCRCVEDI;" json:"fcrcvedi"  form:"fcrcvedi" `
	FCRECSTAT  string    `gorm:"column:FCRECSTAT;" json:"fcrecstat"  form:"fcrecstat" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRENEWAL  string    `gorm:"column:FCRENEWAL;" json:"fcrenewal"  form:"fcrenewal" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSENDSTAT string    `gorm:"column:FCSENDSTAT;" json:"fcsendstat"  form:"fcsendstat" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCTXID     string    `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
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
	FDA1_33    string    `gorm:"column:FDA1_33;" json:"fda1_33"  form:"fda1_33" `
	FDA1_35    string    `gorm:"column:FDA1_35;" json:"fda1_35"  form:"fda1_35" `
	FDA1_36    string    `gorm:"column:FDA1_36;" json:"fda1_36"  form:"fda1_36" `
	FDA1_45    string    `gorm:"column:FDA1_45;" json:"fda1_45"  form:"fda1_45" `
	FDA2_26    string    `gorm:"column:FDA2_26;" json:"fda2_26"  form:"fda2_26" `
	FDA2_27    string    `gorm:"column:FDA2_27;" json:"fda2_27"  form:"fda2_27" `
	FIMILISEC  string    `gorm:"column:FIMILISEC;" json:"fimilisec"  form:"fimilisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNA1_11    string    `gorm:"column:FNA1_11;" json:"fna1_11"  form:"fna1_11" `
	FNA1_12    string    `gorm:"column:FNA1_12;" json:"fna1_12"  form:"fna1_12" `
	FNA1_13    string    `gorm:"column:FNA1_13;" json:"fna1_13"  form:"fna1_13" `
	FNA1_46    string    `gorm:"column:FNA1_46;" json:"fna1_46"  form:"fna1_46" `
	FNA1_47    string    `gorm:"column:FNA1_47;" json:"fna1_47"  form:"fna1_47" `
	FNA1_48    string    `gorm:"column:FNA1_48;" json:"fna1_48"  form:"fna1_48" `
	FNA2_11    string    `gorm:"column:FNA2_11;" json:"fna2_11"  form:"fna2_11" `
	FNA2_12    string    `gorm:"column:FNA2_12;" json:"fna2_12"  form:"fna2_12" `
	FNA2_13    string    `gorm:"column:FNA2_13;" json:"fna2_13"  form:"fna2_13" `
	FNA2_22    string    `gorm:"column:FNA2_22;" json:"fna2_22"  form:"fna2_22" `
	FNA2_23    string    `gorm:"column:FNA2_23;" json:"fna2_23"  form:"fna2_23" `
	FNA2_24    string    `gorm:"column:FNA2_24;" json:"fna2_24"  form:"fna2_24" `
	FNA2_28    string    `gorm:"column:FNA2_28;" json:"fna2_28"  form:"fna2_28" `
	FNA3_11    string    `gorm:"column:FNA3_11;" json:"fna3_11"  form:"fna3_11" `
	FNA3_12    string    `gorm:"column:FNA3_12;" json:"fna3_12"  form:"fna3_12" `
	FNA3_13    string    `gorm:"column:FNA3_13;" json:"fna3_13"  form:"fna3_13" `
	FNA4_11    string    `gorm:"column:FNA4_11;" json:"fna4_11"  form:"fna4_11" `
	FNA4_12    string    `gorm:"column:FNA4_12;" json:"fna4_12"  form:"fna4_12" `
	FNA4_13    string    `gorm:"column:FNA4_13;" json:"fna4_13"  form:"fna4_13" `
	FNA5_11    string    `gorm:"column:FNA5_11;" json:"fna5_11"  form:"fna5_11" `
	FNA5_12    string    `gorm:"column:FNA5_12;" json:"fna5_12"  form:"fna5_12" `
	FNA5_13    string    `gorm:"column:FNA5_13;" json:"fna5_13"  form:"fna5_13" `
	FNA6_11    string    `gorm:"column:FNA6_11;" json:"fna6_11"  form:"fna6_11" `
	FNA6_12    string    `gorm:"column:FNA6_12;" json:"fna6_12"  form:"fna6_12" `
	FNA6_13    string    `gorm:"column:FNA6_13;" json:"fna6_13"  form:"fna6_13" `
	FNB1_11    string    `gorm:"column:FNB1_11;" json:"fnb1_11"  form:"fnb1_11" `
	FNB1_12    string    `gorm:"column:FNB1_12;" json:"fnb1_12"  form:"fnb1_12" `
	FNB1_13    string    `gorm:"column:FNB1_13;" json:"fnb1_13"  form:"fnb1_13" `
	FNB2_11    string    `gorm:"column:FNB2_11;" json:"fnb2_11"  form:"fnb2_11" `
	FNB2_12    string    `gorm:"column:FNB2_12;" json:"fnb2_12"  form:"fnb2_12" `
	FNB2_13    string    `gorm:"column:FNB2_13;" json:"fnb2_13"  form:"fnb2_13" `
	FNC1_11    string    `gorm:"column:FNC1_11;" json:"fnc1_11"  form:"fnc1_11" `
	FNC1_12    string    `gorm:"column:FNC1_12;" json:"fnc1_12"  form:"fnc1_12" `
	FNC1_13    string    `gorm:"column:FNC1_13;" json:"fnc1_13"  form:"fnc1_13" `
	FNI1_16    string    `gorm:"column:FNI1_16;" json:"fni1_16"  form:"fni1_16" `
	FNI1_17    string    `gorm:"column:FNI1_17;" json:"fni1_17"  form:"fni1_17" `
	FNI1_18    string    `gorm:"column:FNI1_18;" json:"fni1_18"  form:"fni1_18" `
	FNI1_22    string    `gorm:"column:FNI1_22;" json:"fni1_22"  form:"fni1_22" `
	FNI1_23    string    `gorm:"column:FNI1_23;" json:"fni1_23"  form:"fni1_23" `
	FNU1CNT    string    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTCREATED  string    `gorm:"column:FTCREATED;" json:"ftcreated"  form:"ftcreated" `
	FTDATETIME string    `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Vpackh01) TableName() string {
	return "VPACKH01"
}

func (obj *Vpackh01) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
