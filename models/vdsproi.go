package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Vdsproi struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPROD string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUM string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
     FCUMSTD string `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FCVDSPROH string `gorm:"column:FCVDSPROH;" json:"fcvdsproh"  form:"fcvdsproh" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNAMT string `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
     FNQTY string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
     FNUMQTY string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Vdsproi) TableName() string{
return "VDSPROI"
}

func (obj *Vdsproi) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}