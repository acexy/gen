package model

import "github.com/golang-acexy/starter-gorm/gormstarter"

const TableNameTeacher = "demo_teacher"

// Teacher 教师表
type Teacher struct {
	ID         int64                 `gorm:"primaryKey;<-:false" json:"id"`
	CreateTime gormstarter.Timestamp `json:"createTime"` // 创建时间
	UpdateTime gormstarter.Timestamp `json:"updateTime"` // 更新时间
	Name       string                `json:"name"`       // 姓名
	Sex        string                `json:"sex"`        // 性别
	Age        int32                 `json:"age"`        // 年龄
	ClassNo    string                `json:"classNo"`    // 班级编号
}

func (Teacher) TableName() string {
	return TableNameTeacher
}
