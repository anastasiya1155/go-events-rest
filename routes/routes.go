package routes

import (
	"github.com/anastasiya1155/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authRoutes := server.Group("/")
	authRoutes.Use(middleware.CheckAuth)
	authRoutes.POST("/events", createEvent)
	authRoutes.PUT("/events/:id", updateEvent)
	authRoutes.DELETE("/events/:id", deleteEvent)
	authRoutes.POST("/events/:id/register", registerForEvent)
	authRoutes.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
