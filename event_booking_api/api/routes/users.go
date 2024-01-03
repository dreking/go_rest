package routes

import (
	"net/http"

	"github.com/dreking/event-booking-api/api/models"
	"github.com/dreking/event-booking-api/api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": err})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": "Something went wrong"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": true, "data": user})
}

func signIn(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "data": err})
		return
	}

	foundUser, err := models.FindByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "data": "Invalid credentials"})
		return
	}

	isValidPassword := utils.ComparePassword(foundUser.Password, user.Password)
	if !isValidPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "data": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(foundUser.Email, foundUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": token})
}
