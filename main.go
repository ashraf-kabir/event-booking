package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/api/v1/health", getHealth)

	server.Run(":8080") // http://localhost:8080
}

func getHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"error": false, "message": "OK"})
}
