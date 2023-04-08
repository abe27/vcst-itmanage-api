package models

    import (
        "time"

        g "github.com/matoous/go-nanoid/v2"
        "gorm.io/gorm"
    )
    
type Worktab struct{
     FCAFCALOT string `gorm:"column:FCAFCALOT;" json:"fcafcalot"  form:"fcafcalot" `
     FCAFOGR string `gorm:"column:FCAFOGR;" json:"fcafogr"  form:"fcafogr" `
     FCAFRATEBY string `gorm:"column:FCAFRATEBY;" json:"fcafrateby"  form:"fcafrateby" `
     FCAFSALADJ string `gorm:"column:FCAFSALADJ;" json:"fcafsaladj"  form:"fcafsaladj" `
     FCAFTAEMP string `gorm:"column:FCAFTAEMP;" json:"fcaftaemp"  form:"fcaftaemp" `
     FCBEG1FLG string `gorm:"column:FCBEG1FLG;" json:"fcbeg1flg"  form:"fcbeg1flg" `
     FCBEG1TYP string `gorm:"column:FCBEG1TYP;" json:"fcbeg1typ"  form:"fcbeg1typ" `
     FCBEG2FLG string `gorm:"column:FCBEG2FLG;" json:"fcbeg2flg"  form:"fcbeg2flg" `
     FCBEG2TYP string `gorm:"column:FCBEG2TYP;" json:"fcbeg2typ"  form:"fcbeg2typ" `
     FCBEGAFFLG string `gorm:"column:FCBEGAFFLG;" json:"fcbegafflg"  form:"fcbegafflg" `
     FCBEGAFTYP string `gorm:"column:FCBEGAFTYP;" json:"fcbegaftyp"  form:"fcbegaftyp" `
     FCBEGBFFLG string `gorm:"column:FCBEGBFFLG;" json:"fcbegbfflg"  form:"fcbegbfflg" `
     FCBEGBFTYP string `gorm:"column:FCBEGBFTYP;" json:"fcbegbftyp"  form:"fcbegbftyp" `
     FCBFCALOT string `gorm:"column:FCBFCALOT;" json:"fcbfcalot"  form:"fcbfcalot" `
     FCBFOGR string `gorm:"column:FCBFOGR;" json:"fcbfogr"  form:"fcbfogr" `
     FCBFRATEBY string `gorm:"column:FCBFRATEBY;" json:"fcbfrateby"  form:"fcbfrateby" `
     FCBFSALADJ string `gorm:"column:FCBFSALADJ;" json:"fcbfsaladj"  form:"fcbfsaladj" `
     FCBFTAEMP string `gorm:"column:FCBFTAEMP;" json:"fcbftaemp"  form:"fcbftaemp" `
     FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
     FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
     FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
     FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
     FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
     FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
     FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
     FCEMPL string `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
     FCEMPLFIX string `gorm:"column:FCEMPLFIX;" json:"fcemplfix"  form:"fcemplfix" `
     FCEMPLTYPE string `gorm:"column:FCEMPLTYPE;" json:"fcempltype"  form:"fcempltype" `
     FCEMPLX1 string `gorm:"column:FCEMPLX1;" json:"fcemplx1"  form:"fcemplx1" `
     FCEND1FLG string `gorm:"column:FCEND1FLG;" json:"fcend1flg"  form:"fcend1flg" `
     FCEND1TYP string `gorm:"column:FCEND1TYP;" json:"fcend1typ"  form:"fcend1typ" `
     FCEND2FLG string `gorm:"column:FCEND2FLG;" json:"fcend2flg"  form:"fcend2flg" `
     FCEND2TYP string `gorm:"column:FCEND2TYP;" json:"fcend2typ"  form:"fcend2typ" `
     FCENDAFFLG string `gorm:"column:FCENDAFFLG;" json:"fcendafflg"  form:"fcendafflg" `
     FCENDAFTYP string `gorm:"column:FCENDAFTYP;" json:"fcendaftyp"  form:"fcendaftyp" `
     FCENDBFFLG string `gorm:"column:FCENDBFFLG;" json:"fcendbfflg"  form:"fcendbfflg" `
     FCENDBFTYP string `gorm:"column:FCENDBFTYP;" json:"fcendbftyp"  form:"fcendbftyp" `
     FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
     FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
     FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
     FCRANK string `gorm:"column:FCRANK;" json:"fcrank"  form:"fcrank" `
     FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
     FCSELTAG string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
     FCSHIFT string `gorm:"column:FCSHIFT;" json:"fcshift"  form:"fcshift" `
     FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
     FCSSESKID string `gorm:"column:FCSSESKID;" json:"fcsseskid"  form:"fcsseskid" `
     FCSTAT string `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
     FCSTDCALOT string `gorm:"column:FCSTDCALOT;" json:"fcstdcalot"  form:"fcstdcalot" `
     FCSTDOGR string `gorm:"column:FCSTDOGR;" json:"fcstdogr"  form:"fcstdogr" `
     FCSTDRATEB string `gorm:"column:FCSTDRATEB;" json:"fcstdrateb"  form:"fcstdrateb" `
     FCSTDSALAD string `gorm:"column:FCSTDSALAD;" json:"fcstdsalad"  form:"fcstdsalad" `
     FCSTDTAEMP string `gorm:"column:FCSTDTAEMP;" json:"fcstdtaemp"  form:"fcstdtaemp" `
     FCSTEP string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
     FCTAEMP string `gorm:"column:FCTAEMP;" json:"fctaemp"  form:"fctaemp" `
     FCTYPE string `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
     FCUDATE string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
     FCUTIME string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
     FDDATE string `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" `
     FIMILLISEC string `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
     FNAFDIF string `gorm:"column:FNAFDIF;" json:"fnafdif"  form:"fnafdif" `
     FNAFOTADJ string `gorm:"column:FNAFOTADJ;" json:"fnafotadj"  form:"fnafotadj" `
     FNAFRATE string `gorm:"column:FNAFRATE;" json:"fnafrate"  form:"fnafrate" `
     FNBEG1DIF string `gorm:"column:FNBEG1DIF;" json:"fnbeg1dif"  form:"fnbeg1dif" `
     FNBEG2DIF string `gorm:"column:FNBEG2DIF;" json:"fnbeg2dif"  form:"fnbeg2dif" `
     FNBEGAFDIF string `gorm:"column:FNBEGAFDIF;" json:"fnbegafdif"  form:"fnbegafdif" `
     FNBEGBFDIF string `gorm:"column:FNBEGBFDIF;" json:"fnbegbfdif"  form:"fnbegbfdif" `
     FNBFDIF string `gorm:"column:FNBFDIF;" json:"fnbfdif"  form:"fnbfdif" `
     FNBFOTADJ string `gorm:"column:FNBFOTADJ;" json:"fnbfotadj"  form:"fnbfotadj" `
     FNBFRATE string `gorm:"column:FNBFRATE;" json:"fnbfrate"  form:"fnbfrate" `
     FNEND1DIF string `gorm:"column:FNEND1DIF;" json:"fnend1dif"  form:"fnend1dif" `
     FNEND2DIF string `gorm:"column:FNEND2DIF;" json:"fnend2dif"  form:"fnend2dif" `
     FNENDAFDIF string `gorm:"column:FNENDAFDIF;" json:"fnendafdif"  form:"fnendafdif" `
     FNENDBFDIF string `gorm:"column:FNENDBFDIF;" json:"fnendbfdif"  form:"fnendbfdif" `
     FNSTD1DIF string `gorm:"column:FNSTD1DIF;" json:"fnstd1dif"  form:"fnstd1dif" `
     FNSTD2DIF string `gorm:"column:FNSTD2DIF;" json:"fnstd2dif"  form:"fnstd2dif" `
     FNSTDDIF string `gorm:"column:FNSTDDIF;" json:"fnstddif"  form:"fnstddif" `
     FNSTDOTADJ string `gorm:"column:FNSTDOTADJ;" json:"fnstdotadj"  form:"fnstdotadj" `
     FNSTDRATE string `gorm:"column:FNSTDRATE;" json:"fnstdrate"  form:"fnstdrate" `
     FTBEG1 string `gorm:"column:FTBEG1;" json:"ftbeg1"  form:"ftbeg1" `
     FTBEG1REG string `gorm:"column:FTBEG1REG;" json:"ftbeg1reg"  form:"ftbeg1reg" `
     FTBEG2 string `gorm:"column:FTBEG2;" json:"ftbeg2"  form:"ftbeg2" `
     FTBEG2REG string `gorm:"column:FTBEG2REG;" json:"ftbeg2reg"  form:"ftbeg2reg" `
     FTBEGAF string `gorm:"column:FTBEGAF;" json:"ftbegaf"  form:"ftbegaf" `
     FTBEGAFREG string `gorm:"column:FTBEGAFREG;" json:"ftbegafreg"  form:"ftbegafreg" `
     FTBEGBF string `gorm:"column:FTBEGBF;" json:"ftbegbf"  form:"ftbegbf" `
     FTBEGBFREG string `gorm:"column:FTBEGBFREG;" json:"ftbegbfreg"  form:"ftbegbfreg" `
     FTBEGREG string `gorm:"column:FTBEGREG;" json:"ftbegreg"  form:"ftbegreg" `
     FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
     FTEND1 string `gorm:"column:FTEND1;" json:"ftend1"  form:"ftend1" `
     FTEND1REG string `gorm:"column:FTEND1REG;" json:"ftend1reg"  form:"ftend1reg" `
     FTEND2 string `gorm:"column:FTEND2;" json:"ftend2"  form:"ftend2" `
     FTEND2REG string `gorm:"column:FTEND2REG;" json:"ftend2reg"  form:"ftend2reg" `
     FTENDAF string `gorm:"column:FTENDAF;" json:"ftendaf"  form:"ftendaf" `
     FTENDAFREG string `gorm:"column:FTENDAFREG;" json:"ftendafreg"  form:"ftendafreg" `
     FTENDBF string `gorm:"column:FTENDBF;" json:"ftendbf"  form:"ftendbf" `
     FTENDBFREG string `gorm:"column:FTENDBFREG;" json:"ftendbfreg"  form:"ftendbfreg" `
     FTENDREG string `gorm:"column:FTENDREG;" json:"ftendreg"  form:"ftendreg" `
     FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
     FTLASTUPD string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}
func (Worktab) TableName() string{
return "WORKTAB"
}

func (obj *Worktab) BeforeCreate(tx *gorm.DB) (err error){
id, _ := g.New(8)
obj.FCSKID = id
return}
