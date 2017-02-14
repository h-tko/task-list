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
