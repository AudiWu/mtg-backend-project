package controller

import (
	"fmt"

	"github.com/AudiWu/mtg-backend-project/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	fmt.Println("Create user controller")

	user := model.User{
		Username: "testAudi",
		Email: "test@email.com",
	}

	user.Create()
}
