package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Coor struct {
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCISCUST   string    `gorm:"column:FCISCUST;" json:"fciscust"  form:"fciscust" `
	FCISSUPP   string    `gorm:"column:FCISSUPP;" json:"fcissupp"  form:"fcissupp" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCRGRP    string    `gorm:"column:FCCRGRP;" json:"fccrgrp"  form:"fccrgrp" `
	FCCRZONE   string    `gorm:"column:FCCRZONE;" json:"fccrzone"  form:"fccrzone" `
	FCEMPL     string    `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	FCSNAME    string    `gorm:"column:FCSNAME;" json:"fcsname"  form:"fcsname" `
	FCSNAME2   string    `gorm:"column:FCSNAME2;" json:"fcsname2"  form:"fcsname2" `
	FCNAME2    string    `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FNCREDTERM float64   `gorm:"column:FNCREDTERM;" json:"fncredterm"  form:"fncredterm" `
	FNCREDLIM  float64   `gorm:"column:FNCREDLIM;" json:"fncredlim"  form:"fncredlim" `
	FCZIP      string    `gorm:"column:FCZIP;" json:"fczip"  form:"fczip" `
	FNPAYFIFO  float64   `gorm:"column:FNPAYFIFO;" json:"fnpayfifo"  form:"fnpayfifo" `
	FCGRADE    string    `gorm:"column:FCGRADE;" json:"fcgrade"  form:"fcgrade" `
	FCPRICENO  string    `gorm:"column:FCPRICENO;" json:"fcpriceno"  form:"fcpriceno" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FNDISCPCN  float64   `gorm:"column:FNDISCPCN;" json:"fndiscpcn"  form:"fndiscpcn" `
	FCACCHART  string    `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FCPERSONTY string    `gorm:"column:FCPERSONTY;" json:"fcpersonty"  form:"fcpersonty" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FCVATCOOR  string    `gorm:"column:FCVATCOOR;" json:"fcvatcoor"  form:"fcvatcoor" `
	FCCOLCOOR  string    `gorm:"column:FCCOLCOOR;" json:"fccolcoor"  form:"fccolcoor" `
	FCPOLICYPR string    `gorm:"column:FCPOLICYPR;" json:"fcpolicypr"  form:"fcpolicypr" `
	FCPOLICYDI string    `gorm:"column:FCPOLICYDI;" json:"fcpolicydi"  form:"fcpolicydi" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCCONTACTN string    `gorm:"column:FCCONTACTN;" json:"fccontactn"  form:"fccontactn" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCFCHRS    string    `gorm:"column:FCFCHRS;" json:"fcfchrs"  form:"fcfchrs" `
	FNBLACK    float64   `gorm:"column:FNBLACK;" json:"fnblack"  form:"fnblack" `
	FCCONTPOS  string    `gorm:"column:FCCONTPOS;" json:"fccontpos"  form:"fccontpos" `
	FCDELICOOR string    `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	FMMEMDATA2 string    `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
	FMMEMDATA3 string    `gorm:"column:FMMEMDATA3;" json:"fmmemdata3"  form:"fmmemdata3" `
	FMMEMDATA4 string    `gorm:"column:FMMEMDATA4;" json:"fmmemdata4"  form:"fmmemdata4" `
	FCBACCMTH  string    `gorm:"column:FCBACCMTH;" json:"fcbaccmth"  form:"fcbaccmth" `
	FCCURRENCY string    `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	FMMAPPICT  string    `gorm:"column:FMMAPPICT;" json:"fmmappict"  form:"fmmappict" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FDFSTCONT  string    `gorm:"column:FDFSTCONT;" json:"fdfstcont"  form:"fdfstcont" `
	FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FDINACTIVE string    `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	FCPAYTERM  string    `gorm:"column:FCPAYTERM;" json:"fcpayterm"  form:"fcpayterm" `
	FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCDEBTAAC  string    `gorm:"column:FCDEBTAAC;" json:"fcdebtaac"  form:"fcdebtaac" `
	FCBDEBTAAC string    `gorm:"column:FCBDEBTAAC;" json:"fcbdebtaac"  form:"fcbdebtaac" `
	FCPAYGRP   string    `gorm:"column:FCPAYGRP;" json:"fcpaygrp"  form:"fcpaygrp" `
	FCBANK     string    `gorm:"column:FCBANK;" json:"fcbank"  form:"fcbank" `
	FCBBRANCH  string    `gorm:"column:FCBBRANCH;" json:"fcbbranch"  form:"fcbbranch" `
	FCBANKNO   string    `gorm:"column:FCBANKNO;" json:"fcbankno"  form:"fcbankno" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FCEMAIL    string    `gorm:"column:FCEMAIL;" json:"fcemail"  form:"fcemail" `
	FCLAYMETHD string    `gorm:"column:FCLAYMETHD;" json:"fclaymethd"  form:"fclaymethd" `
	FCCHQMETHD string    `gorm:"column:FCCHQMETHD;" json:"fcchqmethd"  form:"fcchqmethd" `
	FMMAPNAME  string    `gorm:"column:FMMAPNAME;" json:"fmmapname"  form:"fmmapname" `
	FCTRADETRM string    `gorm:"column:FCTRADETRM;" json:"fctradetrm"  form:"fctradetrm" `
	FCDELIEMPL string    `gorm:"column:FCDELIEMPL;" json:"fcdeliempl"  form:"fcdeliempl" `
	FCLAYEMPL  string    `gorm:"column:FCLAYEMPL;" json:"fclayempl"  form:"fclayempl" `
	FMMEMDATA5 string    `gorm:"column:FMMEMDATA5;" json:"fmmemdata5"  form:"fmmemdata5" `
	FCTEL      string    `gorm:"column:FCTEL;" json:"fctel"  form:"fctel" `
	FCDCHANNEL string    `gorm:"column:FCDCHANNEL;" json:"fcdchannel"  form:"fcdchannel" `
	FMPICSHMK  string    `gorm:"column:FMPICSHMK;" json:"fmpicshmk"  form:"fmpicshmk" `
	FNLAT      float64   `gorm:"column:FNLAT;" json:"fnlat"  form:"fnlat" `
	FNLONG     float64   `gorm:"column:FNLONG;" json:"fnlong"  form:"fnlong" `
	FNLTDPO    float64   `gorm:"column:FNLTDPO;" json:"fnltdpo"  form:"fnltdpo" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FCCOORPGRP string    `gorm:"column:FCCOORPGRP;" json:"fccoorpgrp"  form:"fccoorpgrp" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCIS1SHIP  string    `gorm:"column:FCIS1SHIP;" json:"fcis1ship"  form:"fcis1ship" `
	FCROUTEDEL string    `gorm:"column:FCROUTEDEL;" json:"fcroutedel"  form:"fcroutedel" `
	FCLINKAPP1 string    `gorm:"column:FCLINKAPP1;" json:"fclinkapp1"  form:"fclinkapp1" `
	FCLINKSTP1 string    `gorm:"column:FCLINKSTP1;" json:"fclinkstp1"  form:"fclinkstp1" `
	FCLINKAPP2 string    `gorm:"column:FCLINKAPP2;" json:"fclinkapp2"  form:"fclinkapp2" `
	FCLINKSTP2 string    `gorm:"column:FCLINKSTP2;" json:"fclinkstp2"  form:"fclinkstp2" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCCOLSEND  string    `gorm:"column:FCCOLSEND;" json:"fccolsend"  form:"fccolsend" `
	FCBKBRCODE string    `gorm:"column:FCBKBRCODE;" json:"fcbkbrcode"  form:"fcbkbrcode" `
	FCPYMETHOD string    `gorm:"column:FCPYMETHOD;" json:"fcpymethod"  form:"fcpymethod" `
	FCCHGBEAR  string    `gorm:"column:FCCHGBEAR;" json:"fcchgbear"  form:"fcchgbear" `
	FCADVICEBY string    `gorm:"column:FCADVICEBY;" json:"fcadviceby"  form:"fcadviceby" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCBANKACCT string    `gorm:"column:FCBANKACCT;" json:"fcbankacct"  form:"fcbankacct" `
	FCNPYMETHD string    `gorm:"column:FCNPYMETHD;" json:"fcnpymethd"  form:"fcnpymethd" `
	FCSVLEVEL  string    `gorm:"column:FCSVLEVEL;" json:"fcsvlevel"  form:"fcsvlevel" `
	FCLOCALINS string    `gorm:"column:FCLOCALINS;" json:"fclocalins"  form:"fclocalins" `
	FCNCHRGBR  string    `gorm:"column:FCNCHRGBR;" json:"fcnchrgbr"  form:"fcnchrgbr" `
	FCBANKBR   string    `gorm:"column:FCBANKBR;" json:"fcbankbr"  form:"fcbankbr" `
	FUNAME     string    `gorm:"column:FUNAME;" json:"funame"  form:"funame" `
	FUADDR11   string    `gorm:"column:FUADDR11;" json:"fuaddr11"  form:"fuaddr11" `
	FUADDR21   string    `gorm:"column:FUADDR21;" json:"fuaddr21"  form:"fuaddr21" `
	FUADDR31   string    `gorm:"column:FUADDR31;" json:"fuaddr31"  form:"fuaddr31" `
	FCCHEQTYP  string    `gorm:"column:FCCHEQTYP;" json:"fccheqtyp"  form:"fccheqtyp" `
	FCDELIBY   string    `gorm:"column:FCDELIBY;" json:"fcdeliby"  form:"fcdeliby" `
	FCPICKLOCA string    `gorm:"column:FCPICKLOCA;" json:"fcpickloca"  form:"fcpickloca" `
	FCBICCODE  string    `gorm:"column:FCBICCODE;" json:"fcbiccode"  form:"fcbiccode" `
	FCPRENAME  string    `gorm:"column:FCPRENAME;" json:"fcprename"  form:"fcprename" `
	FCPRENAME2 string    `gorm:"column:FCPRENAME2;" json:"fcprename2"  form:"fcprename2" `
	FCNO       string    `gorm:"column:FCNO;" json:"fcno"  form:"fcno" `
	FCFLOOR    string    `gorm:"column:FCFLOOR;" json:"fcfloor"  form:"fcfloor" `
	FCROOMNO   string    `gorm:"column:FCROOMNO;" json:"fcroomno"  form:"fcroomno" `
	FCBUILDING string    `gorm:"column:FCBUILDING;" json:"fcbuilding"  form:"fcbuilding" `
	FCMOO      string    `gorm:"column:FCMOO;" json:"fcmoo"  form:"fcmoo" `
	FCSOI      string    `gorm:"column:FCSOI;" json:"fcsoi"  form:"fcsoi" `
	FCROAD     string    `gorm:"column:FCROAD;" json:"fcroad"  form:"fcroad" `
	FCTAMBON   string    `gorm:"column:FCTAMBON;" json:"fctambon"  form:"fctambon" `
	FCAMPHUR   string    `gorm:"column:FCAMPHUR;" json:"fcamphur"  form:"fcamphur" `
	FCPROVINCE string    `gorm:"column:FCPROVINCE;" json:"fcprovince"  form:"fcprovince" `
	FCISHEADBR string    `gorm:"column:FCISHEADBR;" json:"fcisheadbr"  form:"fcisheadbr" `
	FCBRANCHNO string    `gorm:"column:FCBRANCHNO;" json:"fcbranchno"  form:"fcbranchno" `
	FCBRANCHNM string    `gorm:"column:FCBRANCHNM;" json:"fcbranchnm"  form:"fcbranchnm" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FCACCDEP   string    `gorm:"column:FCACCDEP;" json:"fcaccdep"  form:"fcaccdep" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCBUILDIN2 string    `gorm:"column:FCBUILDIN2;" json:"fcbuildin2"  form:"fcbuildin2" `
	FCSOI2     string    `gorm:"column:FCSOI2;" json:"fcsoi2"  form:"fcsoi2" `
	FCROAD2    string    `gorm:"column:FCROAD2;" json:"fcroad2"  form:"fcroad2" `
	FCQCJANGWA string    `gorm:"column:FCQCJANGWA;" json:"fcqcjangwa"  form:"fcqcjangwa" `
	FCQNJANGW2 string    `gorm:"column:FCQNJANGW2;" json:"fcqnjangw2"  form:"fcqnjangw2" `
	FCQCAMPHUR string    `gorm:"column:FCQCAMPHUR;" json:"fcqcamphur"  form:"fcqcamphur" `
	FCQNAMPHU2 string    `gorm:"column:FCQNAMPHU2;" json:"fcqnamphu2"  form:"fcqnamphu2" `
	FCQCTAMBON string    `gorm:"column:FCQCTAMBON;" json:"fcqctambon"  form:"fcqctambon" `
	FCQNTAMBO2 string    `gorm:"column:FCQNTAMBO2;" json:"fcqntambo2"  form:"fcqntambo2" `
	FCQCCOUNTY string    `gorm:"column:FCQCCOUNTY;" json:"fcqccounty"  form:"fcqccounty" `
	FCCARRIER  string    `gorm:"column:FCCARRIER;" json:"fccarrier"  form:"fccarrier" `
	FCCARRYTYP string    `gorm:"column:FCCARRYTYP;" json:"fccarrytyp"  form:"fccarrytyp" `
	FMMEMDATA6 string    `gorm:"column:FMMEMDATA6;" json:"fmmemdata6"  form:"fmmemdata6" `
	FMMEMDATA7 string    `gorm:"column:FMMEMDATA7;" json:"fmmemdata7"  form:"fmmemdata7" `
	FMMEMDATA8 string    `gorm:"column:FMMEMDATA8;" json:"fmmemdata8"  form:"fmmemdata8" `
	FMMEMDATA9 string    `gorm:"column:FMMEMDATA9;" json:"fmmemdata9"  form:"fmmemdata9" `
	FMMEMDATAA string    `gorm:"column:FMMEMDATAA;" json:"fmmemdataa"  form:"fmmemdataa" `
	FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FCSENDETAX string    `gorm:"column:FCSENDETAX;" json:"fcsendetax"  form:"fcsendetax" `
	FDSTETAX   string    `gorm:"column:FDSTETAX;" json:"fdstetax"  form:"fdstetax" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Coor) TableName() string {
	return "COOR"
}

func (obj *Coor) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
