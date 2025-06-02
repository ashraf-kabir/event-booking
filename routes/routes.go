package routes

import (
	"event-booking/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// health check
	server.GET("/api/v1/health", controllers.GetHealth)

	// events
	server.GET("/api/v1/events", controllers.GetEvents)
	server.GET("/api/v1/events/:id", controllers.GetOneEvent)
	server.POST("/api/v1/event", controllers.CreateEvent)
	server.PUT("/api/v1/event/:id", controllers.UpdateEvent)
	server.DELETE("/api/v1/event/:id", controllers.DeleteEvent)
}
