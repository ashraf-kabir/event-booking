package main

import (
	"event-booking/db"
	"event-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/api/v1/health", getHealth)

	// events
	server.GET("/api/v1/events", getEvents)
	server.POST("/api/v1/event", createEvent)

	server.Run(":8080") // http://localhost:8080
}

func getHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"error": false, "message": "OK"})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	event.Id = 1
	event.UserId = 1

	event.SaveEvent()

	context.JSON(http.StatusCreated, gin.H{"error": false, "message": "Event created!", "event": event})
}
