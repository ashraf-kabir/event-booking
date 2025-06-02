package controllers

import (
	"event-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

	event.Id = 1
	event.UserId = 1

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

	_, err = models.GetEventById(int(eventId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	updatedEvent.Id = int(eventId)
	updatedEvent.UpdatedAt = time.Now()
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"error": false, "message": "Event updated!", "event": updatedEvent})
}
