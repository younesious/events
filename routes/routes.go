package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", signup)
	r.POST("/login", login)

	r.GET("/events", getAllEvents)
	r.GET("/events/:id", getEvent)

	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)
	auth.POST("/events/:id/register", registerForEvent)
	auth.DELETE("/events/:id/register", cancelRegistration)
}
