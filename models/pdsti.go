package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Pdsti struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCAUTOTY01 string `gorm:"column:FCAUTOTY01;" json:"fcautoty01"  form:"fcautoty01" `
     FCAUTOTY02 string `gorm:"column:FCAUTOTY02;" json:"fcautoty02"  form:"fcautoty02" `
     FCAUTOTY03 string `gorm:"column:FCAUTOTY03;" json:"fcautoty03"  form:"fcautoty03" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDRAWNO string `gorm:"column:FCDRAWNO;" json:"fcdrawno"  form:"fcdrawno" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCECONO string `gorm:"column:FCECONO;" json:"fcecono"  form:"fcecono" `
     FCIOTYPE string `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
     FCISMAINPD string `gorm:"column:FCISMAINPD;" json:"fcismainpd"  form:"fcismainpd" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCOPSEQ string `gorm:"column:FCOPSEQ;" json:"fcopseq"  form:"fcopseq" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPDSTH string `gorm:"column:FCPDSTH;" json:"fcpdsth"  form:"fcpdsth" `
     FCPROCURE string `gorm:"column:FCPROCURE;" json:"fcprocure"  form:"fcprocure" `
     FCPROD string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
     FCREVISION string `gorm:"column:FCREVISION;" json:"fcrevision"  form:"fcrevision" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTUM string `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
     FCSUBSTI01 string `gorm:"column:FCSUBSTI01;" json:"fcsubsti01"  form:"fcsubsti01" `
     FCSUBSTI02 string `gorm:"column:FCSUBSTI02;" json:"fcsubsti02"  form:"fcsubsti02" `
     FCSUBSTI03 string `gorm:"column:FCSUBSTI03;" json:"fcsubsti03"  form:"fcsubsti03" `
     FCSUBSTIT1 string `gorm:"column:FCSUBSTIT1;" json:"fcsubstit1"  form:"fcsubstit1" `
     FCSUBSTIT2 string `gorm:"column:FCSUBSTIT2;" json:"fcsubstit2"  form:"fcsubstit2" `
     FCTYPE string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUM string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
     FCUM01 string `gorm:"column:FCUM01;" json:"fcum01"  form:"fcum01" `
     FCUM02 string `gorm:"column:FCUM02;" json:"fcum02"  form:"fcum02" `
     FCUM03 string `gorm:"column:FCUM03;" json:"fcum03"  form:"fcum03" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FDEFFECT string `gorm:"column:FDEFFECT;" json:"fdeffect"  form:"fdeffect" `
     FDUSERUPD string `gorm:"column:FDUSERUPD;" json:"fduserupd"  form:"fduserupd" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNCONSQTY string `gorm:"column:FNCONSQTY;" json:"fnconsqty"  form:"fnconsqty" `
     FNCONSUNIT string `gorm:"column:FNCONSUNIT;" json:"fnconsunit"  form:"fnconsunit" `
     FNDAYSTART string `gorm:"column:FNDAYSTART;" json:"fndaystart"  form:"fndaystart" `
     FNFSTART string `gorm:"column:FNFSTART;" json:"fnfstart"  form:"fnfstart" `
     FNGROSSWG string `gorm:"column:FNGROSSWG;" json:"fngrosswg"  form:"fngrosswg" `
     FNISIGCOST string `gorm:"column:FNISIGCOST;" json:"fnisigcost"  form:"fnisigcost" `
     FNISMRP string `gorm:"column:FNISMRP;" json:"fnismrp"  form:"fnismrp" `
     FNISOVERHD string `gorm:"column:FNISOVERHD;" json:"fnisoverhd"  form:"fnisoverhd" `
     FNPACKSIZE string `gorm:"column:FNPACKSIZE;" json:"fnpacksize"  form:"fnpacksize" `
     FNPORTION string `gorm:"column:FNPORTION;" json:"fnportion"  form:"fnportion" `
     FNQTY string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
     FNQTY01 string `gorm:"column:FNQTY01;" json:"fnqty01"  form:"fnqty01" `
     FNQTY02 string `gorm:"column:FNQTY02;" json:"fnqty02"  form:"fnqty02" `
     FNQTY03 string `gorm:"column:FNQTY03;" json:"fnqty03"  form:"fnqty03" `
     FNSTDYEILD string `gorm:"column:FNSTDYEILD;" json:"fnstdyeild"  form:"fnstdyeild" `
     FNSTQTY string `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
     FNSTUMQTY string `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
     FNUMQTY string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
     FNUMQTY01 string `gorm:"column:FNUMQTY01;" json:"fnumqty01"  form:"fnumqty01" `
     FNUMQTY02 string `gorm:"column:FNUMQTY02;" json:"fnumqty02"  form:"fnumqty02" `
     FNUMQTY03 string `gorm:"column:FNUMQTY03;" json:"fnumqty03"  form:"fnumqty03" `
     FNYEILDPCN string `gorm:"column:FNYEILDPCN;" json:"fnyeildpcn"  form:"fnyeildpcn" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Pdsti) TableName() string{
return "PDSTI"
}

func (obj *Pdsti) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
