package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/younesious/events/docs"
	"github.com/younesious/events/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	r.GET("/events", handlers.GetAllEvents)
	r.GET("/events/:id", handlers.GetEvent)

	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	auth.POST("/events", handlers.CreateEvent)
	auth.PUT("/events/:id", handlers.UpdateEvent)
	auth.DELETE("/events/:id", handlers.DeleteEvent)
	auth.POST("/events/:id/register", handlers.RegisterForEvent)
	auth.DELETE("/events/:id/register", handlers.CancelRegistration)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
