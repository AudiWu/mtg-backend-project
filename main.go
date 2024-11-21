package main

import (
	"log"

	"github.com/AudiWu/mtg-backend-project/db"
	"github.com/AudiWu/mtg-backend-project/migrate"
	"github.com/AudiWu/mtg-backend-project/router"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()

	router.Serve()
}

func loadDatabase() {
	db.Init()

	migrate.Check()
}

func loadEnv() {
	err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}
