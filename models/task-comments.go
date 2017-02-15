package models

import (
	"github.com/jinzhu/gorm"
)

type TaskComment struct {
	gorm.Model
	TaskID   uint
	MemberID uint
	Comment  string
	Member   Member `gorm:members;AssociationForeignKey:ID;ForeignKey:MemberID`
}

func (TaskComment) TableName() string {
	return "task_comments"
}

func NewTaskComment() *TaskComment {
	return new(TaskComment)
}

func (model *TaskComment) ListByTaskID(taskID int) []*TaskComment {
	var taskComments []*TaskComment

	db.Where("task_id = ?", taskID).Find(&taskComments)

	for _, data := range taskComments {
		db.Model(&data).Related(&data.Member, "MemberID")
	}

	return taskComments
}
