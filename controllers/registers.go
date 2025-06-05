package controllers

import (
	"event-booking/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	event, err := models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Could not find event"})
		return
	}

	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()
	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"error": false, "message": "Event registration completed"})
}

func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	var event models.Event
	event.Id = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"error": false, "message": "Event registration cancelled successfully"})
}
