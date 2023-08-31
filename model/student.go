package model

import "time"

type Student struct {
	BaseModel
	Sno    int
	Tel    string `gorm:"type:char(11)"`
	Gender byte   `gorm:"default 1"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`

	//多对一
	ClassID int
	Class   Class

	//多对多
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`

	//一对一用户表
	UserID int `gorm:"unique"`
	User   User
}
