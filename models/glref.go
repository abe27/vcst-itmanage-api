package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Glref struct {
	// FCACSTEP       string     `gorm:"column:FCACSTEP;" json:"fcacstep"  form:"fcacstep" `
	// FCAPPNAME      string     `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	// FCAPPROVEB     string     `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCATSTEP string `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	// FCBOICARD      string     `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	// FCBOIGROUP     string     `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBOOK   string `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	// FCBRANCHNM     string     `gorm:"column:FCBRANCHNM;" json:"fcbranchnm"  form:"fcbranchnm" `
	// FCBRANCHNO     string     `gorm:"column:FCBRANCHNO;" json:"fcbranchno"  form:"fcbranchno" `
	// FCCANCELBY     string     `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	// FCCARRIER      string     `gorm:"column:FCCARRIER;" json:"fccarrier"  form:"fccarrier" `
	// FCCARRYTYP     string     `gorm:"column:FCCARRYTYP;" json:"fccarrytyp"  form:"fccarrytyp" `
	// FCCASHIER      string     `gorm:"column:FCCASHIER;" json:"fccashier"  form:"fccashier" `
	// FCCHKDT        string     `gorm:"column:FCCHKDT;" json:"fcchkdt"  form:"fcchkdt" `
	FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	// FCCONSIGN      string     `gorm:"column:FCCONSIGN;" json:"fcconsign"  form:"fcconsign" `
	// FCCONTACTN     string     `gorm:"column:FCCONTACTN;" json:"fccontactn"  form:"fccontactn" `
	FCCOOR     string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	// FCCOUNTER      string     `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	// FCCREATEAP     string     `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCCUACC        string     `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCCURRENCY     string     `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	// FCDATAIMP      string     `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp"`
	FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" default:"$$$+"`
	// FCDEBTCODE     string     `gorm:"column:FCDEBTCODE;" json:"fcdebtcode"  form:"fcdebtcode" `
	FCDELICOOR string `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	// FCDELIEMPL     string     `gorm:"column:FCDELIEMPL;" json:"fcdeliempl"  form:"fcdeliempl" `
	FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	// FCDISCSTR      string     `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	// FCDOCAPRVB     string     `gorm:"column:FCDOCAPRVB;" json:"fcdocaprvb"  form:"fcdocaprvb" `
	// FCDOCFLOWI     string     `gorm:"column:FCDOCFLOWI;" json:"fcdocflowi"  form:"fcdocflowi" `
	// FCDOCOWNER     string     `gorm:"column:FCDOCOWNER;" json:"fcdocowner"  form:"fcdocowner" `
	// FCDOCSTEP      string     `gorm:"column:FCDOCSTEP;" json:"fcdocstep"  form:"fcdocstep" `
	// FCDTYPE1       string     `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FCDTYPE2       string     `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FCDTYPE3       string     `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FCDTYPE4       string     `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FCDTYPE5       string     `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FCDTYPE6       string     `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FCDTYPE7       string     `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FCDTYPE8       string     `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FCDTYPE9       string     `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	// FCEMPL         string     `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	// FCFRLOCATE     string     `gorm:"column:FCFRLOCATE;" json:"fcfrlocate"  form:"fcfrlocate" `
	// FCFRSTORE      string     `gorm:"column:FCFRSTORE;" json:"fcfrstore"  form:"fcfrstore" `
	// FCFRWHOUSE     string     `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCGID    string `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD string `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	// FCGLHEADCA     string     `gorm:"column:FCGLHEADCA;" json:"fcglheadca"  form:"fcglheadca" `
	// FCHASCHAIN     string     `gorm:"column:FCHASCHAIN;" json:"fchaschain"  form:"fchaschain" `
	// FCHASRET       string     `gorm:"column:FCHASRET;" json:"fchasret"  form:"fchasret" `
	// FCINVIA        string     `gorm:"column:FCINVIA;" json:"fcinvia"  form:"fcinvia" `
	// FCISCASH       string     `gorm:"column:FCISCASH;" json:"fciscash"  form:"fciscash" `
	FCJOB string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	// FCLAYH         string     `gorm:"column:FCLAYH;" json:"fclayh"  form:"fclayh" `
	// FCLAYMETHD     string     `gorm:"column:FCLAYMETHD;" json:"fclaymethd"  form:"fclaymethd" `
	// FCLAYSEQ       string     `gorm:"column:FCLAYSEQ;" json:"fclayseq"  form:"fclayseq" `
	// FCLID          string     `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	// FCLINKAPP1     string     `gorm:"column:FCLINKAPP1;" json:"fclinkapp1"  form:"fclinkapp1" `
	// FCLINKAPP2     string     `gorm:"column:FCLINKAPP2;" json:"fclinkapp2"  form:"fclinkapp2" `
	// FCLINKSTP1     string     `gorm:"column:FCLINKSTP1;" json:"fclinkstp1"  form:"fclinkstp1" `
	// FCLINKSTP2     string     `gorm:"column:FCLINKSTP2;" json:"fclinkstp2"  form:"fclinkstp2" `
	// FCLUPDAPP      string     `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	// FCMACHINE      string     `gorm:"column:FCMACHINE;" json:"fcmachine"  form:"fcmachine" `
	// FCMNMGLH       string     `gorm:"column:FCMNMGLH;" json:"fcmnmglh"  form:"fcmnmglh" `
	// FCMNMGLHCA     string     `gorm:"column:FCMNMGLHCA;" json:"fcmnmglhca"  form:"fcmnmglhca" `
	// FCMORDERH      string     `gorm:"column:FCMORDERH;" json:"fcmorderh"  form:"fcmorderh" `
	// FCORGCODE      string     `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCOUTVIA       string     `gorm:"column:FCOUTVIA;" json:"fcoutvia"  form:"fcoutvia" `
	// FCPAYTERM      string     `gorm:"column:FCPAYTERM;" json:"fcpayterm"  form:"fcpayterm" `
	// FCPCONTRAC     string     `gorm:"column:FCPCONTRAC;" json:"fcpcontrac"  form:"fcpcontrac" `
	// FCPDBRAND      string     `gorm:"column:FCPDBRAND;" json:"fcpdbrand"  form:"fcpdbrand" `
	// FCPERIODS      string     `gorm:"column:FCPERIODS;" json:"fcperiods"  form:"fcperiods" `
	// FCPERSON       string     `gorm:"column:FCPERSON;" json:"fcperson"  form:"fcperson" `
	// FCPLANT        string     `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	// FCPROMOTE      string     `gorm:"column:FCPROMOTE;" json:"fcpromote"  form:"fcpromote" `
	// FCQCMO         string     `gorm:"column:FCQCMO;" json:"fcqcmo"  form:"fcqcmo" `
	// FCQCMOBK       string     `gorm:"column:FCQCMOBK;" json:"fcqcmobk"  form:"fcqcmobk" `
	// FCRDSTEP       string     `gorm:"column:FCRDSTEP;" json:"fcrdstep"  form:"fcrdstep" `
	// FCRECALBY      string     `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	// FCRECVMAN      string     `gorm:"column:FCRECVMAN;" json:"fcrecvman"  form:"fcrecvman" `
	// FCREFDO        string     `gorm:"column:FCREFDO;" json:"fcrefdo"  form:"fcrefdo" `
	FCREFNO   string `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE  string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT    string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	// FCSELTAG       string     `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	// FCSEQ          string     `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	// FCSHOWPLED     string     `gorm:"column:FCSHOWPLED;" json:"fcshowpled"  form:"fcshowpled" `
	// FCSIGNETAX     string     `gorm:"column:FCSIGNETAX;" json:"fcsignetax"  form:"fcsignetax" `
	FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"`
	// FCSRCUPD       string     `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FCSTAT         string     `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" default:"I"`
	// FCSTEPD        string     `gorm:"column:FCSTEPD;" json:"fcstepd"  form:"fcstepd" `
	// FCSTEPX1       string     `gorm:"column:FCSTEPX1;" json:"fcstepx1"  form:"fcstepx1" `
	// FCSTEPX2       string     `gorm:"column:FCSTEPX2;" json:"fcstepx2"  form:"fcstepx2" `
	// FCSUBBR        string     `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	// FCSUBTXID      string     `gorm:"column:FCSUBTXID;" json:"fcsubtxid"  form:"fcsubtxid" `
	// FCTIMESTAM     string     `gorm:"column:FCTIMESTAM;" json:"fctimestam"  form:"fctimestam" `
	// FCTOLOCATE     string     `gorm:"column:FCTOLOCATE;" json:"fctolocate"  form:"fctolocate" `
	// FCTOSTORE      string     `gorm:"column:FCTOSTORE;" json:"fctostore"  form:"fctostore" `
	// FCTOWHOUSE     string     `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	// FCTRADETRM     string     `gorm:"column:FCTRADETRM;" json:"fctradetrm"  form:"fctradetrm" `
	// FCTRANWHY      string     `gorm:"column:FCTRANWHY;" json:"fctranwhy"  form:"fctranwhy" `
	// FCTXBOOK       string     `gorm:"column:FCTXBOOK;" json:"fctxbook"  form:"fctxbook" `
	// FCTXBRANCH     string     `gorm:"column:FCTXBRANCH;" json:"fctxbranch"  form:"fctxbranch" `
	// FCTXID         string     `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
	// FCTXQCBOOK     string     `gorm:"column:FCTXQCBOOK;" json:"fctxqcbook"  form:"fctxqcbook" `
	// FCTXQCBRAN     string     `gorm:"column:FCTXQCBRAN;" json:"fctxqcbran"  form:"fctxqcbran" `
	// FCTXQCGLRE     string     `gorm:"column:FCTXQCGLRE;" json:"fctxqcglre"  form:"fctxqcglre" `
	// FCTXQCWHOU     string     `gorm:"column:FCTXQCWHOU;" json:"fctxqcwhou"  form:"fctxqcwhou" `
	// FCTXWHOUSE     string     `gorm:"column:FCTXWHOUSE;" json:"fctxwhouse"  form:"fctxwhouse" `
	// FCU1ACC        string     `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCU1STATUS     string     `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	// FCU2STATUS     string     `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	// FCU3STATUS     string     `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	// FCU4STATUS     string     `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	// FCU5STATUS     string     `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	// FCU6STATUS     string     `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	// FCU7STATUS     string     `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	// FCU8STATUS     string     `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	// FCU9STATUS     string     `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	// FCUDATE        string     `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	// FCUTIME        string     `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVATCOOR string `gorm:"column:FCVATCOOR;" json:"fcvatcoor"  form:"fcvatcoor" `
	// FCVATDUE       string     `gorm:"column:FCVATDUE;" json:"fcvatdue"  form:"fcvatdue" `
	// FCVATISOUT     string     `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	// FCVATSEQ       string     `gorm:"column:FCVATSEQ;" json:"fcvatseq"  form:"fcvatseq" `
	// FCVATTYPE      string     `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	// FCVLDCTRL      string     `gorm:"column:FCVLDCTRL;" json:"fcvldctrl"  form:"fcvldctrl" `
	// FCWHTINT       string     `gorm:"column:FCWHTINT;" json:"fcwhtint"  form:"fcwhtint" `
	// FCXFERSTEP     string     `gorm:"column:FCXFERSTEP;" json:"fcxferstep"  form:"fcxferstep" `
	// FCZIP          string     `gorm:"column:FCZIP;" json:"fczip"  form:"fczip" `
	// FDAPPROVE      string     `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" `
	FDDATE time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	// FDDEBTDATE     string     `gorm:"column:FDDEBTDATE;" json:"fddebtdate"  form:"fddebtdate" `
	// FDDOCAPRV      string     `gorm:"column:FDDOCAPRV;" json:"fddocaprv"  form:"fddocaprv" `
	// FDDUEDATE      string     `gorm:"column:FDDUEDATE;" json:"fdduedate"  form:"fdduedate" `
	// FDFULLFILL     string     `gorm:"column:FDFULLFILL;" json:"fdfullfill"  form:"fdfullfill" `
	// FDLAYDATE      string     `gorm:"column:FDLAYDATE;" json:"fdlaydate"  form:"fdlaydate" `
	// FDRECEDATE     string     `gorm:"column:FDRECEDATE;" json:"fdrecedate"  form:"fdrecedate" `
	// FDVATDATE      string     `gorm:"column:FDVATDATE;" json:"fdvatdate"  form:"fdvatdate" `
	FIMILLISEC int64 `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	// FMDOCFLOW      string     `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
	// FMERRMSG       string     `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	// FMEXTRATAG     string     `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	// FMMEMDATA      string     `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	// FMMEMDATA2     string     `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
	// FMMEMDATA3     string     `gorm:"column:FMMEMDATA3;" json:"fmmemdata3"  form:"fmmemdata3" `
	// FMMEMDATA4     string     `gorm:"column:FMMEMDATA4;" json:"fmmemdata4"  form:"fmmemdata4" `
	// FMMEMDATA5     string     `gorm:"column:FMMEMDATA5;" json:"fmmemdata5"  form:"fmmemdata5" `
	// FMMEMDATA6     string     `gorm:"column:FMMEMDATA6;" json:"fmmemdata6"  form:"fmmemdata6" `
	// FMMEMDATA7     string     `gorm:"column:FMMEMDATA7;" json:"fmmemdata7"  form:"fmmemdata7" `
	// FMMEMDATA8     string     `gorm:"column:FMMEMDATA8;" json:"fmmemdata8"  form:"fmmemdata8" `
	// FMMEMDATA9     string     `gorm:"column:FMMEMDATA9;" json:"fmmemdata9"  form:"fmmemdata9" `
	// FMMEMDATAA     string     `gorm:"column:FMMEMDATAA;" json:"fmmemdataa"  form:"fmmemdataa" `
	FNAFTDEP   float64 `gorm:"column:FNAFTDEP;" json:"fnaftdep"  form:"fnaftdep" defualt:"0"`
	FNAFTDEPKE float64 `gorm:"column:FNAFTDEPKE;" json:"fnaftdepke"  form:"fnaftdepke" defualt:"0"`
	FNAMT      float64 `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	// FNAMT2         float64    `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
	// FNAMTKE        float64    `gorm:"column:FNAMTKE;" json:"fnamtke"  form:"fnamtke" `
	// FNAMTNOVAT     float64    `gorm:"column:FNAMTNOVAT;" json:"fnamtnovat"  form:"fnamtnovat" `
	// FNBEFDEP       float64    `gorm:"column:FNBEFDEP;" json:"fnbefdep"  form:"fnbefdep" `
	// FNBEFDEPKE     float64    `gorm:"column:FNBEFDEPKE;" json:"fnbefdepke"  form:"fnbefdepke" `
	// FNBEFOAMT      float64    `gorm:"column:FNBEFOAMT;" json:"fnbefoamt"  form:"fnbefoamt" `
	// FNBPAYINV      float64    `gorm:"column:FNBPAYINV;" json:"fnbpayinv"  form:"fnbpayinv" `
	// FNCREDTERM     float64    `gorm:"column:FNCREDTERM;" json:"fncredterm"  form:"fncredterm" `
	// FNCRXRATE      float64    `gorm:"column:FNCRXRATE;" json:"fncrxrate"  form:"fncrxrate" `
	// FNDEBTZERO     float64    `gorm:"column:FNDEBTZERO;" json:"fndebtzero"  form:"fndebtzero" `
	// FNDEPAMT       float64    `gorm:"column:FNDEPAMT;" json:"fndepamt"  form:"fndepamt" `
	// FNDEPAMTKE     float64    `gorm:"column:FNDEPAMTKE;" json:"fndepamtke"  form:"fndepamtke" `
	// FNDISCAMT1     float64    `gorm:"column:FNDISCAMT1;" json:"fndiscamt1"  form:"fndiscamt1" `
	// FNDISCAMT2     float64    `gorm:"column:FNDISCAMT2;" json:"fndiscamt2"  form:"fndiscamt2" `
	// FNDISCAMTA     float64    `gorm:"column:FNDISCAMTA;" json:"fndiscamta"  form:"fndiscamta" `
	// FNDISCAMTI     float64    `gorm:"column:FNDISCAMTI;" json:"fndiscamti"  form:"fndiscamti" `
	// FNDISCAMTK     float64    `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	// FNDISCPCN1     float64    `gorm:"column:FNDISCPCN1;" json:"fndiscpcn1"  form:"fndiscpcn1" `
	// FNDISCPCN2     float64    `gorm:"column:FNDISCPCN2;" json:"fndiscpcn2"  form:"fndiscpcn2" `
	// FNDISCPCN3     float64    `gorm:"column:FNDISCPCN3;" json:"fndiscpcn3"  form:"fndiscpcn3" `
	// FNEXPAMT       float64    `gorm:"column:FNEXPAMT;" json:"fnexpamt"  form:"fnexpamt" `
	// FNLOCKED       float64    `gorm:"column:FNLOCKED;" json:"fnlocked"  form:"fnlocked" `
	// FNPAYAMT       float64    `gorm:"column:FNPAYAMT;" json:"fnpayamt"  form:"fnpayamt" `
	// FNPAYAMTKE     float64    `gorm:"column:FNPAYAMTKE;" json:"fnpayamtke"  form:"fnpayamtke" `
	// FNPPAYAMT      float64    `gorm:"column:FNPPAYAMT;" json:"fnppayamt"  form:"fnppayamt" `
	// FNPPAYAMTK     float64    `gorm:"column:FNPPAYAMTK;" json:"fnppayamtk"  form:"fnppayamtk" `
	// FNRETENAMT     float64    `gorm:"column:FNRETENAMT;" json:"fnretenamt"  form:"fnretenamt" `
	// FNREVALAMT     float64    `gorm:"column:FNREVALAMT;" json:"fnrevalamt"  form:"fnrevalamt" `
	// FNSPAYAMT      float64    `gorm:"column:FNSPAYAMT;" json:"fnspayamt"  form:"fnspayamt" `
	// FNSTOCKUPD     float64    `gorm:"column:FNSTOCKUPD;" json:"fnstockupd"  form:"fnstockupd" `
	// FNU1CNT        float64    `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FNU2CNT        float64    `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FNU3CNT        float64    `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FNU4CNT        float64    `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FNU5CNT        float64    `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FNU6CNT        float64    `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FNU7CNT        float64    `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FNU8CNT        float64    `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FNU9CNT        float64    `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	// FNVATAMT       float64    `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	// FNVATAMTKE     float64    `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
	// FNVATPORT      float64    `gorm:"column:FNVATPORT;" json:"fnvatport"  form:"fnvatport" `
	// FNVATPORTA     float64    `gorm:"column:FNVATPORTA;" json:"fnvatporta"  form:"fnvatporta" `
	// FNVATRATE      float64    `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	// FNXRATE        float64    `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	// FTLASRECAL     string     `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	// FTRDSTEP       string     `gorm:"column:FTRDSTEP;" json:"ftrdstep"  form:"ftrdstep" `
	// FTRECEIVE      string     `gorm:"column:FTRECEIVE;" json:"ftreceive"  form:"ftreceive" `
	// FTSRCUPD       string     `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	// FTXFER         string     `gorm:"column:FTXFER;" json:"ftxfer"  form:"ftxfer" `
	Corp           *Corp   `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	Branch         *Branch `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	Dept           *Dept   `gorm:"foreignKey:FCDEPT;references:FCSKID;" json:"department"`
	Sect           *Sect   `gorm:"foreignKey:FCSECT;references:FCSKID;" json:"section"`
	Job            *Job    `gorm:"foreignKey:FCJOB;references:FCSKID;" json:"job"`
	Glhead         *Glhead `gorm:"foreignKey:FCGLHEAD;references:FCSKID;" json:"glhead"`
	Book           *Book   `gorm:"foreignKey:FCBOOK;references:FCSKID;" json:"book"`
	Coor           *Coor   `gorm:"foreignKey:FCCOOR;reference:FCSKID;" json:"coor"`
	CreatedBy      *Empl   `gorm:"foreignKey:FCCREATEBY;references:FCSKID;" json:"created_by"`
	UpdatedBy      *Empl   `gorm:"foreignKey:FCCORRECTB;references:FCSKID;" json:"updated_by"`
	VatCoor        *Coor   `gorm:"foreignKey:FCVATCOOR;reference:FCSKID;" json:"vat_coor"`
	Proj           *Proj   `gorm:"foreignKey:FCPROJ;reference:FCSKID;" json:"proj"`
	DeliveryToCoor *Coor   `gorm:"foreignKey:FCDELICOOR;reference:FCSKID;" json:"delivery_to_coor"`
}

