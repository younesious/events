package routes

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

	userID, _ := c.Get("userID")

	err = models.RegisterForEvent(id, userID.(int64))
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

	userID, _ := c.Get("userID")

	err = models.CancelRegistration(id, userID.(int64))
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

func signup(c *gin.Context) {
	var u models.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload!",
			"error":   err.Error(),
		})
		return
	}

	err = u.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating user!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully!",
		"user":    u,
	})
}

func login(c *gin.Context) {
	var u models.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload!",
			"error":   err.Error(),
		})
		return
	}

	user, err := models.AuthenticateUser(u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication failed!",
			"error":   err.Error(),
		})
		return
	}

	token, err := models.GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating token!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Authenticated successfully!",
		"token":   token,
	})
}
