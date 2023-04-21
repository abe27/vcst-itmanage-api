package models

import (
    "fmt"
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Product struct {
	FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID    string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	// FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	// FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	// FCLUPDAPP  string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	// FCBAKYRHIS string `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCTYPE string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	// FCPDGRP    string `gorm:"column:FCPDGRP;" json:"fcpdgrp"  form:"fcpdgrp" `
	// FCBRANCH   string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE   string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCSNAME  string `gorm:"column:FCSNAME;" json:"fcsname"  form:"fcsname" `
	FCSNAME2 string `gorm:"column:FCSNAME2;" json:"fcsname2"  form:"fcsname2" `
	FCNAME   string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2  string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	// FCSUBSTITU  string      `gorm:"column:FCSUBSTITU;" json:"fcsubstitu"  form:"fcsubstitu" `
	// FNAVGCOST   float64     `gorm:"column:FNAVGCOST;" json:"fnavgcost"  form:"fnavgcost" `
	// FNSTDCOST   float64     `gorm:"column:FNSTDCOST;" json:"fnstdcost"  form:"fnstdcost" `
	// FNISCONSUM  float64     `gorm:"column:FNISCONSUM;" json:"fnisconsum"  form:"fnisconsum" `
	// FCACCBCASH  string      `gorm:"column:FCACCBCASH;" json:"fcaccbcash"  form:"fcaccbcash" `
	// FCACCBCRED  string      `gorm:"column:FCACCBCRED;" json:"fcaccbcred"  form:"fcaccbcred" `
	// FCACCSCASH  string      `gorm:"column:FCACCSCASH;" json:"fcaccscash"  form:"fcaccscash" `
	// FCACCSCRED  string      `gorm:"column:FCACCSCRED;" json:"fcaccscred"  form:"fcaccscred" `
	// FCLASTUM    string      `gorm:"column:FCLASTUM;" json:"fclastum"  form:"fclastum" `
	FCUM string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	// FCSTUM      string      `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	// FCUM1       string      `gorm:"column:FCUM1;" json:"fcum1"  form:"fcum1" `
	// FNUMQTY1    float64     `gorm:"column:FNUMQTY1;" json:"fnumqty1"  form:"fnumqty1" `
	// FCSTUM1     string      `gorm:"column:FCSTUM1;" json:"fcstum1"  form:"fcstum1" `
	// FNSTUMQTY1  float64     `gorm:"column:FNSTUMQTY1;" json:"fnstumqty1"  form:"fnstumqty1" `
	// FCUM2       string      `gorm:"column:FCUM2;" json:"fcum2"  form:"fcum2" `
	// FNUMQTY2    float64     `gorm:"column:FNUMQTY2;" json:"fnumqty2"  form:"fnumqty2" `
	// FCSTUM2     string      `gorm:"column:FCSTUM2;" json:"fcstum2"  form:"fcstum2" `
	// FNSTUMQTY2  float64     `gorm:"column:FNSTUMQTY2;" json:"fnstumqty2"  form:"fnstumqty2" `
	// FCUM3       string      `gorm:"column:FCUM3;" json:"fcum3"  form:"fcum3" `
	// FNUMQTY3    float64     `gorm:"column:FNUMQTY3;" json:"fnumqty3"  form:"fnumqty3" `
	// FCSTUM3     string      `gorm:"column:FCSTUM3;" json:"fcstum3"  form:"fcstum3" `
	// FNSTUMQTY3  float64     `gorm:"column:FNSTUMQTY3;" json:"fnstumqty3"  form:"fnstumqty3" `
	// FCUM4       string      `gorm:"column:FCUM4;" json:"fcum4"  form:"fcum4" `
	// FNUMQTY4    float64     `gorm:"column:FNUMQTY4;" json:"fnumqty4"  form:"fnumqty4" `
	// FCSTUM4     string      `gorm:"column:FCSTUM4;" json:"fcstum4"  form:"fcstum4" `
	// FNSTUMQTY4  float64     `gorm:"column:FNSTUMQTY4;" json:"fnstumqty4"  form:"fnstumqty4" `
	// FCCOPYREMA  string      `gorm:"column:FCCOPYREMA;" json:"fccopyrema"  form:"fccopyrema" `
	// FNPRICE     float64     `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	// FCEDPRICE   string      `gorm:"column:FCEDPRICE;" json:"fcedprice"  form:"fcedprice" `
	// FNCOMMISSI  float64     `gorm:"column:FNCOMMISSI;" json:"fncommissi"  form:"fncommissi" `
	// FCCTRLSTOC  string      `gorm:"column:FCCTRLSTOC;" json:"fcctrlstoc"  form:"fcctrlstoc" `
	// FNSAFETYQT  float64     `gorm:"column:FNSAFETYQT;" json:"fnsafetyqt"  form:"fnsafetyqt" `
	// FNREORDERP  float64     `gorm:"column:FNREORDERP;" json:"fnreorderp"  form:"fnreorderp" `
	// FNEOQTY     float64     `gorm:"column:FNEOQTY;" json:"fneoqty"  form:"fneoqty" `
	// FCVATISOUT  string      `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	// FCVATTYPE   string      `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	// FNVATRATE   float64     `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	// FCCREATETY  string      `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCEAFTERR   string      `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	// FCSELTAG    string      `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	// FTDATETIME  time.Time   `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	// FIMILLISEC  int64       `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	// FCLOCATION  string      `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
	// FDLASTBUYD  string      `gorm:"column:FDLASTBUYD;" json:"fdlastbuyd"  form:"fdlastbuyd" `
	// FDLASTSELL  string      `gorm:"column:FDLASTSELL;" json:"fdlastsell"  form:"fdlastsell" `
	// FNLASTBUYP  float64     `gorm:"column:FNLASTBUYP;" json:"fnlastbuyp"  form:"fnlastbuyp" `
	// FCFCHR      string      `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	// FCFCHRS     string      `gorm:"column:FCFCHRS;" json:"fcfchrs"  form:"fcfchrs" `
	// FCSTEP      string      `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	// FCAPPROVEB  string      `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	// FCHASSALE   string      `gorm:"column:FCHASSALE;" json:"fchassale"  form:"fchassale" `
	// FNSALEPRIC  float64     `gorm:"column:FNSALEPRIC;" json:"fnsalepric"  form:"fnsalepric" `
	// FCFORMULAS  string      `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
	// FCACCWR1    string      `gorm:"column:FCACCWR1;" json:"fcaccwr1"  form:"fcaccwr1" `
	// FCACCWR2    string      `gorm:"column:FCACCWR2;" json:"fcaccwr2"  form:"fcaccwr2" `
	// FCACCEXP1   string      `gorm:"column:FCACCEXP1;" json:"fcaccexp1"  form:"fcaccexp1" `
	// FCACCVER    string      `gorm:"column:FCACCVER;" json:"fcaccver"  form:"fcaccver" `
	// FCSUPP      string      `gorm:"column:FCSUPP;" json:"fcsupp"  form:"fcsupp" `
	// FCMEASURE   string      `gorm:"column:FCMEASURE;" json:"fcmeasure"  form:"fcmeasure" `
	// FCWHATPRIC  string      `gorm:"column:FCWHATPRIC;" json:"fcwhatpric"  form:"fcwhatpric" `
	// FCPLANT     string      `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	// FCSOURCE    string      `gorm:"column:FCSOURCE;" json:"fcsource"  form:"fcsource" `
	// FCPROCURE   string      `gorm:"column:FCPROCURE;" json:"fcprocure"  form:"fcprocure" `
	// FNLTDPO     float64     `gorm:"column:FNLTDPO;" json:"fnltdpo"  form:"fnltdpo" `
	// FNLTDMFG    float64     `gorm:"column:FNLTDMFG;" json:"fnltdmfg"  form:"fnltdmfg" `
	// FNLTDSUBC   float64     `gorm:"column:FNLTDSUBC;" json:"fnltdsubc"  form:"fnltdsubc" `
	// FCSUPPBOI   string      `gorm:"column:FCSUPPBOI;" json:"fcsuppboi"  form:"fcsuppboi" `
	// FNSCOSTKE1  float64     `gorm:"column:FNSCOSTKE1;" json:"fnscostke1"  form:"fnscostke1" `
	// FNSCOSTKE2  float64     `gorm:"column:FNSCOSTKE2;" json:"fnscostke2"  form:"fnscostke2" `
	// FCCURR1     string      `gorm:"column:FCCURR1;" json:"fccurr1"  form:"fccurr1" `
	// FCCURR2     string      `gorm:"column:FCCURR2;" json:"fccurr2"  form:"fccurr2" `
	// FNMINPRICE  float64     `gorm:"column:FNMINPRICE;" json:"fnminprice"  form:"fnminprice" `
	// FNMAXPRICE  float64     `gorm:"column:FNMAXPRICE;" json:"fnmaxprice"  form:"fnmaxprice" `
	// FMPICNAME   string      `gorm:"column:FMPICNAME;" json:"fmpicname"  form:"fmpicname" `
	// FTLASTUPD   time.Time   `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	// FCCURRENCY  string      `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	// FNUMRATIO   float64     `gorm:"column:FNUMRATIO;" json:"fnumratio"  form:"fnumratio" `
	// FNEXPFAC    float64     `gorm:"column:FNEXPFAC;" json:"fnexpfac"  form:"fnexpfac" `
	// FCEXPCUR1   string      `gorm:"column:FCEXPCUR1;" json:"fcexpcur1"  form:"fcexpcur1" `
	// FCEXPCUR2   string      `gorm:"column:FCEXPCUR2;" json:"fcexpcur2"  form:"fcexpcur2" `
	// FNLTDQC     float64     `gorm:"column:FNLTDQC;" json:"fnltdqc"  form:"fnltdqc" `
	// FNLTDPPARE  float64     `gorm:"column:FNLTDPPARE;" json:"fnltdppare"  form:"fnltdppare" `
	// FCPRCGRP    string      `gorm:"column:FCPRCGRP;" json:"fcprcgrp"  form:"fcprcgrp" `
	// FCDISCSTR   string      `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	// FNAGELONG   float64     `gorm:"column:FNAGELONG;" json:"fnagelong"  form:"fnagelong" `
	// FCSTATUS    string      `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	// FDINACTIVE  string      `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	// FCTAXTYPE   string      `gorm:"column:FCTAXTYPE;" json:"fctaxtype"  form:"fctaxtype" `
	// FNGROSSWG   float64     `gorm:"column:FNGROSSWG;" json:"fngrosswg"  form:"fngrosswg" `
	// FCACCSCOST  string      `gorm:"column:FCACCSCOST;" json:"fcaccscost"  form:"fcaccscost" `
	// FCACCITEM   string      `gorm:"column:FCACCITEM;" json:"fcaccitem"  form:"fcaccitem" `
	// FCMFUM      string      `gorm:"column:FCMFUM;" json:"fcmfum"  form:"fcmfum" `
	// FNMFUMQTY   float64     `gorm:"column:FNMFUMQTY;" json:"fnmfumqty"  form:"fnmfumqty" `
	// FCMANUFAC   string      `gorm:"column:FCMANUFAC;" json:"fcmanufac"  form:"fcmanufac" `
	// FNUMRATIOA  float64     `gorm:"column:FNUMRATIOA;" json:"fnumratioa"  form:"fnumratioa" `
	// FNUMRATIOB  float64     `gorm:"column:FNUMRATIOB;" json:"fnumratiob"  form:"fnumratiob" `
	// FCDFAUMRAT  string      `gorm:"column:FCDFAUMRAT;" json:"fcdfaumrat"  form:"fcdfaumrat" `
	// FNDUTYPCN   float64     `gorm:"column:FNDUTYPCN;" json:"fndutypcn"  form:"fndutypcn" `
	// FCBVATISOU  string      `gorm:"column:FCBVATISOU;" json:"fcbvatisou"  form:"fcbvatisou" `
	// FNBATCHSIZ  float64     `gorm:"column:FNBATCHSIZ;" json:"fnbatchsiz"  form:"fnbatchsiz" `
	// FCAUTOGW1   string      `gorm:"column:FCAUTOGW1;" json:"fcautogw1"  form:"fcautogw1" `
	// FCMCTRLSTO  string      `gorm:"column:FCMCTRLSTO;" json:"fcmctrlsto"  form:"fcmctrlsto" `
	// FNWTAXRATE  float64     `gorm:"column:FNWTAXRATE;" json:"fnwtaxrate"  form:"fnwtaxrate" `
	// FDADD       string      `gorm:"column:FDADD;" json:"fdadd"  form:"fdadd" `
	// FMINACTIVE  string      `gorm:"column:FMINACTIVE;" json:"fminactive"  form:"fminactive" `
	// FCBUDCHART  string      `gorm:"column:FCBUDCHART;" json:"fcbudchart"  form:"fcbudchart" `
	// FCLID       string      `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	// FCPDCAT     string      `gorm:"column:FCPDCAT;" json:"fcpdcat"  form:"fcpdcat" `
	// FCCOSTBASE  string      `gorm:"column:FCCOSTBASE;" json:"fccostbase"  form:"fccostbase" `
	// FCSECT      string      `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	// FCCUST      string      `gorm:"column:FCCUST;" json:"fccust"  form:"fccust" `
	// FTLASTEDIT  string      `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	// FCPRTYPE    string      `gorm:"column:FCPRTYPE;" json:"fcprtype"  form:"fcprtype" `
	// FNFIXCOST   float64     `gorm:"column:FNFIXCOST;" json:"fnfixcost"  form:"fnfixcost" `
	// FNVARCOST   float64     `gorm:"column:FNVARCOST;" json:"fnvarcost"  form:"fnvarcost" `
	// FNDMCOST    float64     `gorm:"column:FNDMCOST;" json:"fndmcost"  form:"fndmcost" `
	// FNDLCOST    float64     `gorm:"column:FNDLCOST;" json:"fndlcost"  form:"fndlcost" `
	// FNOHCOST    float64     `gorm:"column:FNOHCOST;" json:"fnohcost"  form:"fnohcost" `
	// FCSWHOUSE   string      `gorm:"column:FCSWHOUSE;" json:"fcswhouse"  form:"fcswhouse" `
	// FCBWHOUSE   string      `gorm:"column:FCBWHOUSE;" json:"fcbwhouse"  form:"fcbwhouse" `
	// FCU1STATUS  string      `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	// FCU2STATUS  string      `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	// FCDTYPE1    string      `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FCDTYPE2    string      `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FNU1CNT     float64     `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FNU2CNT     float64     `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FMERRMSG    string      `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	// FCCREATEAP  string      `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	// FCACCWIP    string      `gorm:"column:FCACCWIP;" json:"fcaccwip"  form:"fcaccwip" `
	// FCMRPXFER   string      `gorm:"column:FCMRPXFER;" json:"fcmrpxfer"  form:"fcmrpxfer" `
	// FCACCCONSP  string      `gorm:"column:FCACCCONSP;" json:"fcaccconsp"  form:"fcaccconsp" `
	// FCU1ACC     string      `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCDATAIMP   string      `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	// FCCOLSEND   string      `gorm:"column:FCCOLSEND;" json:"fccolsend"  form:"fccolsend" `
	// FCMANFLAG   string      `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	// FCADDAPVBY  string      `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	// FCEDTAPVBY  string      `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	// FCDELAPVBY  string      `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	// FCISUSED    string      `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	// FCCREATEBY  string      `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	// FCCORRECTB  string      `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	// FNWIDTH     float64     `gorm:"column:FNWIDTH;" json:"fnwidth"  form:"fnwidth" `
	// FNLENGTH    float64     `gorm:"column:FNLENGTH;" json:"fnlength"  form:"fnlength" `
	// FNHEIGHT    float64     `gorm:"column:FNHEIGHT;" json:"fnheight"  form:"fnheight" `
	// FNVOLUME    float64     `gorm:"column:FNVOLUME;" json:"fnvolume"  form:"fnvolume" `
	// FNTRANSEXP  float64     `gorm:"column:FNTRANSEXP;" json:"fntransexp"  form:"fntransexp" `
	// FCU3STATUS  string      `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	// FCDTYPE3    string      `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FNU3CNT     float64     `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FCU4STATUS  string      `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	// FCDTYPE4    string      `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FNU4CNT     float64     `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FCU5STATUS  string      `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	// FCDTYPE5    string      `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FNU5CNT     float64     `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FCU6STATUS  string      `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	// FCDTYPE6    string      `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FNU6CNT     float64     `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FCU7STATUS  string      `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	// FCDTYPE7    string      `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FNU7CNT     float64     `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FCU8STATUS  string      `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	// FCDTYPE8    string      `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FNU8CNT     float64     `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FCU9STATUS  string      `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	// FCDTYPE9    string      `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	// FNU9CNT     float64     `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	// FCSTCTRLST  string      `gorm:"column:FCSTCTRLST;" json:"fcstctrlst"  form:"fcstctrlst" `
	// FCGID       string      `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	// FCBARCODE   string      `gorm:"column:FCBARCODE;" json:"fcbarcode"  form:"fcbarcode" `
	// FTSRCUPD    time.Time   `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
	// FCSRCUPD    string      `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FMEXTRATAG  string      `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	// FCIUM       string      `gorm:"column:FCIUM;" json:"fcium"  form:"fcium" `
	// FNIUMQTY    float64     `gorm:"column:FNIUMQTY;" json:"fniumqty"  form:"fniumqty" `
	// FCSTIUM     string      `gorm:"column:FCSTIUM;" json:"fcstium"  form:"fcstium" `
	// FNSTIUMQTY  float64     `gorm:"column:FNSTIUMQTY;" json:"fnstiumqty"  form:"fnstiumqty" `
	// FCUSEPDSER  string      `gorm:"column:FCUSEPDSER;" json:"fcusepdser"  form:"fcusepdser" `
	// FCORGCODE   string      `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCCUACC     string      `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCAPPNAME   string      `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	ProductType ProductType `gorm:"foreignKey:FCTYPE;references:FCCODE;" json:"product_type"`
	Unit        Unit        `gorm:"foreignKey:FCUM;references:FCSKID;" json:"product_unit"`
}

func (Product) TableName() string {
	return "PROD"
}

func (obj *Product) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}

