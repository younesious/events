package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/younesious/events/models"
)

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

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"error":   err.Error(),
		})
		return
	}
	event, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func createEvent(c *gin.Context) {
	var e models.Event
	userID := c.GetInt64("userID")

	err := c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	e.UserID = userID
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

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"error":   err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	var e models.Event
	err = c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	existingEvent, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching event",
			"error":   err.Error(),
		})
		return
	}

	if existingEvent.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to update this event",
		})
		return
	}

	e.ID = id
	e.UserID = userID
	err = e.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully!",
		"event":   e,
	})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"error":   err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	existingEvent, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching event",
			"error":   err.Error(),
		})
		return
	}

	if existingEvent.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to delete this event",
		})
		return
	}

	err = models.DeleteEvent(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully!",
	})
}
