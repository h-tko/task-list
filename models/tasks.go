package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TaskStatus int

const (
	Entry  TaskStatus = 0
	Ready  TaskStatus = 1
	Done   TaskStatus = 2
	Delete TaskStatus = 9
)

type Task struct {
	gorm.Model
	Title            string
	Body             string
	Status           uint
	RegistMemberID   uint
	InChargeMemberID uint
	FixedAt          time.Time
	RegistMember     Member `gorm:members;AssociationForeignKey:ID;ForeignKey:RegistMemberID`
	InChargeMember   Member `gorm:members;AssociationForeignKey:ID;FOreignKey:InChargeMemberID`
}

func (Task) TableName() string {
	return "tasks"
}

func NewTask() *Task {
	return new(Task)
}

func (model *Task) All() []*Task {
	var tasks []*Task

	model.base(db).Order("tasks.created_at desc").Find(&tasks)

	for _, data := range tasks {
		db.Model(&data).Related(&data.RegistMember, "RegistMemberID")
		db.Model(&data).Related(&data.InChargeMember, "InChargeMemberID")
	}

	return tasks
}

func (model *Task) Search(freeword string) []*Task {
	var tasks []*Task
	likeFreeword := "%" + freeword + "%"

	model.base(db).Where("(title like ? OR body like ?)", likeFreeword, likeFreeword).Order("tasks.created_at desc").Find(&tasks)

	for _, data := range tasks {
		db.Model(&data).Related(&data.RegistMember, "RegistMemberID")
		db.Model(&data).Related(&data.InChargeMember, "InChargeMemberID")
	}

	return tasks
}

func (model *Task) base(db *gorm.DB) *gorm.DB {
	db.
		Joins("inner join members as regist on regist.id = tasks.regist_member_id").
		Joins("left join members as incharge on incharge.id = tasks.in_charge_member_id").
		Where("status <> ?", Delete)

	return db
}

func (model *Task) Regist() {
	db.NewRecord(model)
	db.Create(&model)
}

func (model *Task) One(id int) {
	model.base(db).First(&model, id)

	db.Model(&model).Related(&model.RegistMember, "RegistMemberID")
	db.Model(&model).Related(&model.InChargeMember, "InChargeMemberID")
}
