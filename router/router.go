package router

import (
	"os"

	"github.com/AudiWu/mtg-backend-project/controller"
	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()

	publicRoutes := router.Group("/user")
	publicRoutes.POST("/create", controller.CreateUser)
	publicRoutes.GET("/find/:username", controller.FindUser)

	// for third party api
	// router.GET("/land-types", func(c *gin.Context) {
	// 	client := http.Client{}
	// 	url := os.Getenv("SCRYFALL_API_URL") + "catalog/land-types"
	// 	req, err := http.NewRequest("GET", url, nil)

	// 	if err != nil {
	// 		log.Fatalf("Failed to create request object for /GET endpoint: %v", err)
	// 	}

	// 	req.Header.Add("Content-type", "application/json; charset=utf-8")

	// 	res, err := client.Do(req)

	// 	if err != nil || res.StatusCode != 200 {
	// 		log.Fatalf("Failed to send HTTP request: %v", err)
	// 	}

	// 	defer res.Body.Close()

	// 	body, err := io.ReadAll(res.Body)

	// 	if err != nil {
	// 		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.Data(200, "application/json", body)

	// })

	router.Run(":" + os.Getenv("PORT"))
}
