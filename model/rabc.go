package model

type User struct {
	ID      int `gorm:"primaryKey"`
	Account string
	Pwd     string `gorm:"type:varchar(255);not null"`
	//一对一角色表
	RoleID int
	Role   Role
}

type Role struct {
	BaseModel
	Permissions []Permission `gorm:"many2many:role2permission;constraint:OnDelete:CASCADE"`
}

type Permission struct {
	BaseModel
	Url string `gorm:"not null"`
}

type Admin struct {
	BaseModel
	//一对一用户表
	UserID int `gorm:"unique"`
	User   User
}
