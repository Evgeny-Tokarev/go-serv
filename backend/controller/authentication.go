package controller

import (
	"fmt"
	"github.com/Evgeny-Tokarev/go-serv/helper"
	"github.com/Evgeny-Tokarev/go-serv/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context) {
	var input model.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err)
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}
	fmt.Println(user)
	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cookie, err := helper.GenerateJwtCookie(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.SetCookie(cookie.Name, cookie.Value, int(cookie.Expires.Unix()), cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	context.JSON(http.StatusOK, user)
}
func GetCurrentUser(c *gin.Context) {
	user, err := helper.CurrentUser(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, user)
}
