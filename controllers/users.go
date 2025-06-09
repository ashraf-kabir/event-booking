package controllers

import (
	"event-booking/models"
	"event-booking/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = user.SaveUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"error": false, "message": "User created successfully!"})
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Could not authenticate"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Could not authenticate"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "token": token})
}