type ProductType struct {
	FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCCODE    string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCNAME    string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2   string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	// FCNOTSHOW  string    `gorm:"column:FCNOTSHOW;" json:"fcnotshow"  form:"fcnotshow" `
	// FCDESC     string    `gorm:"column:FCDESC;" json:"fcdesc"  form:"fcdesc" `
	// FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	// FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	// FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	// FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	// FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	// FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	// FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	// FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	// FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	// FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	// FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	// FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	// FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	// FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	// FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	// FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	// FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	// FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	// FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	// FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	// FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	// FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	// FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	// FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	// FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	// FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	// FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	// FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	// FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	// FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	// FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (ProductType) TableName() string {
	return "PRODTYPE"
}

func (obj *ProductType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}

type Unit struct {
	// FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	// FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	// FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	// FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	// FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	// FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCCORP  string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCODE  string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCNAME  string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2 string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCTYPE  string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	// FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	// FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	// FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	// FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	// FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	// FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	// FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	// FTLASTEDIT time.Time    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	// FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	// FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	// FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	// FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	// FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	// FCCOLSEND  string    `gorm:"column:FCCOLSEND;" json:"fccolsend"  form:"fccolsend" `
	// FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	// FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	// FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	// FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	// FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	// FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	// FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	// FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	// FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	// FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	// FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	// FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	// FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	// FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	// FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	// FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	// FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	// FCQCRDUNIT string    `gorm:"column:FCQCRDUNIT;" json:"fcqcrdunit"  form:"fcqcrdunit" `
	// FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
	// FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	// FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Unit) TableName() string {
	return "UM"
}

func (obj *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
