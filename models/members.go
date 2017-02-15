package models

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name        string
	MailAddress string
	Password    string
}

func (Member) TableName() string {
	return "members"
}

func NewMember() *Member {
	return new(Member)
}

func (model *Member) One(mailAddress, password string) {
	db.Select("id, name, mail_address").
		Where("mail_address = ?", mailAddress).
		Where("password = ?", password).
		First(&model)
}
