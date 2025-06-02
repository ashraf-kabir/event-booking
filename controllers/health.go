package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"error": false, "message": "OK"})
}
