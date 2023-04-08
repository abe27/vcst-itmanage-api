package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Vbudkpi struct{
     FCACCHART string `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCTRGUM string `gorm:"column:FCTRGUM;" json:"fctrgum"  form:"fctrgum" `
     FCTYPE string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FCVBUDRQH string `gorm:"column:FCVBUDRQH;" json:"fcvbudrqh"  form:"fcvbudrqh" `
     FCVBUDRQI string `gorm:"column:FCVBUDRQI;" json:"fcvbudrqi"  form:"fcvbudrqi" `
     FCYEAR string `gorm:"column:FCYEAR;" json:"fcyear"  form:"fcyear" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FMREMARK string `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
     FMREMARK01 string `gorm:"column:FMREMARK01;" json:"fmremark01"  form:"fmremark01" `
     FMREMARK02 string `gorm:"column:FMREMARK02;" json:"fmremark02"  form:"fmremark02" `
     FMREMARK03 string `gorm:"column:FMREMARK03;" json:"fmremark03"  form:"fmremark03" `
     FMREMARK04 string `gorm:"column:FMREMARK04;" json:"fmremark04"  form:"fmremark04" `
     FMREMARK05 string `gorm:"column:FMREMARK05;" json:"fmremark05"  form:"fmremark05" `
     FMREMARK06 string `gorm:"column:FMREMARK06;" json:"fmremark06"  form:"fmremark06" `
     FMREMARK07 string `gorm:"column:FMREMARK07;" json:"fmremark07"  form:"fmremark07" `
     FMREMARK08 string `gorm:"column:FMREMARK08;" json:"fmremark08"  form:"fmremark08" `
     FMREMARK09 string `gorm:"column:FMREMARK09;" json:"fmremark09"  form:"fmremark09" `
     FMREMARK10 string `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
     FMREMARK11 string `gorm:"column:FMREMARK11;" json:"fmremark11"  form:"fmremark11" `
     FMREMARK12 string `gorm:"column:FMREMARK12;" json:"fmremark12"  form:"fmremark12" `
     FNACTQTY string `gorm:"column:FNACTQTY;" json:"fnactqty"  form:"fnactqty" `
     FNQTY01 string `gorm:"column:FNQTY01;" json:"fnqty01"  form:"fnqty01" `
     FNQTY02 string `gorm:"column:FNQTY02;" json:"fnqty02"  form:"fnqty02" `
     FNQTY03 string `gorm:"column:FNQTY03;" json:"fnqty03"  form:"fnqty03" `
     FNQTY04 string `gorm:"column:FNQTY04;" json:"fnqty04"  form:"fnqty04" `
     FNQTY05 string `gorm:"column:FNQTY05;" json:"fnqty05"  form:"fnqty05" `
     FNQTY06 string `gorm:"column:FNQTY06;" json:"fnqty06"  form:"fnqty06" `
     FNQTY07 string `gorm:"column:FNQTY07;" json:"fnqty07"  form:"fnqty07" `
     FNQTY08 string `gorm:"column:FNQTY08;" json:"fnqty08"  form:"fnqty08" `
     FNQTY09 string `gorm:"column:FNQTY09;" json:"fnqty09"  form:"fnqty09" `
     FNQTY10 string `gorm:"column:FNQTY10;" json:"fnqty10"  form:"fnqty10" `
     FNQTY11 string `gorm:"column:FNQTY11;" json:"fnqty11"  form:"fnqty11" `
     FNQTY12 string `gorm:"column:FNQTY12;" json:"fnqty12"  form:"fnqty12" `
     FNTRGQTY string `gorm:"column:FNTRGQTY;" json:"fntrgqty"  form:"fntrgqty" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Vbudkpi) TableName() string{
return "VBUDKPI"
}

func (obj *Vbudkpi) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
