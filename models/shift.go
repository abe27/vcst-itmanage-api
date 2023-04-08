package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Shift struct {
	FCAFSALADJ string `gorm:"column:FCAFSALADJ;" json:"fcafsaladj"  form:"fcafsaladj" `
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCASALADJH string `gorm:"column:FCASALADJH;" json:"fcasaladjh"  form:"fcasaladjh" `
	FCASALADJW string `gorm:"column:FCASALADJW;" json:"fcasaladjw"  form:"fcasaladjw" `
	FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBFSALADJ string `gorm:"column:FCBFSALADJ;" json:"fcbfsaladj"  form:"fcbfsaladj" `
	FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCALOT    string `gorm:"column:FCCALOT;" json:"fccalot"  form:"fccalot" `
	FCCODE     string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATASER  string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDATEDF   string `gorm:"column:FCDATEDF;" json:"fcdatedf"  form:"fcdatedf" `
	FCDTYPE1   string `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR  string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEMFIXDED string `gorm:"column:FCEMFIXDED;" json:"fcemfixded"  form:"fcemfixded" `
	FCEMFIXINC string `gorm:"column:FCEMFIXINC;" json:"fcemfixinc"  form:"fcemfixinc" `
	FCFCHR     string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCFRIBEG1  string `gorm:"column:FCFRIBEG1;" json:"fcfribeg1"  form:"fcfribeg1" `
	FCFRIBEG2  string `gorm:"column:FCFRIBEG2;" json:"fcfribeg2"  form:"fcfribeg2" `
	FCFRIEND1  string `gorm:"column:FCFRIEND1;" json:"fcfriend1"  form:"fcfriend1" `
	FCFRIEND2  string `gorm:"column:FCFRIEND2;" json:"fcfriend2"  form:"fcfriend2" `
	FCGID      string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCISAPPROV string `gorm:"column:FCISAPPROV;" json:"fcisapprov"  form:"fcisapprov" `
	FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMAXREG   string `gorm:"column:FCMAXREG;" json:"fcmaxreg"  form:"fcmaxreg" `
	FCMINREG   string `gorm:"column:FCMINREG;" json:"fcminreg"  form:"fcminreg" `
	FCMONBEG1  string `gorm:"column:FCMONBEG1;" json:"fcmonbeg1"  form:"fcmonbeg1" `
	FCMONBEG2  string `gorm:"column:FCMONBEG2;" json:"fcmonbeg2"  form:"fcmonbeg2" `
	FCMONEND1  string `gorm:"column:FCMONEND1;" json:"fcmonend1"  form:"fcmonend1" `
	FCMONEND2  string `gorm:"column:FCMONEND2;" json:"fcmonend2"  form:"fcmonend2" `
	FCNAME     string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCREGBEG1  string `gorm:"column:FCREGBEG1;" json:"fcregbeg1"  form:"fcregbeg1" `
	FCREGBEG2  string `gorm:"column:FCREGBEG2;" json:"fcregbeg2"  form:"fcregbeg2" `
	FCREGBEGOT string `gorm:"column:FCREGBEGOT;" json:"fcregbegot"  form:"fcregbegot" `
	FCREGEND1  string `gorm:"column:FCREGEND1;" json:"fcregend1"  form:"fcregend1" `
	FCREGEND2  string `gorm:"column:FCREGEND2;" json:"fcregend2"  form:"fcregend2" `
	FCSATBEG1  string `gorm:"column:FCSATBEG1;" json:"fcsatbeg1"  form:"fcsatbeg1" `
	FCSATBEG2  string `gorm:"column:FCSATBEG2;" json:"fcsatbeg2"  form:"fcsatbeg2" `
	FCSATEND1  string `gorm:"column:FCSATEND1;" json:"fcsatend1"  form:"fcsatend1" `
	FCSATEND2  string `gorm:"column:FCSATEND2;" json:"fcsatend2"  form:"fcsatend2" `
	FCSECT     string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSSALADJH string `gorm:"column:FCSSALADJH;" json:"fcssaladjh"  form:"fcssaladjh" `
	FCSSALADJW string `gorm:"column:FCSSALADJW;" json:"fcssaladjw"  form:"fcssaladjw" `
	FCSTATUS   string `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCSTDBEGH1 string `gorm:"column:FCSTDBEGH1;" json:"fcstdbegh1"  form:"fcstdbegh1" `
	FCSTDBEGH2 string `gorm:"column:FCSTDBEGH2;" json:"fcstdbegh2"  form:"fcstdbegh2" `
	FCSTDBEGW1 string `gorm:"column:FCSTDBEGW1;" json:"fcstdbegw1"  form:"fcstdbegw1" `
	FCSTDBEGW2 string `gorm:"column:FCSTDBEGW2;" json:"fcstdbegw2"  form:"fcstdbegw2" `
	FCSTDENDH1 string `gorm:"column:FCSTDENDH1;" json:"fcstdendh1"  form:"fcstdendh1" `
	FCSTDENDH2 string `gorm:"column:FCSTDENDH2;" json:"fcstdendh2"  form:"fcstdendh2" `
	FCSTDENDW1 string `gorm:"column:FCSTDENDW1;" json:"fcstdendw1"  form:"fcstdendw1" `
	FCSTDENDW2 string `gorm:"column:FCSTDENDW2;" json:"fcstdendw2"  form:"fcstdendw2" `
	FCSUNBEG1  string `gorm:"column:FCSUNBEG1;" json:"fcsunbeg1"  form:"fcsunbeg1" `
	FCSUNBEG2  string `gorm:"column:FCSUNBEG2;" json:"fcsunbeg2"  form:"fcsunbeg2" `
	FCSUNEND1  string `gorm:"column:FCSUNEND1;" json:"fcsunend1"  form:"fcsunend1" `
	FCSUNEND2  string `gorm:"column:FCSUNEND2;" json:"fcsunend2"  form:"fcsunend2" `
	FCTHUBEG1  string `gorm:"column:FCTHUBEG1;" json:"fcthubeg1"  form:"fcthubeg1" `
	FCTHUBEG2  string `gorm:"column:FCTHUBEG2;" json:"fcthubeg2"  form:"fcthubeg2" `
	FCTHUEND1  string `gorm:"column:FCTHUEND1;" json:"fcthuend1"  form:"fcthuend1" `
	FCTHUEND2  string `gorm:"column:FCTHUEND2;" json:"fcthuend2"  form:"fcthuend2" `
	FCTUEBEG1  string `gorm:"column:FCTUEBEG1;" json:"fctuebeg1"  form:"fctuebeg1" `
	FCTUEBEG2  string `gorm:"column:FCTUEBEG2;" json:"fctuebeg2"  form:"fctuebeg2" `
	FCTUEEND1  string `gorm:"column:FCTUEEND1;" json:"fctueend1"  form:"fctueend1" `
	FCTUEEND2  string `gorm:"column:FCTUEEND2;" json:"fctueend2"  form:"fctueend2" `
	FCTYEMDED  string `gorm:"column:FCTYEMDED;" json:"fctyemded"  form:"fctyemded" `
	FCTYEMINC  string `gorm:"column:FCTYEMINC;" json:"fctyeminc"  form:"fctyeminc" `
	FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCWEDBEG1  string `gorm:"column:FCWEDBEG1;" json:"fcwedbeg1"  form:"fcwedbeg1" `
	FCWEDBEG2  string `gorm:"column:FCWEDBEG2;" json:"fcwedbeg2"  form:"fcwedbeg2" `
	FCWEDEND1  string `gorm:"column:FCWEDEND1;" json:"fcwedend1"  form:"fcwedend1" `
	FCWEDEND2  string `gorm:"column:FCWEDEND2;" json:"fcwedend2"  form:"fcwedend2" `
	FDINACTIVE string `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNAFOTMAX  string `gorm:"column:FNAFOTMAX;" json:"fnafotmax"  form:"fnafotmax" `
	FNAFOTMAXH string `gorm:"column:FNAFOTMAXH;" json:"fnafotmaxh"  form:"fnafotmaxh" `
	FNAFOTMAXW string `gorm:"column:FNAFOTMAXW;" json:"fnafotmaxw"  form:"fnafotmaxw" `
	FNAFOTMIN  string `gorm:"column:FNAFOTMIN;" json:"fnafotmin"  form:"fnafotmin" `
	FNAFOTMINH string `gorm:"column:FNAFOTMINH;" json:"fnafotminh"  form:"fnafotminh" `
	FNAFOTMINW string `gorm:"column:FNAFOTMINW;" json:"fnafotminw"  form:"fnafotminw" `
	FNBFOTMAX  string `gorm:"column:FNBFOTMAX;" json:"fnbfotmax"  form:"fnbfotmax" `
	FNBFOTMAXH string `gorm:"column:FNBFOTMAXH;" json:"fnbfotmaxh"  form:"fnbfotmaxh" `
	FNBFOTMAXW string `gorm:"column:FNBFOTMAXW;" json:"fnbfotmaxw"  form:"fnbfotmaxw" `
	FNBFOTMIN  string `gorm:"column:FNBFOTMIN;" json:"fnbfotmin"  form:"fnbfotmin" `
	FNBFOTMINH string `gorm:"column:FNBFOTMINH;" json:"fnbfotminh"  form:"fnbfotminh" `
	FNBFOTMINW string `gorm:"column:FNBFOTMINW;" json:"fnbfotminw"  form:"fnbfotminw" `
	FNBWBREAKH string `gorm:"column:FNBWBREAKH;" json:"fnbwbreakh"  form:"fnbwbreakh" `
	FNBWBREAKW string `gorm:"column:FNBWBREAKW;" json:"fnbwbreakw"  form:"fnbwbreakw" `
	FNOTBREAK  string `gorm:"column:FNOTBREAK;" json:"fnotbreak"  form:"fnotbreak" `
	FNOTBREAKH string `gorm:"column:FNOTBREAKH;" json:"fnotbreakh"  form:"fnotbreakh" `
	FNOTBREAKW string `gorm:"column:FNOTBREAKW;" json:"fnotbreakw"  form:"fnotbreakw" `
	FNSTOTMAXH string `gorm:"column:FNSTOTMAXH;" json:"fnstotmaxh"  form:"fnstotmaxh" `
	FNSTOTMAXW string `gorm:"column:FNSTOTMAXW;" json:"fnstotmaxw"  form:"fnstotmaxw" `
	FNSTOTMINH string `gorm:"column:FNSTOTMINH;" json:"fnstotminh"  form:"fnstotminh" `
	FNSTOTMINW string `gorm:"column:FNSTOTMINW;" json:"fnstotminw"  form:"fnstotminw" `
	FNU1CNT    string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    string `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    string `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    string `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    string `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    string `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    string `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    string `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
	FTSRCUPD   string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Shift) TableName() string {
	return "SHIFT"
}

func (obj *Shift) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
