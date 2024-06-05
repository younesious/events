package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getAllEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/events", createEvent)
}
