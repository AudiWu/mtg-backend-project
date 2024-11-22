package controller

import (
	"fmt"
	"net/http"

	"github.com/AudiWu/mtg-backend-project/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	fmt.Println("Create user controller")

	var input model.User

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error validating input",
		})
		return
	}

	user, _ := input.FindUser(input.Username)

	if user.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "User already exists",
			"message": "User already exists",
		})
		return
	}

	err := input.Create()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error creating user!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create user success!",
	})
}

func FindUser(c *gin.Context) {
	fmt.Println("Find user controller")

	var input model.User

	username := c.Param("username")

	user, err := input.FindUser(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error getting user!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "Find user success!",
	})
}
