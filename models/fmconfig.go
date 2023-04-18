package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Fmconfig struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCON000   string    `gorm:"column:FCCON000;" json:"fccon000"  form:"fccon000" `
	FCCON001   string    `gorm:"column:FCCON001;" json:"fccon001"  form:"fccon001" `
	FCCON002   string    `gorm:"column:FCCON002;" json:"fccon002"  form:"fccon002" `
	FCCON003   string    `gorm:"column:FCCON003;" json:"fccon003"  form:"fccon003" `
	FCCON004   string    `gorm:"column:FCCON004;" json:"fccon004"  form:"fccon004" `
	FCCON005   string    `gorm:"column:FCCON005;" json:"fccon005"  form:"fccon005" `
	FCCON006   string    `gorm:"column:FCCON006;" json:"fccon006"  form:"fccon006" `
	FCCON007   string    `gorm:"column:FCCON007;" json:"fccon007"  form:"fccon007" `
	FCCON008   string    `gorm:"column:FCCON008;" json:"fccon008"  form:"fccon008" `
	FCCON009   string    `gorm:"column:FCCON009;" json:"fccon009"  form:"fccon009" `
	FCCON010   string    `gorm:"column:FCCON010;" json:"fccon010"  form:"fccon010" `
	FCCON011   string    `gorm:"column:FCCON011;" json:"fccon011"  form:"fccon011" `
	FCCON012   string    `gorm:"column:FCCON012;" json:"fccon012"  form:"fccon012" `
	FCCON013   string    `gorm:"column:FCCON013;" json:"fccon013"  form:"fccon013" `
	FCCON014   string    `gorm:"column:FCCON014;" json:"fccon014"  form:"fccon014" `
	FCCON015   string    `gorm:"column:FCCON015;" json:"fccon015"  form:"fccon015" `
	FCCON016   string    `gorm:"column:FCCON016;" json:"fccon016"  form:"fccon016" `
	FCCON017   string    `gorm:"column:FCCON017;" json:"fccon017"  form:"fccon017" `
	FCCON018   string    `gorm:"column:FCCON018;" json:"fccon018"  form:"fccon018" `
	FCCON019   string    `gorm:"column:FCCON019;" json:"fccon019"  form:"fccon019" `
	FCCON020   string    `gorm:"column:FCCON020;" json:"fccon020"  form:"fccon020" `
	FCCON021   string    `gorm:"column:FCCON021;" json:"fccon021"  form:"fccon021" `
	FCCON022   string    `gorm:"column:FCCON022;" json:"fccon022"  form:"fccon022" `
	FCCON023   string    `gorm:"column:FCCON023;" json:"fccon023"  form:"fccon023" `
	FCCON024   string    `gorm:"column:FCCON024;" json:"fccon024"  form:"fccon024" `
	FCCON025   string    `gorm:"column:FCCON025;" json:"fccon025"  form:"fccon025" `
	FCCON026   string    `gorm:"column:FCCON026;" json:"fccon026"  form:"fccon026" `
	FCCON027   string    `gorm:"column:FCCON027;" json:"fccon027"  form:"fccon027" `
	FCCON028   string    `gorm:"column:FCCON028;" json:"fccon028"  form:"fccon028" `
	FCCON029   string    `gorm:"column:FCCON029;" json:"fccon029"  form:"fccon029" `
	FCCON030   string    `gorm:"column:FCCON030;" json:"fccon030"  form:"fccon030" `
	FCCON031   string    `gorm:"column:FCCON031;" json:"fccon031"  form:"fccon031" `
	FCCON032   string    `gorm:"column:FCCON032;" json:"fccon032"  form:"fccon032" `
	FCCON033   string    `gorm:"column:FCCON033;" json:"fccon033"  form:"fccon033" `
	FCCON034   string    `gorm:"column:FCCON034;" json:"fccon034"  form:"fccon034" `
	FCCON035   string    `gorm:"column:FCCON035;" json:"fccon035"  form:"fccon035" `
	FCCON036   string    `gorm:"column:FCCON036;" json:"fccon036"  form:"fccon036" `
	FCCON037   string    `gorm:"column:FCCON037;" json:"fccon037"  form:"fccon037" `
	FCCON038   string    `gorm:"column:FCCON038;" json:"fccon038"  form:"fccon038" `
	FCCON039   string    `gorm:"column:FCCON039;" json:"fccon039"  form:"fccon039" `
	FCCON040   string    `gorm:"column:FCCON040;" json:"fccon040"  form:"fccon040" `
	FCCON041   string    `gorm:"column:FCCON041;" json:"fccon041"  form:"fccon041" `
	FCCON042   string    `gorm:"column:FCCON042;" json:"fccon042"  form:"fccon042" `
	FCCON043   string    `gorm:"column:FCCON043;" json:"fccon043"  form:"fccon043" `
	FCCON044   string    `gorm:"column:FCCON044;" json:"fccon044"  form:"fccon044" `
	FCCON045   string    `gorm:"column:FCCON045;" json:"fccon045"  form:"fccon045" `
	FCCON046   string    `gorm:"column:FCCON046;" json:"fccon046"  form:"fccon046" `
	FCCON047   string    `gorm:"column:FCCON047;" json:"fccon047"  form:"fccon047" `
	FCCON048   string    `gorm:"column:FCCON048;" json:"fccon048"  form:"fccon048" `
	FCCON049   string    `gorm:"column:FCCON049;" json:"fccon049"  form:"fccon049" `
	FCCON050   string    `gorm:"column:FCCON050;" json:"fccon050"  form:"fccon050" `
	FCCON051   string    `gorm:"column:FCCON051;" json:"fccon051"  form:"fccon051" `
	FCCON052   string    `gorm:"column:FCCON052;" json:"fccon052"  form:"fccon052" `
	FCCON053   string    `gorm:"column:FCCON053;" json:"fccon053"  form:"fccon053" `
	FCCON054   string    `gorm:"column:FCCON054;" json:"fccon054"  form:"fccon054" `
	FCCON055   string    `gorm:"column:FCCON055;" json:"fccon055"  form:"fccon055" `
	FCCON056   string    `gorm:"column:FCCON056;" json:"fccon056"  form:"fccon056" `
	FCCON057   string    `gorm:"column:FCCON057;" json:"fccon057"  form:"fccon057" `
	FCCON058   string    `gorm:"column:FCCON058;" json:"fccon058"  form:"fccon058" `
	FCCON059   string    `gorm:"column:FCCON059;" json:"fccon059"  form:"fccon059" `
	FCCON060   string    `gorm:"column:FCCON060;" json:"fccon060"  form:"fccon060" `
	FCCON061   string    `gorm:"column:FCCON061;" json:"fccon061"  form:"fccon061" `
	FCCON062   string    `gorm:"column:FCCON062;" json:"fccon062"  form:"fccon062" `
	FCCON063   string    `gorm:"column:FCCON063;" json:"fccon063"  form:"fccon063" `
	FCCON064   string    `gorm:"column:FCCON064;" json:"fccon064"  form:"fccon064" `
	FCCON065   string    `gorm:"column:FCCON065;" json:"fccon065"  form:"fccon065" `
	FCCON066   string    `gorm:"column:FCCON066;" json:"fccon066"  form:"fccon066" `
	FCCON067   string    `gorm:"column:FCCON067;" json:"fccon067"  form:"fccon067" `
	FCCON068   string    `gorm:"column:FCCON068;" json:"fccon068"  form:"fccon068" `
	FCCON069   string    `gorm:"column:FCCON069;" json:"fccon069"  form:"fccon069" `
	FCCON070   string    `gorm:"column:FCCON070;" json:"fccon070"  form:"fccon070" `
	FCCON141   string    `gorm:"column:FCCON141;" json:"fccon141"  form:"fccon141" `
	FCCON142   string    `gorm:"column:FCCON142;" json:"fccon142"  form:"fccon142" `
	FCCON143   string    `gorm:"column:FCCON143;" json:"fccon143"  form:"fccon143" `
	FCCON144   string    `gorm:"column:FCCON144;" json:"fccon144"  form:"fccon144" `
	FCCON145   string    `gorm:"column:FCCON145;" json:"fccon145"  form:"fccon145" `
	FCCON146   string    `gorm:"column:FCCON146;" json:"fccon146"  form:"fccon146" `
	FCCON147   string    `gorm:"column:FCCON147;" json:"fccon147"  form:"fccon147" `
	FCCON148   string    `gorm:"column:FCCON148;" json:"fccon148"  form:"fccon148" `
	FCCON149   string    `gorm:"column:FCCON149;" json:"fccon149"  form:"fccon149" `
	FCCON150   string    `gorm:"column:FCCON150;" json:"fccon150"  form:"fccon150" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCENCOTYPE string    `gorm:"column:FCENCOTYPE;" json:"fcencotype"  form:"fcencotype" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT string    `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string    `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Fmconfig) TableName() string {
	return "FMCONFIG"
}

func (obj *Fmconfig) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
