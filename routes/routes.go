package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.POST("signup", signup)
	r.POST("/login", login)

	r.GET("/events", getAllEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/events", createEvent)
	r.PUT("/events/:id", updateEvent)
	r.DELETE("/events/:id", deleteEvent)
}
