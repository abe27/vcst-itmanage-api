package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Adjboqi struct{
     FCADJBOQH string `gorm:"column:FCADJBOQH;" json:"fcadjboqh"  form:"fcadjboqh" `
     FCAPPNAME string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
     FCBOICARD string `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
     FCBOIGROUP string `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
     FCCUACC string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
     FCDATAIMP string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCESTBOQH string `gorm:"column:FCESTBOQH;" json:"fcestboqh"  form:"fcestboqh" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCMAINJOBH string `gorm:"column:FCMAINJOBH;" json:"fcmainjobh"  form:"fcmainjobh" `
     FCMAINUM string `gorm:"column:FCMAINUM;" json:"fcmainum"  form:"fcmainum" `
     FCORGCODE string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
     FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSEQ string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
     FCSJOBUM string `gorm:"column:FCSJOBUM;" json:"fcsjobum"  form:"fcsjobum" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
     FCSUBJOBH string `gorm:"column:FCSUBJOBH;" json:"fcsubjobh"  form:"fcsubjobh" `
     FCU1ACC string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
     FNINVLBAMT string `gorm:"column:FNINVLBAMT;" json:"fninvlbamt"  form:"fninvlbamt" `
     FNINVPDAMT string `gorm:"column:FNINVPDAMT;" json:"fninvpdamt"  form:"fninvpdamt" `
     FNMAINQTY string `gorm:"column:FNMAINQTY;" json:"fnmainqty"  form:"fnmainqty" `
     FNPOLBAMT string `gorm:"column:FNPOLBAMT;" json:"fnpolbamt"  form:"fnpolbamt" `
     FNPOPDAMT string `gorm:"column:FNPOPDAMT;" json:"fnpopdamt"  form:"fnpopdamt" `
     FNPRLBAMT string `gorm:"column:FNPRLBAMT;" json:"fnprlbamt"  form:"fnprlbamt" `
     FNPRPDAMT string `gorm:"column:FNPRPDAMT;" json:"fnprpdamt"  form:"fnprpdamt" `
     FNSJOBLBOR string `gorm:"column:FNSJOBLBOR;" json:"fnsjoblbor"  form:"fnsjoblbor" `
     FNSJOBPROD string `gorm:"column:FNSJOBPROD;" json:"fnsjobprod"  form:"fnsjobprod" `
     FNSJOBQTY string `gorm:"column:FNSJOBQTY;" json:"fnsjobqty"  form:"fnsjobqty" `
     FNTOTLABOR string `gorm:"column:FNTOTLABOR;" json:"fntotlabor"  form:"fntotlabor" `
     FNTOTPROD string `gorm:"column:FNTOTPROD;" json:"fntotprod"  form:"fntotprod" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
     FTSRCUPD string `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}
func (Adjboqi) TableName() string{
return "ADJBOQI"
}

func (obj *Adjboqi) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
