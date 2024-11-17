package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/land-types", func(c *gin.Context) {
		client := http.Client{}
		url := os.Getenv("SCRYFALL_API_URL") + "catalog/land-types"
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			log.Fatalf("Failed to create request object for /GET endpoint: %v", err)
		}

		req.Header.Add("Content-type", "application/json; charset=utf-8")

		res, err := client.Do(req)

		if err != nil || res.StatusCode != 200 {
			log.Fatalf("Failed to send HTTP request: %v", err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			c.JSON(res.StatusCode, gin.H{"error": err.Error()})
			return
		}

		c.Data(200, "application/json", body)
		
	})

	r.Run(":" + os.Getenv("PORT"))
}
