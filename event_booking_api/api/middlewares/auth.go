package middlewares

import (
	"net/http"
	"strings"

	"github.com/dreking/event-booking-api/api/models"
	"github.com/dreking/event-booking-api/api/utils"
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "data": "Token invalid"})
		return
	}

	parsedToken := strings.Split(token, " ")
	signature, token := parsedToken[0], parsedToken[1]
	if signature != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "data": "Token invalid"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "data": "Token invalid"})
		return
	}

	var user *models.User
	user, err = models.FindById(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "data": "Token invalid"})
		return
	}

	c.Set("user", user)
	c.Next()
}
