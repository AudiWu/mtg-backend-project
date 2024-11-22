package model

import (
	"log"

	"github.com/AudiWu/mtg-backend-project/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username" binding:"required"`
	Email    string `gorm:"size:255;not null;unique" json:"email" binding:"required"`
}

func (user *User) Create() error {
	db := db.GetDB()

	result := db.Create(&user)
	err := result.Error

	if err != nil {
		log.Panic("Create user failed!")
		return err
	}

	return nil
}

func (user *User) FindUser(username string) (*User, error) {
	db := db.GetDB()

	result := db.Where("username=?", username).Find(&user)
	err := result.Error

	if err != nil {
		log.Panic("Find user failed!")
		return &User{}, err
	}
	
	return user, nil
}
