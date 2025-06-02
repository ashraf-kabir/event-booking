package main

import (
	"event-booking/db"
	"event-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/api/v1/health", getHealth)

	// events
	server.GET("/api/v1/events", getEvents)
	server.GET("/api/v1/events/:id", getOneEvent)
	server.POST("/api/v1/event", createEvent)

	server.Run(":8080") // http://localhost:8080
}

func getHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"error": false, "message": "OK"})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getOneEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	event, err := models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
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

	err = event.SaveEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"error": false, "message": "Event created!", "event": event})
}
