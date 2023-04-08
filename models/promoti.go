package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Promoti struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCCURRENCY string `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDISCSTR string `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCFORMULAS string `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
     FCITEMTYPE string `gorm:"column:FCITEMTYPE;" json:"fcitemtype"  form:"fcitemtype" `
     FCLIST string `gorm:"column:FCLIST;" json:"fclist"  form:"fclist" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPDGRP string `gorm:"column:FCPDGRP;" json:"fcpdgrp"  form:"fcpdgrp" `
     FCPROD string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
     FCPROMOTH string `gorm:"column:FCPROMOTH;" json:"fcpromoth"  form:"fcpromoth" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCTYPE string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUM string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMMEMDATA string `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
     FNDISCAMT string `gorm:"column:FNDISCAMT;" json:"fndiscamt"  form:"fndiscamt" `
     FNPOINT1 string `gorm:"column:FNPOINT1;" json:"fnpoint1"  form:"fnpoint1" `
     FNPOINT2 string `gorm:"column:FNPOINT2;" json:"fnpoint2"  form:"fnpoint2" `
     FNPOINT3 string `gorm:"column:FNPOINT3;" json:"fnpoint3"  form:"fnpoint3" `
     FNPRICE string `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
     FNPRICEKE string `gorm:"column:FNPRICEKE;" json:"fnpriceke"  form:"fnpriceke" `
     FNQTY string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
     FNUMQTY string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
     FTBEGDATE string `gorm:"column:FTBEGDATE;" json:"ftbegdate"  form:"ftbegdate" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTENDDATE string `gorm:"column:FTENDDATE;" json:"ftenddate"  form:"ftenddate" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Promoti) TableName() string{
return "PROMOTI"
}

func (obj *Promoti) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
