package model

import "time"

type BaseModel struct {
	ID         int        `gorm:"primaryKey"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
	Name       string     `gorm:"type:varchar(32);unique;not null"`
}
