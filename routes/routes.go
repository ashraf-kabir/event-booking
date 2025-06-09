package routes

import (
	"event-booking/controllers"
	"event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// health check
	server.GET("/api/v1/health", controllers.GetHealth)

	// authenticated event routes
	authenticated := server.Group("/api/v1/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/", controllers.GetEvents)
	authenticated.GET("/:id", controllers.GetOneEvent)
	authenticated.POST("/", controllers.CreateEvent)
	authenticated.PUT("/:id", controllers.UpdateEvent)
	authenticated.DELETE("/:id", controllers.DeleteEvent)
	authenticated.POST("/:id/register", controllers.RegisterForEvent)
	authenticated.DELETE("/:id/register", controllers.CancelRegistration)

	// public auth routes
	server.POST("/api/v1/signup", controllers.SignUp)
	server.POST("/api/v1/login", controllers.Login)
}
