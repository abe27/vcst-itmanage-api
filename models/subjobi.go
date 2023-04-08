package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Subjobi struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDTYPE1 string `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
     FCDTYPE2 string `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCLID string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPROD string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
     FCPRODTYPE string `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSUBJOBH string `gorm:"column:FCSUBJOBH;" json:"fcsubjobh"  form:"fcsubjobh" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
     FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUM string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
     FCUMSTD string `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
     FMERRMSG string `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNLBORCOST string `gorm:"column:FNLBORCOST;" json:"fnlborcost"  form:"fnlborcost" `
     FNPRODCOST string `gorm:"column:FNPRODCOST;" json:"fnprodcost"  form:"fnprodcost" `
     FNQTY string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
     FNU1CNT string `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
     FNU2CNT string `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
     FNUMQTY string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Subjobi) TableName() string{
return "SUBJOBI"
}

func (obj *Subjobi) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
