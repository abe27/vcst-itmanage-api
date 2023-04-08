package models

type Itablis1 struct {
	FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCFCHR     string `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCIDTYPE   string `gorm:"column:FCIDTYPE;" json:"fcidtype"  form:"fcidtype" `
	FCNAME     string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSTATUS   string `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FIMAXSKIDV string `gorm:"column:FIMAXSKIDV;" json:"fimaxskidv"  form:"fimaxskidv" `
	FINSKIDV   string `gorm:"column:FINSKIDV;" json:"finskidv"  form:"finskidv" `
	FIRECC01   string `gorm:"column:FIRECC01;" json:"firecc01"  form:"firecc01" `
	FIRECCOUNT string `gorm:"column:FIRECCOUNT;" json:"fireccount"  form:"fireccount" `
	FISKIDLEN  string `gorm:"column:FISKIDLEN;" json:"fiskidlen"  form:"fiskidlen" `
	FISKIDV    string `gorm:"column:FISKIDV;" json:"fiskidv"  form:"fiskidv" `
	FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FTDATETIME string `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" `
	FTLASTEDIT string `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  string `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" `
}

func (Itablis1) TableName() string {
	return "ITABLIS1"
}
