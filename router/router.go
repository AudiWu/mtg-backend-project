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

	router.Run(":" + os.Getenv("PORT"))
}
