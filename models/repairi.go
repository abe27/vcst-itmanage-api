package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Repairi struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCAPPROVEB string `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
     FCBOICARD string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
     FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCOOR string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
     FCDISCSTR string `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCFORMULAS string `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
     FCGAG string `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLOT string `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPFORMULA string `gorm:"column:FCPFORMULA;" json:"fcpformula"  form:"fcpformula" `
     FCPROD string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
     FCPRODTYPE string `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
     FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
     FCREFPDTYP string `gorm:"column:FCREFPDTYP;" json:"fcrefpdtyp"  form:"fcrefpdtyp" `
     FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
     FCREPAIRH string `gorm:"column:FCREPAIRH;" json:"fcrepairh"  form:"fcrepairh" `
     FCROOTSEQ string `gorm:"column:FCROOTSEQ;" json:"fcrootseq"  form:"fcrootseq" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSHOWCOMP string `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCSTEP string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUM string `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
     FCUMSTD string `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FCVATISOUT string `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
     FCVATTYPE string `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
     FCWHOUSE string `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
     FDAPPROVE string `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" `
     FDDATE string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
     FDDELIVERY string `gorm:"column:FDDELIVERY;" json:"fddelivery"  form:"fddelivery" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMREMARK string `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
     FMREMARK10 string `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
     FMREMARK2 string `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
     FMREMARK3 string `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
     FMREMARK4 string `gorm:"column:FMREMARK4;" json:"fmremark4"  form:"fmremark4" `
     FMREMARK5 string `gorm:"column:FMREMARK5;" json:"fmremark5"  form:"fmremark5" `
     FMREMARK6 string `gorm:"column:FMREMARK6;" json:"fmremark6"  form:"fmremark6" `
     FMREMARK7 string `gorm:"column:FMREMARK7;" json:"fmremark7"  form:"fmremark7" `
     FMREMARK8 string `gorm:"column:FMREMARK8;" json:"fmremark8"  form:"fmremark8" `
     FMREMARK9 string `gorm:"column:FMREMARK9;" json:"fmremark9"  form:"fmremark9" `
     FMTEXTKE string `gorm:"column:FMTEXTKE;" json:"fmtextke"  form:"fmtextke" `
     FNBACKQTY string `gorm:"column:FNBACKQTY;" json:"fnbackqty"  form:"fnbackqty" `
     FNDISCAMT string `gorm:"column:FNDISCAMT;" json:"fndiscamt"  form:"fndiscamt" `
     FNDISCAMTK string `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
     FNPRICE string `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
     FNPRICEKE string `gorm:"column:FNPRICEKE;" json:"fnpriceke"  form:"fnpriceke" `
     FNQTY string `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
     FNUMQTY string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
     FNVATAMT string `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
     FNVATRATE string `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
     FNXRATE string `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Repairi) TableName() string{
return "REPAIRI"
}

func (obj *Repairi) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}