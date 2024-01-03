package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dreking/event-booking-api/api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	user, _ := c.Keys["user"].(models.User)
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": "Event id is required"})
		return
	}

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": "Event id is not an integer"})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "data": "Event not found"})
		return
	}

	isRegistered, err := models.FindUserRegistration(user.ID, eventId)
	fmt.Println(err)
	if err == nil && isRegistered != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": "You already registered to this event"})
		return
	}

	var registration models.Registration
	registration.UserID = user.ID
	registration.EventID = eventId
	err = registration.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": true, "data": "Something went wrong"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": true, "data": "Registration done"})
}

func cancelRegistration(c *gin.Context) {
	user, _ := c.Keys["user"].(models.User)
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": "Event id is required"})
		return
	}

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": "Event id is not an integer"})
		return
	}

	err = models.CancelRegistration(user.ID, eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": "Deleted registration"})
}
