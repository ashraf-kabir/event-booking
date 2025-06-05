package controllers

import (
	"event-booking/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetOneEvent(context *gin.Context) {
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

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	err = event.SaveEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"error": false, "message": "Event created!", "event": event})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Not authorized"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	updatedEvent.Id = eventId
	updatedEvent.UserId = userId
	updatedEvent.CreatedAt = event.CreatedAt
	updatedEvent.UpdatedAt = time.Now()
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"error": false, "message": "Event updated!", "event": updatedEvent})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(int(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Not authorized"})
		return
	}

	err = event.DeleteEventById()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"error": false, "message": "Event deleted successfully!"})
}
