package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	host := os.Getenv("DB_HOST")
    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    databaseName := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", host, username, password, databaseName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
        panic(err)
    } else {
        fmt.Println("Successfully connected to the database")
    }

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}