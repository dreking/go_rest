package routes

import (
	"net/http"
	"strconv"

	"github.com/dreking/event-booking-api/api/models"
	"github.com/gin-gonic/gin"
)

type GetEventsQuery struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

func getAllEventsPaginated(c *gin.Context) {
	var query GetEventsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": err})
		return
	}

	events, err := models.GetAllEventsPaginated(query.Page, query.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": err})
		return
	}

	c.JSON(http.StatusOK, events)
}

func getAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": err})
		return
	}

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	// user, ok := c.Get("user") // didnt give me a way to tell go the returned data is a User struct
	user, ok := c.Keys["user"].(*models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": "Token invalid"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": err})
		return
	}

	event.UserID = user.ID
	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": true, "data": "Something went wrong"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": true, "data": event})
}

type GetEventByIdParams struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func getEventById(c *gin.Context) {
	var params GetEventByIdParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": err})
		return
	}

	event, err := models.GetEventById(params.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "data": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": event})
}

func updateEventById(c *gin.Context) {
	user, _ := c.Keys["user"].(*models.User)

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": "Event id is not an integer"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": "Something went wrong"})
		return
	}

	if event.UserID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "data": "Not authorized"})
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": err})
		return
	}

	updatedEvent.ID = 1
	updatedEvent.UserID = 1
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": "Something went wrong", "err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": updatedEvent})
}

func deleteEventById(c *gin.Context) {
	user, _ := c.Keys["user"].(*models.User)

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": "Event id is not an integer"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": "Something went wrong"})
		return
	}

	if event.UserID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "data": "Not authorized"})
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": "Deleted successully"})
}
