package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID           string      `gorm:"size:21;primaryKey;not null;" json:"id"`
	AvatarUrl    string      `json:"avatar_url"`
	FirstName    string      `gorm:"size:20;not null;" json:"first_name"`
	LastName     string      `gorm:"size:20;not null;" json:"last_name"`
	Email        string      `gorm:"size:120;" json:"email"`
	UserName     string      `gorm:"column:empid;size:50;index;unique;" json:"username"` //-->emp_code
	Password     string      `gorm:"size:60;" json:"-"`
	CompanyID    *string     `json:"company_id"`
	PositionID   *string     `json:"position_id"`
	SectionID    *string     `json:"section_id"`
	DepartmentID *string     `json:"department_id"`
	PermissionID *string     `json:"permision_id"`
	WhsID        *string     `json:"whs_id"`
	IsActive     bool        `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt    time.Time   `json:"created_at" default:"now"`
	UpdatedAt    time.Time   `json:"updated_at" default:"now"`
	Company      *Company    `gorm:"foreignKey:CompanyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`
	Position     *Position   `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"position"`
	Section      *Section    `gorm:"foreignKey:SectionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"section"`
	Department   *Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"department"`
	Permission   *Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"permission"`
	Whs          *Whs        `gorm:"foreignKey:WhsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"whs"`
}

func (obj *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type FrmRegister struct {
	AvatarUrl    string `form:"avatar_url"`
	FirstName    string `form:"first_name"`
	LastName     string `form:"last_name"`
	Email        string `form:"email"`
	UserName     string `form:"user_name"` //-->emp_code
	Password     string `form:"password"`
	CompanyID    string `form:"company_id"`
	PositionID   string `form:"position_id"`
	SectionID    string `form:"section_id"`
	DepartmentID string `form:"department_id"`
	PermissionID string `form:"permision_id"`
	WhsID        string `form:"whs_id"`
	IsActive     bool   `form:"is_active"`
}

type ActivityLogging struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	IsActive    bool      `gorm:"null" json:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
	User        *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

func (obj *ActivityLogging) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type AuthSession struct {
	ID       string `json:"id"`
	Header   string `json:"header"`
	JwtType  string `json:"type"`
	JwtToken string `json:"token"`
	User     *User  `json:"user"`
}

type UserLoginForm struct {
	UserName string `json:"username"` //-->emp_code
	Password string `json:"password"`
}
