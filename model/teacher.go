package model

import (
	"time"
)

type Teacher struct {
	BaseModel
	Tel    string `gorm:"type:char(11)"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255)"`

	UserID int
	User   User
}
