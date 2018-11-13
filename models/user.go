package models

import (
	"github.com/kenken64/saloon-server/utils"
)

type User struct {
	BaseModel
	//Id        int    `gorm:"AUTO_INCREMENT"`
	Firstname string `sql:"size:100"`
	Lastname  string `sql:"size:100"`
	Password  string `sql:"size:200"`
	Email     string `gorm:"unique_index;not null"`
	IsAdmin   bool   `sql:"size:1"`
}

func (u *User) HashPassword(plain string) (string, error) {
	return utils.HashAndSalt([]byte(plain))
}

func (u *User) CheckPassword(plain string) bool {
	err := utils.VerifyPassword(plain, []byte(u.Password))
	return err == nil
}
