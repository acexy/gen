package model

import "github.com/golang-acexy/starter-gorm/gormstarter"

const TableNameStudent = "demo_student"

// Student mapped from table <demo_student>
type Student struct {
	ID         int64                 `gorm:"primaryKey;<-:false" json:"id"`
	CreateTime gormstarter.Timestamp `json:"createTime"`
	UpdateTime gormstarter.Timestamp `json:"updateTime"`
	Name       string                `json:"name"`
	Sex        string                `json:"sex"`
	Age        int32                 `json:"age"`
	TeacherID  int64                 `json:"teacherId"`
}

func (Student) TableName() string {
	return TableNameStudent
}
