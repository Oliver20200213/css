package model

type Class struct {
	BaseModel
	Num     int
	TutorID int
	Tutor   Teacher `gorm:"foreignKey:TutorID;constraint:OnDelete:CASCADE;"`
}
