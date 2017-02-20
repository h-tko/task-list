package models

import (
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type Member struct {
	gorm.Model
	Name        string `validate:"required"`
	MailAddress string `validate:"required"`
	Password    string `validate:"required"`
}

func (Member) TableName() string {
	return "members"
}

func NewMember() *Member {
	return new(Member)
}

func (model *Member) One(mailAddress, password string) {
	cryptedPassword := encryptPassword(password)

	db.Select("id, name, mail_address").
		Where("mail_address = ?", mailAddress).
		Where("password = ?", cryptedPassword).
		First(&model)
}

func (model *Member) FromID(id uint) {
	db.Select("id, name, mail_address").
		Where("id = ?", id).
		First(&model)
}

func (model *Member) Regist() {
	model.Password = encryptPassword(model.Password)
	db.NewRecord(model)
	db.Create(&model)
}

func encryptPassword(password string) string {
	crypt, _ := scrypt.Key([]byte(password), []byte("feaw98f3jk"), 16384, 8, 1, 16)

	return hex.EncodeToString(crypt[:])
}
