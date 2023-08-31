package model

type Course struct {
	BaseModel
	Credit int
	Period int
	//多对一
	TeacherID int
	Teacher   Teacher
}
