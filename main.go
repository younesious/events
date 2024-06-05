package main

import (
	"github.com/gin-gonic/gin"
	"github.com/younesious/events/db"
	"github.com/younesious/events/routes"
)

func main() {
	db.InitDB()
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run("localhost:8080")
}
