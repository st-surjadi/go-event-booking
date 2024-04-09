package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEventById)
	server.DELETE("/events/:id", deleteEventById)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
