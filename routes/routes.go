package routes

import (
	"example.com/go-event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEventById)
	authenticated.DELETE("/events/:id", deleteEventById)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
