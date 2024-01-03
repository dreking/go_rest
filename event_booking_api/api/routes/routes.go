package routes

import (
	"github.com/dreking/event-booking-api/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getAllEvents)
	server.GET("/events/paginated", getAllEventsPaginated)
	server.GET("/events/:id", getEventById)

	authenticated := server.Group("/", middlewares.Authorization)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEventById)
	authenticated.DELETE("/events/:id", deleteEventById)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/cancel", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/signin", signIn)
}
