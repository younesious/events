package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/younesious/events/db"
	"github.com/younesious/events/models"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/events", getAllEvents)
	r.POST("/events", createEvent)

	fmt.Println("\n-----------------")

	r.Run("localhost:8080")
}

func getAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching events",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func createEvent(c *gin.Context) {
	var e models.Event

	err := c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	err = e.CreateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully!",
		"event":   e,
	})
}
