package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Locmovi1 struct{
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
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
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCGAG string `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
     FCGVPOLICY string `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
     FCIOTYPE string `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLOCA1 string `gorm:"column:FCLOCA1;" json:"fcloca1"  form:"fcloca1" `
     FCLOCMOVH1 string `gorm:"column:FCLOCMOVH1;" json:"fclocmovh1"  form:"fclocmovh1" `
     FCLOT string `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPLANT string `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
     FCPROD string `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
     FCPRODTYPE string `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
     FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
     FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
     FCRFTYPE string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCSTORE string `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
     FCSTUMSTD string `gorm:"column:FCSTUMSTD;" json:"fcstumstd"  form:"fcstumstd" `
     FCSUBBR string `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
     FCTIMESTAM string `gorm:"column:FCTIMESTAM;" json:"fctimestam"  form:"fctimestam" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FCWHOUSE string `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
     FDDATE string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
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
     FNQTYATDAT string `gorm:"column:FNQTYATDAT;" json:"fnqtyatdat"  form:"fnqtyatdat" `
     FNSTDQTY string `gorm:"column:FNSTDQTY;" json:"fnstdqty"  form:"fnstdqty" `
     FNSTDSTQTY string `gorm:"column:FNSTDSTQTY;" json:"fnstdstqty"  form:"fnstdstqty" `
     FNSTQTYATD string `gorm:"column:FNSTQTYATD;" json:"fnstqtyatd"  form:"fnstqtyatd" `
     FNUMQTY string `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Locmovi1) TableName() string{
return "LOCMOVI1"
}

func (obj *Locmovi1) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