func (Glref) TableName() string {
	return "GLREF"
}

func (obj *Glref) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}

type GlRefForm struct {
	FCBOOK     string       `json:"fcbook" form:"fcbook"`
	FCBRANCH   string       `json:"fcbranch" form:"fcbranch"`
	FCWHOUSE   string       `json:"fcwhouse" form:"fcwhouse"`
	FCCODE     string       `json:"fccode"  form:"fccode" `
	FCGID      string       `json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string       `json:"fcglhead"  form:"fcglhead" `
	FCJOB      string       `json:"fcjob"  form:"fcjob" `
	FCPROJ     string       `json:"fcproj"  form:"fcproj" `
	FCREFNO    string       `json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string       `json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string       `json:"fcrftype"  form:"fcrftype" `
	FCSTEP     string       `json:"fcstep"  form:"fcstep" default:"I"`
	FDDATE     time.Time    `json:"fddate"  form:"fddate" default:"now"`
	FNAFTDEP   float64      `json:"fnaftdep"  form:"fnaftdep" defualt:"0"`
	FNAFTDEPKE float64      `json:"fnaftdepke"  form:"fnaftdepke" defualt:"0"`
	REFPROD    []FrmRefProd `json:"ref_prod" form:"ref_prod"`
}

type GlrefHistory struct {
}
