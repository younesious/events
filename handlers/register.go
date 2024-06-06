package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/younesious/events/models"
)

func registerForEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID!",
			"error":   err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	err = models.RegisterForEvent(id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error registering for event!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered for event successfully!",
	})
}

func cancelRegistration(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID!",
			"error":   err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	err = models.CancelRegistration(id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error canceling registration!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration canceled successfully!",
	})
}
