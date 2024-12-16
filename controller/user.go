package controller

import (
	"fmt"
	"net/http"

	"github.com/AudiWu/mtg-backend-project/model"
	"github.com/AudiWu/mtg-backend-project/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	fmt.Println("Create user controller")

	var input model.User

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		res := utils.ApiErrorResponse(http.StatusBadRequest, "Error validating input", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	
	user, _ := input.FindUser(input.Username)
	
	if user.Username != "" {
		res := utils.ApiErrorResponse(http.StatusBadRequest, "user already exists", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err := input.Create()

	if err != nil {
		res := utils.ApiErrorResponse(http.StatusInternalServerError, "Error creating user!", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func FindUser(c *gin.Context) {
	fmt.Println("Find user controller")

	var input model.User

	username := c.Param("username")

	user, err := input.FindUser(username)

	if err != nil {
		res := utils.ApiErrorResponse(http.StatusInternalServerError,  "Error getting user!", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ApiResponse(http.StatusOK, user)
	c.JSON(http.StatusOK, res)
}
