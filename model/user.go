package model

import (
	"log"

	"github.com/AudiWu/mtg-backend-project/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
}

func (user *User) Create() {
	db := db.GetDB()

	result := db.Create(&user)
	err := result.Error

	if err != nil {
		log.Fatal("Create user failed!")
	}
}
